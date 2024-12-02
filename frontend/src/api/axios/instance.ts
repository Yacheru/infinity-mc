import axios, { AxiosInstance } from "axios";
import { apis } from '@config/config.json'
import { Response } from "@api/axios/entities/Response.ts";

export const api: AxiosInstance = axios.create({
    withCredentials: true,
})

api.interceptors.request.use((config) => {
    config.headers.Authorization = `Bearer ${localStorage.getItem('token')}`
    return config;
})

api.interceptors.response.use((config) => {
    return config
}, async (err) => {
    const req = err.config
    if (err.response.status == 401 && err.config && !err.config._IsRetry) {
        req._IsRetry = true;
        try {
            const response = await axios.post<Response>(`${apis.auth}/refresh`, {}, { withCredentials: true })
            localStorage.setItem('token', response.data.data.tokens.access_token)
            return api.request(req)
        } catch (e: any) {
            console.log(e.response?.data?.message);
        }
    }
    if (err.response.status == 403) {
        return window.location.assign(`/`);
    }
    throw err
})
