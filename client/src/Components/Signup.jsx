import './App.css';
import React, { useState } from 'react';
import axios from 'axios';

function Signup() {

    const [name, setName] = useState("")
    const [dob, setDob] = useState(new Date())
    const [phone, setPhone] = useState("")
    const [email, setEmail] = useState("")
    const [address, setAddress] = useState("")

    function handleNameChange(event) {
        setName(event.target.value)
        console.log(name);
    }
    function handleDobChange(event) {
        setDob(event.target.value)
        console.log(dob);
    }
    function handlePhoneChange(event) {
        setPhone(event.target.value)
        console.log(phone);
    }
    function handleEmailChange(event) {
        setEmail(event.target.value)
        console.log(email);
    }
    function handleAddressChange(event) {
        setAddress(event.target.value)
        console.log(address);
    }
    function submitData(event) {
        if(name === "" || dob === "" || phone === "" || email === "" || address === "") {
            return;
        }

        // send post request to API
        axios.post('http://localhost:8080/api/patients', {
            name: name,
            dob: dob,
            phone: phone,
            email: email,
            address: address
        }, {headers:{}}, { crossdomain: true })
            .then (res => {
                console.log(res);
                console.log(res.data);
            })
        alert("You are registered!");
    }

    return(
        <div>
            <form>
                <fieldset>
                    <label htmlFor="name">Name</label>
                    <input type="text" placeholder="John Doe" onChange={handleNameChange} value={name} required></input>

                    <label htmlFor="dob">Date of Birth</label>
                    <input type="date" onChange={handleDobChange} value={dob} required></input>

                    <label htmlFor="phone">Phone Number</label>
                    <input type="tel" onChange={handlePhoneChange} value={phone} required></input>

                    <label htmlFor="email">Email Address</label>
                    <input type="email" onChange={handleEmailChange} value={email} required></input>

                    <label htmlFor="address">Address</label>
                    <textarea rows="2" placeholder="Street Address" onChange={handleAddressChange} value={address} required></textarea>

                    <a href="/">
                        <input className="button button-outline" type="submit" value="Sign Up" onClick={submitData}></input>
                    </a>
                </fieldset>
            </form>
        </div>
    )
}

export default Signup;