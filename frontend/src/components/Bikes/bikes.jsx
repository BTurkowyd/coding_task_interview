import React, { useEffect } from "react";
import { fetchAllBikes } from "../../api/api";

const Bikes = () => {

    useEffect(() => {
        fetchAllBikes()
    }, [])

    return (
        <div>
            <h1>Kurwa w koncu</h1>
        </div>
    )
}

export default Bikes