import axios from 'axios'

const url = 'http://localhost:3000'

const axiosInstance = axios.create({
    withCredentials: true
 })

export const loginUser = async (credentials) => await axiosInstance.post(`${url}/login`, credentials)
export const logoutUser = async () => await axiosInstance.get(`${url}/logout`)
export const registerUser = () => axiosInstance.post(`${url}/register`)
export const fetchBikes = () => axiosInstance.get(`${url}/bikes`)
export const fetchBikeByID = (bike_id) => axiosInstance.get(`${url}/bikes/${bike_id}`)
export const rentBike = (bike_id) => axiosInstance.patch(`${url}/rent/${bike_id}`, null)
export const returnBike = (bike_id) => axiosInstance.patch(`${url}/return/${bike_id}`, null)
export const fetchUserData = async (user_id) => await axiosInstance.get(`${url}/fetchUserData`, user_id)