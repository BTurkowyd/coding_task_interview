import React, { useEffect, useState } from "react";
import { fetchBikeByID, returnBike } from "../../api/api";
import { useParams } from "react-router-dom";

const Return = () => {

    const params = useParams()
    const {id} = params

    useEffect(() => {
        returnBike(id)
        .then(response => console.log(response.data))
        .catch(response => console.log(response.message))
    }, [])

    return(
        <div> {id}

        </div>
    )
}

export default Return