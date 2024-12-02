import { api } from "../instance.js";
import { apis } from '@config/config.json'
import { AxiosResponse } from "axios";
import { AdminResponse } from "@api/axios/entities/Response.ts";

export default class AdminService {
    static async getUsers(): Promise<AxiosResponse<AdminResponse>> {
        return api.get<AdminResponse>(
            `${apis.auth}/users`,
        )
    }

    static async updateRole(id: string, role: string): Promise<AxiosResponse<AdminResponse>> {
        return api.patch<AdminResponse>(
            `${apis.auth}/update-role/${id}?role=${role}`
        )
    }

    static async deleteUser(id: string): Promise<AxiosResponse<AdminResponse>> {
        return api.delete<AdminResponse>(
            `${apis.auth}/delete-user/${id}`,
        )
    }
}