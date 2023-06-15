import React, { useState, useEffect } from "react";
import { registerUser } from "../../api/api";
import { useNavigate } from "react-router-dom";

const Register = () => {
    const [userName, setUserName] = useState("")
    const [password, setPassword] = useState("")
    const [confirmPassword, setConfirmPassword] = useState("")
    const navigate = useNavigate()

    useEffect(() => {
        if (localStorage.name) {
            navigate("/bikes")
        }
    }, [])


    const handleRegister = (e) => {
        e.preventDefault()
        if (password === confirmPassword) {
            const credentials = {
                "name": userName,
                "password": password
            }
            registerUser(JSON.stringify(credentials))
            .then(response => console.log(response.data))
            .catch(error => console.log(error.data))
        } else {
            console.log("Passwords in both fields do not match.")
        }

        setUserName("")
        setPassword("")
        setConfirmPassword("")
    }

    return (
        <div className="form">
        <form onSubmit={handleRegister}>
            <div>
                <label htmlFor="name"> User name: </label>
                <input type="text" onChange={(e) => setUserName(e.target.value)} value={userName} name="name" required />
            </div>
            <div>
                <label htmlFor="password"> Password: </label>
                <input type="password" onChange={(e) => setPassword(e.target.value)} value={password} name="password" required />
            </div>
            <div>
                <label htmlFor="confirmPassword"> Confirm: </label>
                <input type="password" onChange={(e) => setConfirmPassword(e.target.value)} value={confirmPassword} name="password" required />
            </div>
            <div>
                <input type="submit"/>
            </div>
        </form>
    </div>
    )
}

export default Register