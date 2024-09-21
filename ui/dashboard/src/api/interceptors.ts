import { AxiosInstance, AxiosError } from "axios";
import { AppDispatch, RootState } from "@app/store";
import { useDispatch, useSelector } from "react-redux";
import { identityAPIHandler } from "./handlers";
import { setTokens } from "@app/store/auth/authSlice";

export const setupInterceptors = (axiosInstance: AxiosInstance) => {
  const { accessToken, refreshToken } = useSelector((state: RootState) => state.auth);
  const dispatch = useDispatch<AppDispatch>();

  let isRefreshing = false;
  let refreshQueue: Array<(token: string) => void> = [];

  axiosInstance.interceptors.request.use(
    (config) => {
      if (accessToken) {
        config.headers.Authorization = `Bearer ${accessToken}`;
      }
      return config;
    },
    (error) => Promise.reject(error)
  );

  axiosInstance.interceptors.response.use(
    (response) => response,
    async (error: AxiosError) => {
      const originalRequest = error.config as any;

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
            dispatch(setTokens({ accessToken: data.access_token }));
            axiosInstance.defaults.headers.common["Authorization"] = `Bearer ${data.access_token}`;

            refreshQueue.forEach((callback) => callback(data.access_token));
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
    }
  );
};