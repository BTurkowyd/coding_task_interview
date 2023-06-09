import React, { useEffect, useState } from "react";
import { fetchBikeByID } from "../../api/api";
import { useParams } from "react-router-dom";

const Bike = () => {

    const [bike, setBike] = useState()

    const params = useParams()
    const {id} = params

    useEffect(() => {
        fetchBikeByID(id)
        .then(response => setBike(response.data))
        .catch(response => console.log(response.message))
    }, [])

    return(
        <div> {bike && (
            <>
            <p>Bike name: {bike.name}</p>
            <p>Bike id: {bike.bike_id}</p>
            <p>Bike latitude: {bike.latitude}</p>
            <p>Bike longtitude: {bike.longtitude}</p></>
        )}

        </div>
    )
}

export default Bike