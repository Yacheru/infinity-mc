import {AxiosResponse} from "axios";
import {Response} from "@api/axios/entities/Response.ts";

export interface IPaymentsService {
    createPayment(price: string, email: string, service: string, nickname: string, duration: string): any;
}

export interface IPunishmentsService {
    getPunishments(limit: string, type: string): any;
}

export interface IAuthService {
    sendCode(email: string): Promise<void>
    register(nickname: string, email: string, password: string, code: string): Promise<AxiosResponse<Response>>
    login(email: string, password: string): Promise<AxiosResponse<Response>>
    refresh(refreshToken: string): Promise<AxiosResponse<Response>>
}