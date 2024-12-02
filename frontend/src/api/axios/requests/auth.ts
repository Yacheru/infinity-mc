import { api } from "../instance.js";
import { apis } from '@config/config.json'
import { AxiosResponse } from "axios";
import { Response } from "@api/axios/entities/Response.ts";
import { LogoutResponse } from "@api/axios/entities/authResponse.ts";

export default class AuthService {
    static async sendCode(email: string): Promise<void> {
        return api.post(
            `${apis.auth}/send-code`,
            {
                email: email
            },
        )
    }

    static async register(nickname: string, email: string, password: string, code: string): Promise<AxiosResponse<Response>> {
        return api.post<Response>(
            `${apis.auth}/register?code=${code}`,
            {
                nickname: nickname,
                email: email,
                password: password,
            }
        )
    }

    static async login(email: string, password: string): Promise<AxiosResponse<Response>> {
        return api.post<Response>(
            `${apis.auth}/login`,
            {
                email: email,
                password: password,
            }
        )
    }

    static async logout(): Promise<AxiosResponse<LogoutResponse>> {
        return api.get<LogoutResponse>(
            `${apis.auth}/logout`,
        )
    }

    static async getUsers(): Promise<AxiosResponse<Response>> {
        return api.get<Response>(
            `${apis.auth}/users`,
        )
    }
}