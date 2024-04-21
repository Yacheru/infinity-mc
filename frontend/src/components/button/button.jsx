import './button.css'
import Modal from '../modal/modal'
import { useState } from 'react'

export default function Button({ item }) {
    const [modalActive, setModalActive] = useState(false)
    const [modalType, setModalType] = useState('') 
    
    return (
        <div className="main__item-navbar">
            <button className='main__item-button about' onClick={() => {setModalActive(true); setModalType('about')}}>?</button>
            <button className='main__item-button buy' onClick={() => {setModalActive(true); setModalType('buy')}}>Купить</button>
            <Modal active={modalActive} setActive={setModalActive} modalType={modalType} item={item}/>
        </div>
    )
}