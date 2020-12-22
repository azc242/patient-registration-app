import axios from "axios";
import React, { useState } from "react";
import "./App.css";

function Login(props) {
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");

    function authenticate(event) {
        axios.post('http://localhost:8080/api/login', {
            username: username,
            password: password
        }, {headers:{}}, { crossdomain: true })
            .then (res => {
                console.log(res);
                props.displayPatients();
            })
            .catch(function (error) {
                if (error.response) {
                    console.log(error.response.data);
                    console.log(error.response.status);
                    console.log(error.response.headers);
                    alert("Invalid credentials. Try again.");
                }
            });
        event.preventDefault();
    }
    function handleUsernameChange(event) {
        setUsername(event.target.value);
    }
    function handlePasswordChange(event) {
        setPassword(event.target.value);
    }
    return (
        <div>
            <h1>Admin Sign In</h1>
            <h5>If you are not an admin, please head to the home page and register.</h5>
            <form>
                <fieldset>
                    <label>Admin ID</label>
                    <input type="text" onChange={handleUsernameChange} value={username} required></input>

                    <label>Password</label>
                    <input type="password" onChange={handlePasswordChange} value={password} required></input>

                    <input className="button" id="black-button" type="submit" value="Log in" onClick={authenticate}></input>
                </fieldset>
            </form>
        </div>
    )
}

export default Login;