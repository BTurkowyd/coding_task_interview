import React, { useEffect, useState } from "react";
import { loginUser } from "../../api/api";
import { useNavigate } from "react-router-dom";


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

    return(
        <div className="form">
            <form onSubmit={handleSubmit}>
                <div>
                    <label htmlFor="name"> User name: </label>
                    <input type="text" onChange={(e) => setName(e.target.value)} value={name} name="name" required />
                </div>
                <div>
                    <label htmlFor="password"> Password: </label>
                    <input type="password" onChange={(e) => setPassword(e.target.value)} value={password} name="password" required />
                </div>
                <div>
                    <input type="submit"/>
                </div>
            </form>
            Are you new here? <a href="/register"> Register your account! </a>
        </div>
    )
}

export default Login