import React from "react";
import { Outlet } from "react-router-dom";
import { Provider } from "react-redux";
import { store } from "@app/store";
import { ThemeProvider } from "@app/components/themes/theme-provider";
import { Toaster } from "@/components/ui/sonner";

const App: React.FC = () => {
  return (
    <ThemeProvider defaultTheme="light" storageKey="vite-ui-theme">
      <Provider store={store}>
        <Outlet />
        <Toaster />
      </Provider>
    </ThemeProvider>
  );
};

export default App;
