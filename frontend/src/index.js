import React from "react";
import ReactDOM from "react-dom";
import App from "./app";

import { BrowserRouter as Router } from "react-router-dom";

import "./scss/index.scss";
import 'bootstrap/dist/css/bootstrap.min.css';

ReactDOM.render(

  <Router>
    <App />
  </Router>,

  document.getElementById("root")
);
