import React from "react";
import { Outlet } from "react-router-dom";
import { Provider } from "react-redux";
import { store } from "@app/store";

const App: React.FC = () => {
  return (
    <Provider store={store}>
      <Outlet />
    </Provider>
  );
};

export default App;