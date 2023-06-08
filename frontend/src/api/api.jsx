import axios from 'axios'

const url = 'http://localhost:3000'


export const fetchAllBikes = async () => await axios.get(`${url}/bikes`, {withCredentials: true})
export const loginUser = (credentials) => axios.post(`${url}/login`, credentials, {withCredentials: true})
export const logoutUser = () => axios.get(`${url}/logout`, {withCredentials: true})
export const registerUser = () => axios.post(`${url}/register`)
export const getCookie = () => axios.get(`${url}/getCookie`)
