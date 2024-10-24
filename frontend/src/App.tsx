import React, {lazy, Suspense} from "react";
import { Route } from "react-router-dom";
import { FaroRoutes } from "@grafana/faro-react";

import Loading from './Load.js'

import './load.css'
import './pages/pages.css'
import './pages/root/Root.css'
import './pages/root/container.css'

const Root = lazy(() => import('./pages/root/Root.js'))
const Term = lazy(() => import('./pages/terms/Terms.js'))
const Punishments = lazy(() => import('./pages/punishments/Punishments.js'))
const Stats = lazy(() => import('./pages/stats/Stats.js'))
const News = lazy(() => import('./pages/news/News.js'))

export default function App() {
    return (
        <FaroRoutes>
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
                path='/p'
                element={
                    <Suspense fallback={<Loading />}>
                        <Punishments />
                    </Suspense>
                }
            />
            <Route
                path={'/news'}
                element={
                    <Suspense fallback={<Loading />}>
                        <News />
                    </Suspense>
                }
            />
            <Route
                path={'/news'}
                element={
                    <Suspense fallback={<Loading />}>
                        <News />
                    </Suspense>
                }
            />
        </FaroRoutes>
    );
}