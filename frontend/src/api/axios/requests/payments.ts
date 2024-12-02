import { IPaymentsService } from "$types/api";
import { apis } from '@config/config.json'
import { api } from "../instance.js";

export default class PaymentsService implements IPaymentsService {
    async createPayment(price: string, email: string, service: string, nickname: string, duration: string) {
        return api.post(
            `${apis.pay}/create`,
            {
                price: price,
                email: email,
                service: service,
                nickname: nickname,
                duration: Number(duration),
            }
        )
    }
}