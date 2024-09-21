import React from "react";
import { Outlet } from "react-router-dom";
import { Provider } from "react-redux";
import { persistor, store } from "@app/store";
import { PersistGate } from "redux-persist/integration/react";

const App: React.FC = () => {  
  return (
    <Provider store={store}>
      <PersistGate loading={null} persistor={persistor}>
        <Outlet />
      </PersistGate>
    </Provider>
  );
};

export default App;