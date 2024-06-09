import * as React from "react";
import * as ReactDOM from "react-dom/client";
import { BrowserRouter } from "react-router-dom";

import App from "./App";

import "./index.css";

import '../18n.js'
import Load from "./Load.jsx";

ReactDOM.createRoot(document.getElementById("root")).render(
    <BrowserRouter>
        <React.Suspense fallback={<Load />}>
            <App />
        </React.Suspense>
    </BrowserRouter>
);