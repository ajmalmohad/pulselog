import Home from "@pages/Home";
import type { RouteObject } from "react-router-dom";
import { ProtectedRoute } from "./acl/ProtectedRoute";
import { LoginPage } from "@pages/Login";
import { SignupPage } from "@pages/Signup";
import { LandingPage } from "@pages/LandingPage";

export type AppRouteObject = RouteObject & {
  path?: string;
  element: React.ReactNode;
  children?: AppRouteObject[];
};

export const routes: AppRouteObject[] = [
  {
    path: "/",
    element: <LandingPage />,
  },
  {
    path: "/login",
    element: <LoginPage />,
  },
  {
    path: "/signup",
    element: <SignupPage />,
  },
  {
    path: "/home",
    element: (
      <ProtectedRoute>
        <Home />
      </ProtectedRoute>
    ),
  }
];