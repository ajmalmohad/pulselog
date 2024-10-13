import {
  AxiosInstance,
  AxiosError,
  InternalAxiosRequestConfig,
  AxiosResponse,
} from "axios";
import { identityAPIHandler } from "@app/api/handlers";

export const setupInterceptors = (axiosInstance: AxiosInstance) => {
  let isRefreshing = false;
  let refreshQueue: Array<(token: string) => void> = [];

  const addAuthHeader = (config: InternalAxiosRequestConfig) => {
    const accessToken = localStorage.getItem("access_token");
    if (accessToken) {
      config.headers.Authorization = `Bearer ${accessToken}`;
    }
    return config;
  };

  const handleRequestError = (error: AxiosError) => Promise.reject(error);

  const handleResponseSuccess = (response: AxiosResponse) => response;

  const handleResponseError = async (error: AxiosError) => {
    const originalRequest = error.config as InternalAxiosRequestConfig & {
      _retry: boolean;
    };

    if (!originalRequest) return;

    if (
      originalRequest.url === "/auth/login" ||
      originalRequest.url === "/auth/signup"
    ) {
      return Promise.reject(error);
    }

    if (error.response?.status === 401 && !originalRequest._retry) {
      if (originalRequest.url === "/auth/reauthenticate") {
        window.location.href = "/";
        return Promise.reject(error);
      }

      if (!isRefreshing) {
        isRefreshing = true;
        originalRequest._retry = true;
        const refreshToken = localStorage.getItem("refresh_token");

        try {
          const { data } = await identityAPIHandler.post(
            "/auth/reauthenticate",
            { refresh_token: refreshToken }
          );
          localStorage.setItem("access_token", data.data.access_token);
          axiosInstance.defaults.headers.common[
            "Authorization"
          ] = `Bearer ${data.data.access_token}`;

          refreshQueue.forEach((callback) => callback(data.data.access_token));
          refreshQueue = [];

          return axiosInstance(originalRequest);
        } catch (refreshError) {
          return Promise.reject(refreshError);
        } finally {
          isRefreshing = false;
        }
      }

      return new Promise((resolve) => {
        refreshQueue.push((token) => {
          originalRequest.headers["Authorization"] = `Bearer ${token}`;
          resolve(axiosInstance(originalRequest));
        });
      });
    }

    return Promise.reject(error);
  };

  axiosInstance.interceptors.request.use(addAuthHeader, handleRequestError);
  axiosInstance.interceptors.response.use(
    handleResponseSuccess,
    handleResponseError
  );
};
