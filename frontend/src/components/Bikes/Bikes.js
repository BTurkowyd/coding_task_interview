import React, { useEffect, useState } from "react";
import axios from "axios";
import { useNavigate } from "react-router-dom";
import { logoutUser } from "../../api/api";

const url = 'http://localhost:3000'

const Bikes = () => {

    const navigate = useNavigate()
    const [bikes, setBikes] = useState([])

    useEffect(() => {
        axios.get(`${url}/bikes`, { withCredentials: true })
            .then(response => setBikes(response.data))
            .catch(error => {
                if (error.response && error.response.status === 401) {
                    console.log("unauthorized access")
                    navigate("/login")
                } else {
                    console.log(error)
                }
            })
    }, [])

    const handleLogout = () => {
        logoutUser()
        navigate("/login")
        console.log("logged out")
    }

    const rentBike = (bike_id, user_id) => {
        
    }

    return (
        <div>
            <h1>List of all bikes</h1>
            {bikes.map(item => (
                <div key={item.bike_id}>
                    <p>
                        {item.name} {item.rented.toString()} {item.user_id}
                        <button>Rent</button>
                    </p>
                </div>
            ))}
            <button onClick={handleLogout}>Logout</button>
        </div>
    )
}

export default Bikes