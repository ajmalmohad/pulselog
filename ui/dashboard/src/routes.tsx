import Home from "@pages/Home";
import type { RouteObject } from "react-router-dom";
import { ProtectedRoute } from "./acl/ProtectedRoute";
import { LoginPage } from "@pages/Login";
import { SignupPage } from "@pages/Signup";
import { LandingPage } from "@pages/LandingPage";
import { ProtectedInverseRoute } from "./acl/ProtectedInverseRoute";

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
    element: (<ProtectedInverseRoute>
      <LoginPage />
    </ProtectedInverseRoute>)
    ,
  },
  {
    path: "/signup",
    element: (<ProtectedInverseRoute>
      <SignupPage />
    </ProtectedInverseRoute>)
    ,
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