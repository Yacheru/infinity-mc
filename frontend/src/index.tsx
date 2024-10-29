import * as React from "react";
import * as ReactDOM from "react-dom/client";
import { BrowserRouter } from "react-router-dom";
import { Toaster } from "solid-toast";

import Router from "./routes.tsx";

import '../i18next'

import "./styles/index.css";

ReactDOM.createRoot(document.getElementById("root") as HTMLElement).render(
    <BrowserRouter>
        <Router />
    </BrowserRouter>
);