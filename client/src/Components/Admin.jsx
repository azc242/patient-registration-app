import React from "react";
import App from "./App";
import "./App.css";
import Login from "./Login.jsx";

function Admin() {
    return(
        <div className="container">
            <h1>Admin Sign In</h1>
            <h5>If you are not an admin, please head to the home page and register.</h5>
            <Login />
        </div>
    )
}

export default Admin;