import React, { useState } from "react";
// import "./App.css";
import Login from "./Login.jsx";
import PatientDisplay from "./PatientDisplay";

function Admin() {
    const [isAdmin, setIsAdmin] = useState(false);
    function displayPatients() {
        setIsAdmin(true);
        alert("Welcome Admin. Patient information is displayed below.")
    }
    return(
        <div className="container">
            {/* This will hide the login component when admin is validated */}
            <div hidden={(isAdmin) ? true : false}> 
                <Login 
                    displayPatients={displayPatients}
                />
            </div>

            {/* This will display all patients when admin is validated */}
            <div hidden={(isAdmin) ? false : true}>
                <PatientDisplay />
            </div>
        </div>
    )
}

export default Admin;