import { IPunishmentsService } from "$types/api";
import { api } from "../instance.js";
import { apis } from '@config/config.json'

export default class PunishmentsService implements IPunishmentsService {
    async getPunishments(limit: string, type: string) {
        return api.get(`${apis.punishments}/?limit=${limit}&type=${type}`)
    }
}