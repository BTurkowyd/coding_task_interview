import axios from 'axios'

const url = 'http://localhost:3000'

export const loginUser = async (credentials) => await axios.post(`${url}/login`, credentials, {withCredentials: true})
export const logoutUser = async () => await axios.get(`${url}/logout`, {withCredentials: true})
export const registerUser = () => axios.post(`${url}/register`)
export const fetchBikes = () => axios.get(`${url}/bikes`, { withCredentials: true })
export const fetchBikeByID = (bike_id) => axios.get(`${url}/bikes/${bike_id}`, { withCredentials: true })
export const rentBike = (bike_id) => axios.patch(`${url}/rent/${bike_id}`, null, { withCredentials: true })
export const returnBike = (bike_id) => axios.patch(`${url}/return/${bike_id}`, null, { withCredentials: true })