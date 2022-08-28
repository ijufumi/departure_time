import React, { FC } from "react";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import Top from "./pages/Top";


interface Props {
}

const App: FC<Props> = () => {
  return (
    <BrowserRouter>
      <Routes>
        <Route
          path="/"
          element={<Top />}
        />
      </Routes>
    </BrowserRouter>
  );
};

export default App;
