import Axios from "axios";
import { env } from "@app/api/config";

export const identityAPIHandler = Axios.create({
  baseURL: env.IDENTITY_API_URL,
  headers: {
    Accept: "application/json",
    "Content-Type": "application/json",
  },
});
