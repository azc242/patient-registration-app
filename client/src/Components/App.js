import React from "react";
import { BrowserRouter as Router, Switch, Route } from "react-router-dom";
import './App.css';
import HomeRegister from "./HomeRegister";
import Error from "./Error.jsx";
import Admin from "./Admin.jsx";

function App() {
  return (
    <Router>
      <Switch>
        <Route exact path="/" component={HomeRegister}></Route>
        <Route exact path="/admin" component={Admin}></Route>
        <Route component={Error} />
      </Switch>
    </Router>

  );
}

export default App;
