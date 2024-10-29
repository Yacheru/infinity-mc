export interface IPaymentsService {
    createPayment(price: string, email: string, service: string, nickname: string, duration: string): any;
}

export interface IPunishmentsService {
    getPunishments(limit: string, type: string): any;
}