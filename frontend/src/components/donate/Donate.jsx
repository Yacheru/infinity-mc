import Header from './header/header.jsx'
import Main from './main/main.jsx'
import Aside from './aside/aside.jsx'
import Footer from './footer/footer.jsx'

import './Donate.css'

export default function Donate() {
  return (
      <div className='container'>
          <Header/>
          <div id='wrapper'>
              <Aside/>
              <Main/>
          </div>
          <Footer/>
      </div>
  )
}
