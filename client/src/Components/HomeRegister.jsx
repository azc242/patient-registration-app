import React from "react";
import "./App.css";
import Signup from "./Signup.jsx"

function HomeRegister() {
    return (
        <div className="container">
            <h1>Hello. Ready to enhance your healthcare?</h1>
            <h3>Sign up to see what we have to offer.</h3>
            <Signup />
        </div>
    )
}

export default HomeRegister;