import {api} from "../../instance.js";

export const getPunishments = async (limit, type) => api.get(`/mc/punishments?limit=${limit}&type=${type}`);