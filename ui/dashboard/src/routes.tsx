import type { RouteObject } from "react-router-dom";

export type AppRouteObject = RouteObject & {
  path?: string;
  element: React.ReactNode;
  children?: AppRouteObject[];
};

export const routes: AppRouteObject[] = [
    
];