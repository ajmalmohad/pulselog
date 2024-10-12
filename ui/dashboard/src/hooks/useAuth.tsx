import { identityAPIHandler } from "@app/api/handlers";
import { toast } from "sonner";
import { useNavigate } from "react-router-dom";

export const useAuth = () => {
  const accessToken = localStorage.getItem("access_token");
  const refreshToken = localStorage.getItem("refresh_token");
  const navigate = useNavigate();

  const signup = async (name: string, email: string, password: string) => {
    const signupPromise = identityAPIHandler.post("/auth/signup", {
      name,
      email,
      password,
    });

    toast.promise(signupPromise, {
      loading: "Signing up...",
      success: "Signed up successfully!",
      error: "Signup failed. Please try again.",
    });

    await signupPromise.then(({ data }) => {
      localStorage.setItem("refresh_token", data.data.refresh_token);
      localStorage.setItem("access_token", data.data.access_token);
      navigate("/home");
    });
  };

  const login = async (email: string, password: string) => {
    const loginPromise = identityAPIHandler.post("/auth/login", {
      email,
      password,
    });

    toast.promise(loginPromise, {
      loading: "Logging in...",
      success: "Logged in successfully!",
      error: "Login failed. Please try again.",
    });

    await loginPromise.then(({ data }) => {
      localStorage.setItem("refresh_token", data.data.refresh_token);
      localStorage.setItem("access_token", data.data.access_token);
      navigate("/home");
    });
  };

  const logout = async () => {
    const logoutPromise = identityAPIHandler.delete("/users/logout", {
      data: { refresh_token: refreshToken },
    });

    toast.promise(logoutPromise, {
      loading: "Logging out...",
      success: "Logged out successfully!",
      error: "Logout failed. Please try again.",
    });

    await logoutPromise.then(() => {
      localStorage.clear();
      navigate("/auth/login");
    });
  };

  const isAuthenticated = !!accessToken;

  return { login, logout, signup, isAuthenticated, accessToken, refreshToken };
};
