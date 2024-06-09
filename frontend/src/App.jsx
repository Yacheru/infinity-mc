import React, {lazy, Suspense} from "react";
import { Routes, Route } from "react-router-dom";

import Loading from './Load.jsx'

import './load.css'
import './pages/pages.css'
import './pages/root/Root.css'
import './pages/root/container.css'

const Root = lazy(() => import('./pages/root/Root.jsx'))
const Term = lazy(() => import('./pages/terms/Terms.jsx'))
const Punishments = lazy(() => import('./pages/punishments/Punishments.jsx'))
const Stats = lazy(() => import('./pages/stats/Stats.jsx'))

export default function App() {
    return (
        <Routes>
            <Route
                path='*'
                element={
                    <Suspense fallback={<Loading />}>
                        <Root />
                    </Suspense>
                }
            />
            <Route
                path='/terms'
                element={
                    <Suspense fallback={<Loading />}>
                        <Term />
                    </Suspense>
                }
            />
            <Route
                exact
                path='/punishments'
                element={
                    <Suspense fallback={<Loading />}>
                        <Punishments />
                    </Suspense>
                }
            />
            <Route
                path='/stats'
                element={
                    <Suspense fallback={<Loading />}>
                        <Stats />
                    </Suspense>
                }
            />
        </Routes>
    );
}