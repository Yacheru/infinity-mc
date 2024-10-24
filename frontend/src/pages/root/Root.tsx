import Header from '../../components/header/header.js'
import Main from './components/main/main.js'
import Aside from './components/aside/aside.js'
import Footer from './components/footer/footer.js'

import './Root.css'
import React from 'react'

export default function Root() {
    return (
        <div className='container'>
            <Header/>
            <div id='wrapper' className={'flex'}>
                <Aside/>
                <Main/>
            </div>
            <Footer/>
        </div>
    )
}
