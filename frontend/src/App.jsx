import Header from './components/header/header'
import Main from './components/main/main'
import Aside from './components/aside/aside'
import Footer from './components/footer/footer'

import './App.css'

function App() {
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

export default App
