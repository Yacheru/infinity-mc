import {api} from "../../instance.js";

export const createPayment = async (nickname, email, price, duration, item, dur) =>
    api.get(`/payments?nickname=${nickname}&email=${email}&price=${price[duration]}&donat=${item}&duration=${dur}`);