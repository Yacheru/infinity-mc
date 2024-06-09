import axios from "axios";
import cfg from '../../../config.json'

export const api = axios.create({
    baseURL: cfg.apiUrl
})