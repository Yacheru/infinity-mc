import React, { lazy, Suspense } from "react";
import { Route, Routes } from "react-router-dom";

import Loading from './lazyLoad.js'

import './styles/load.css'
import './styles/pages/pages.css'
import './styles/pages/root/root.css'
import './styles/container.css'

const Root = lazy(() => import('./pages/root/Root.tsx'))
const Term = lazy(() => import('./pages/terms/Terms.tsx'))
const News = lazy(() => import('./pages/news/News.tsx'))
const Login = lazy(() => import('./pages/login/Login.tsx'))
const Punishments = lazy(() => import('./pages/punishments/Punishments.tsx'))

export default function Router() {
    return (
        <Suspense fallback={<Loading />}>
            <Routes>
                <Route path='*' element={<Root />}/>
                <Route path='/terms' element={<Term />} />
                <Route path='/p' element={<Punishments />} />
                <Route path='/news' element={<News />} />
                <Route path='/login' element={<Login />} />
            </Routes>
        </Suspense>
    );
}