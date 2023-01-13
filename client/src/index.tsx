import * as React from "react";
import * as ReactDOM from "react-dom";
import { BrowserRouter, Route, Routes } from "react-router-dom";

import GlobalStyles from "./GlobalStyles";

import Home from "./pages/Home";
import Dashboard from "./pages/Dashboard";

import NavigationBar from "./components/NavigationBar";

const index = (
  <BrowserRouter>
    <GlobalStyles />
    <NavigationBar />
    <Routes>
      <Route path="/" element={<Home />} />
      <Route path="/dashboard" element={<Dashboard />} />
    </Routes>
  </BrowserRouter>
);

ReactDOM.render(index, document.getElementById("root") as HTMLElement);
