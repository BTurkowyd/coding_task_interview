import axios from 'axios'

const url = 'http://localhost:3000'

export const loginUser = async (credentials) => await axios.post(`${url}/login`, credentials, {withCredentials: true})
export const logoutUser = async () => await axios.get(`${url}/logout`, {withCredentials: true})
export const registerUser = () => axios.post(`${url}/register`)