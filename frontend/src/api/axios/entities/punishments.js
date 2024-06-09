import { api } from "../instance.js";

export default class PunishmentsService {
    getPunishments(limit, type) {
        return api.get(`/mc/punishments?limit=${limit}&type=${type}`)
    }
}