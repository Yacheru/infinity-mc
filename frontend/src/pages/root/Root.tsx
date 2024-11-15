import React from 'react'

import { Toaster } from "react-hot-toast";

import Header from '@components/header/header.js'
import Main from './components/main/main.js'
import Aside from './components/aside/aside.js'
import Footer from './components/footer/footer.js'

import '@styles/pages/root/root.css'

export default function Root() {
    return (
        <div className='container'>
            <Header/>
            <div id='wrapper' className={'flex'}>
                <Aside/>
                <Main/>
                <Toaster position="top-right" reverseOrder={false} />
            </div>
            <Footer/>
        </div>
    )
}
