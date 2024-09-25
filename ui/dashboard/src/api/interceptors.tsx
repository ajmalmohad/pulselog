import { useEffect } from "react";
import { useDispatch, useSelector } from "react-redux";
import { AxiosInstance, AxiosError } from "axios";
import { RootState, AppDispatch } from "@app/store";
import { identityAPIHandler } from "@app/api/handlers";
import { setTokens } from "@app/store/auth/authSlice";

const useSetupInterceptors = (axiosInstance: AxiosInstance) => {
  const { accessToken, refreshToken } = useSelector((state: RootState) => state.auth);
  const dispatch = useDispatch<AppDispatch>();

  useEffect(() => {
    let isRefreshing = false;
    let refreshQueue: Array<(token: string) => void> = [];

    const addAuthHeader = (config: any) => {
      if (accessToken) {
        config.headers.Authorization = `Bearer ${accessToken}`;
      }
      return config;
    };

    const handleRequestError = (error: any) => Promise.reject(error);

    const handleResponseSuccess = (response: any) => response;

    const handleResponseError = async (error: AxiosError) => {
      const originalRequest = error.config as any;

      if ((originalRequest.url === "/auth/login" || originalRequest.url === "/auth/signup")) {
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

          try {
            const { data } = await identityAPIHandler.post("/auth/reauthenticate", { refresh_token: refreshToken });
            dispatch(setTokens({ accessToken: data.data.access_token }));
            axiosInstance.defaults.headers.common["Authorization"] = `Bearer ${data.data.access_token}`;

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
    axiosInstance.interceptors.response.use(handleResponseSuccess, handleResponseError);
  }, [accessToken, refreshToken, dispatch, axiosInstance]);
};

export default useSetupInterceptors;