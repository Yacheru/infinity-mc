import { IUser } from "@api/axios/entities/authResponse.ts";
import { Response } from "@api/axios/entities/Response.ts";
import AuthService from "@api/axios/requests/auth.ts";
import { apis } from '@config/config.json'

import { Dispatch, SetStateAction } from "react";
import { makeAutoObservable } from "mobx";
import axios from "axios";

export default class Auth {
    user = {} as IUser;
    isAuth = false
    isAdmin = false

    constructor() {
        makeAutoObservable(this);
    }

    setAuth(bool: boolean) {
        this.isAuth = bool;
    }

    setUser(user: IUser) {
        this.user = user;
    }

    setAdmin(bool: boolean) {
        this.isAdmin = bool;
    }

    async codeSend(email: string, setModalActive: Dispatch<SetStateAction<boolean>>) {
        try {
            await AuthService.sendCode(email);
            setModalActive(true)
        } catch (e: any) {
            console.error(e.response?.data?.message);
        }
    }

    async login(email: string, password: string) {
        try {
            const response = await AuthService.login(email, password);
            localStorage.setItem('token', response.data.data.tokens.access_token)
            this.setAuth(true);
            this.setUser(response.data.data.user);
            window.location.assign(`/`);
        } catch (e: any) {
            console.log(e.response?.data?.message);
        }
    }

    async register(nickname: string, email: string, password: string, code: string) {
        try {
            const response = await AuthService.register(nickname, email, password, code)
            localStorage.setItem('token', response.data.data.tokens.access_token)
            this.setAuth(true);
            this.setUser(response.data.data.user);
            window.location.assign(`/`);
        } catch (e: any) {
            console.log(e.response?.data?.message);
        }
    }

    async logout() {
        try {
            await AuthService.logout();
            localStorage.removeItem('token');
            this.setAuth(false);
            this.setUser({} as IUser);
            window.location.reload();
        } catch (e: any) {
            console.log(e.response?.data?.message);
        }
    }

    async checkAuth() {
        try {
            const response = await axios.post<Response>(`${apis.auth}/refresh`, {}, { withCredentials: true })
            localStorage.setItem('token', response.data.data.tokens.access_token)
            this.setAuth(true);
            this.setAdmin(response.data.data.user.role === 'admin');
            this.setUser(response.data.data.user);
        } catch (e: any) {
            console.log(e.response?.data?.message);
        }
    }
}