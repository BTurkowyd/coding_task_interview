import React, { useEffect } from "react";
import { logoutUser } from "../../api/api";
import { useNavigate } from "react-router-dom";


const Logout = () => {

    const navigate = useNavigate()

    useEffect(() => {
        logoutUser()
        .then(console.log("logged out"))
        .catch(error => console.log(error))
        navigate("/login")
        localStorage.clear()
    }, [])

}

export default Logout