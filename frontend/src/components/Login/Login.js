import React, { useState } from "react";
import { loginUser } from "../../api/api";
import { useNavigate } from "react-router-dom";


const Login = () => {

    const [name, setName] = useState("")
    const [password, setPassword] = useState("")

    const navigate = useNavigate()

    const handleSubmit = () => {
        const credentials = {
            "name": name,
            "password": password
        }
        try {
            const response = loginUser(JSON.stringify(credentials))
            console.log(response)
        } catch (error) {
            console.log(error)
        }

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
        </div>
    )
}

export default Login