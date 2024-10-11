import React from "react";
import { Outlet } from "react-router-dom";
import { Provider } from "react-redux";
import { persistor, store } from "@app/store";
import { PersistGate } from "redux-persist/integration/react";
import { ThemeProvider } from "@app/components/themes/theme-provider"
import { Toaster } from "@/components/ui/sonner"

const App: React.FC = () => {
  return (
    <ThemeProvider defaultTheme="light" storageKey="vite-ui-theme">
      <Provider store={store}>
        <PersistGate loading={null} persistor={persistor}>
          <Outlet />
          <Toaster />
        </PersistGate>
      </Provider>
    </ThemeProvider>
  );
};

export default App;