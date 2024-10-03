import React from "react";
import { Outlet } from "react-router-dom";
import { Provider } from "react-redux";
import { persistor, store } from "@app/store";
import { PersistGate } from "redux-persist/integration/react";
import { ThemeProvider } from "@app/components/themes/theme-provider"

const App: React.FC = () => {
  return (
    <ThemeProvider defaultTheme="dark" storageKey="vite-ui-theme">
      <Provider store={store}>
        <PersistGate loading={null} persistor={persistor}>
          <Outlet />
        </PersistGate>
      </Provider>
    </ThemeProvider>
  );
};

export default App;