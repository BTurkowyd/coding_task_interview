import React, { useEffect } from "react";
import { logoutUser } from "../../api/api";


const Logout = () => {

    useEffect(() => {
        logoutUser()
        .then(console.log("logged out"))
        .catch(error => console.log(error))
        navigate("/login")
        localStorage.clear()
    }, [])

}

export default Logout