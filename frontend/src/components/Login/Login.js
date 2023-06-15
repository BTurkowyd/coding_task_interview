import React, { useEffect, useState } from "react";
import { loginUser } from "../../api/api";
import { useNavigate } from "react-router-dom";
import './style.css'

const Login = () => {

    const [name, setName] = useState("")
    const [password, setPassword] = useState("")
    const navigate = useNavigate()

    useEffect(() => {
        if (localStorage.name) {
            navigate("/bikes")
        }
    }, [])

    const handleSubmit = () => {

        const credentials = {
            "name": name,
            "password": password
        }
        loginUser(JSON.stringify(credentials))
            .then(response => {
                localStorage.setItem("name", response.data.name)
                localStorage.setItem("renting", response.data.renting)
                localStorage.setItem("bike_id", response.data.bike_id)
            })
            .catch(error => {
                if (error.response) {
                    console.log(error.response.data);
                } else {
                    console.error(error);
                }
            })

        setName("")
        setPassword("")
        navigate("/bikes")

    }

    return (
        <div className="login-container">
            <form onSubmit={handleSubmit}>
                <div className="login-heading">
                    Login
                </div>
                <div className="login-form">
                    <input className="login-input" placeholder="Login" type="text" onChange={(e) => setName(e.target.value)} value={name} name="name" required />
                    <input className="login-input" placeholder="Password" type="password" onChange={(e) => setPassword(e.target.value)} value={password} name="password" required />
                    <input className="login-button" type="submit" />
                </div>
            </form>
            <div className="login-link">
                Are you new here? <a href="/register"> Register your account! </a>
            </div>
        </div>
    )
}

export default Login