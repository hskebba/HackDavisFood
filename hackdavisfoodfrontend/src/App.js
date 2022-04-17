import './App.css';
import { BrowserRouter as Router, Route, Link, Routes } from "react-router-dom";


import {Inventory} from "./inventory/inventory"
import {Orders} from "./orders/orders"
import React from "react";

function App() {
  return (
    <div className="App">
        <Router>
            <Routes>
                <Route exact path="/" element={<h1>Home Page</h1>} />
                <Route exact path="orders" element={<Orders/>} />
                <Route exact path="inventory" element={<Inventory/>} />
            </Routes>
        </Router>
    </div>
  );
}

export default App;
