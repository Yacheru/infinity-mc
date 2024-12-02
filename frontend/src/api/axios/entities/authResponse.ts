export interface AuthResponse {
    user: IUser,
    tokens: ITokens
}

export interface LogoutResponse {
    status: number;
    message: string;
}

export interface IUser {
    user_id: string;
    email: string;
    role: string;
    nickname: string;
    password: string;
    ip_addr: string;
}

export interface ITokens {
    access_token: string;
    refresh_token: string;
}
