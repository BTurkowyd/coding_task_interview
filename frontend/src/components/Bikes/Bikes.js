import React, { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { fetchBikes, logoutUser } from "../../api/api";


const Bikes = () => {

    const navigate = useNavigate()
    const [bikes, setBikes] = useState([])
    const [user, setUser] = useState("")

    useEffect(() => {
        setUser({
            "name": localStorage.getItem("name"),
            "renting": localStorage.getItem("renting"),
            "bike_id:": localStorage.getItem("bike_id")
        })
    }, [])

    useEffect(() => {
        fetchBikes()
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
        .then(console.log("logged out"))
        .catch(error => console.log(error))
        navigate("/login")
        localStorage.clear()
        
    }

    const displayBikes = () => {
        return (bikes.map(item => (
            <div key={item.bike_id}>
                <p>
                    {item.name}
                    {item.user_id === user["name"] && <button>Return bike</button>}
                    {(item.rented === false && user["renting"] === "false")&& <button>Rent bike</button>}
                </p>
            </div>
        )))
    }

    return (
        <div>
            Logged in User: {user["name"]}
            <h3>List of all bikes</h3>
            {displayBikes()}
            <button onClick={handleLogout}>Logout</button>
        </div>
    )
}

export default Bikes