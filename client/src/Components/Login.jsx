import axios from "axios";
import React, { useState } from "react";
import "./App.css";

function Login() {
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");

    function authenticate(event) {
        axios.post('http://localhost:8080/api/login', {
            username: username,
            password: password
        }, {headers:{}}, { crossdomain: true })
            .then (res => {
                console.log(res);
                console.log(res.data);
            })
            .catch(function (error) {
                if (error.response) {
                    console.log(error.response.data);
                    console.log(error.response.status);
                    console.log(error.response.headers);
                    alert("Invalid credentials. Try again.");
                    return;
                }
            });
        event.preventDefault();
    }
    function handleUsernameChange(event) {
        setUsername(event.target.value);
        console.log(event.target.value);
    }
    function handlePasswordChange(event) {
        setPassword(event.target.value);
        console.log(event.target.value);
    }
    return (
        <div>
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