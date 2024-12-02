import {AuthResponse, IUser} from "@api/axios/entities/authResponse.ts";

export interface Response {
    status: number;
    message: string;
    data: AuthResponse;
}

export interface AdminResponse {
    status: number;
    message: string;
    data: IUser[];
}
