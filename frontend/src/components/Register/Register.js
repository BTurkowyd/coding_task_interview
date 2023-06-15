import React, { useState, useEffect } from "react";
import { registerUser } from "../../api/api";
import { useNavigate } from "react-router-dom";
import './style.css'


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
                .then(response => {
                    console.log(response.data)
                    if (response.data.code === 200) {
                        navigate('/login')
                    }
                })
                .catch(error => {
                    if (error.response) {
                        console.log(error.response.data);
                    } else {
                        console.error(error);
                    }
                })
        } else {
            console.log("Passwords in both fields do not match.")
        }

        setUserName("")
        setPassword("")
        setConfirmPassword("")
    }

    return (
        <div className="register-container">
            <form onSubmit={handleRegister}>
                <div className="register-heading">
                    Register
                </div>
                <div className="register-form">
                    <input className="register-input" placeholder="Login" type="text" onChange={(e) => setUserName(e.target.value)} value={userName} name="name" required />
                    <input className="register-input" placeholder="Password" type="password" onChange={(e) => setPassword(e.target.value)} value={password} name="password" required />
                    <input className="register-input" placeholder="Confirm password" type="password" onChange={(e) => setConfirmPassword(e.target.value)} value={confirmPassword} name="password" required />
                    <input className="register-button" type="submit" />
                </div>
                <div className="register-link">
                    Already have an account? <a href="/login"> Log in. </a>
                </div>
            </form>
        </div>
    )
}

export default Register