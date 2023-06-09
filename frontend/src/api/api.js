import axios from 'axios'

const url = 'http://localhost:3000'

export const loginUser = async (credentials) => await axios.post(`${url}/login`, credentials, {withCredentials: true})
export const logoutUser = async () => await axios.get(`${url}/logout`, {withCredentials: true})
export const registerUser = () => axios.post(`${url}/register`)
export const fetchBikes = () => axios.get(`${url}/bikes`, { withCredentials: true })
export const fetchBikeByID = (bike_id) => axios.get(`${url}/bikes/${bike_id}`, { withCredentials: true })