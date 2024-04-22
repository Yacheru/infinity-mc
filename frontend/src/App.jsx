import Donate from "./components/donate/Donate.jsx";
import Term from "./components/terms/Term.jsx";
import React from "react";

import { Routes, Route } from "react-router-dom";

export default function App() {
    return (
        <Routes>
            <Route path='*' element={<Donate />} />
            <Route path='/terms' element={<Term />} />
        </Routes>
    );
}