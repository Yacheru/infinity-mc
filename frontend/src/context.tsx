import { createContext } from 'react';
import Auth from "./store/auth.ts";
import { IStore } from "$types/store";

export const auth = new Auth();

export const Context = createContext<IStore>({
    auth
});