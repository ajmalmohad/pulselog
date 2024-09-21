import Axios from "axios";
import { setupInterceptors } from "./interceptors";
import { env } from "./config";

export const identityAPIHandler = Axios.create({
  baseURL: env.IDENTITY_API_URL,
  headers: {
    Accept: "application/json",
    "Content-Type": "application/json",
  },
});

setupInterceptors(identityAPIHandler);