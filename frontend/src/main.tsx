import * as React from "react";
import * as ReactDOM from "react-dom/client";
import { BrowserRouter } from "react-router-dom";

import App from "./App.js";
import Load from "./Load.tsx";

import '../18n.js'
import '../faro.js'

import "./index.css";

ReactDOM.createRoot(document.getElementById("root")).render(
    <BrowserRouter>
        <React.Suspense fallback={<Load />}>
            <App />
        </React.Suspense>
    </BrowserRouter>
);