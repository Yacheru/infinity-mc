import React from "react";
import { Routes, Route } from "react-router-dom";

import Root from "./pages/root/Root.jsx";
import Term from "./pages/terms/Terms.jsx";
import Bans from "./pages/bans/Bans.jsx";
import Stats from "./pages/stats/Stats.jsx";

import './pages/pages.css'

export default function App() {
    return (
        <Routes>
            <Route path='*' element={<Root />} />
            <Route path='/terms' element={<Term />} />
            <Route exact path='/bans' element={<Bans />} />
            <Route path='/stats' element={<Stats />} />
        </Routes>
    );
}