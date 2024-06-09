import Header from '../../components/header/header.jsx'
import Main from './components/main/main.jsx'
import Aside from './components/aside/aside.jsx'
import Footer from './components/footer/footer.jsx'

import './Root.css'

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
