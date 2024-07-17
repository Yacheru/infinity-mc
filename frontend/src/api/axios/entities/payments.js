import { api } from "../instance.js";

export default class PunishmentsService {
    createPayment(nickname, email, price, duration, item, dur) {
        return api.get(`/payments?nickname=${nickname}&email=${email}&price=${price[duration]}&donat=${item}&duration=${dur}`)
    }
}