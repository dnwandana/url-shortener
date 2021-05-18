import axios, { AxiosError, AxiosInstance, AxiosResponse } from "axios"

const axiosInstance: AxiosInstance = axios.create({
  baseURL: process.env.NEXT_PUBLIC_API_ENDPOINT,
  withCredentials: true,
})

axiosInstance.interceptors.response.use(
  (response: AxiosResponse) => {
    return response
  },
  (error: AxiosError) => {
    return Promise.reject(error)
  }
)

export default axiosInstance
