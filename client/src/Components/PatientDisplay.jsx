import React, { useState, useEffect } from "react";
import axios from "axios";
import "./App.css";
import PatientInfo from "./PatientInfo";

function PatientDisplay() {
    const [patients, setPatients] = useState([]);

    useEffect(() => {
        axios.get('http://localhost:8080/api/patients', { crossdomain: true }).then(response => {
            setPatients(response.data);
            console.log(patients);
        });
    }, []);
    
    return (
        <div>
            <h1>Patients:</h1>
            <table>
                <thead>
                    <tr>
                    <th>Name</th>
                    <th>DOB</th>
                    <th>Phone</th>
                    <th>Email</th>
                    <th>Address</th>
                    <th>Time Registered</th>
                    </tr>
                </thead>
                {/* {patients.map((result, index) => {
                    <div>hi</div>
                })} */}
                <tbody>
                {patients.map((patient, index) => {
                    return (
                        <tr>
                            <td>{patient.name}</td>
                            <td>{patient.dob}</td>
                            <td>{patient.phone}</td>
                            <td>{patient.email}</td>
                            <td>{patient.address}</td>
                            <td>{patient.time}</td>
                        </tr>
                    )
                })
                }
                </tbody>
                
            </table>
        </div>
    )
}

export default PatientDisplay;