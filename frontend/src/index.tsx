import * as ReactDOM from "react-dom/client";
import { BrowserRouter } from "react-router-dom";

import Router from "./routes.tsx";
import { Context, auth } from './context';

import '../i18next'
import "./styles/index.css";

ReactDOM.createRoot(document.getElementById("root") as HTMLElement).render(
    <Context.Provider value={{auth}}>
        <BrowserRouter>
            <Router />
        </BrowserRouter>
    </Context.Provider>
);