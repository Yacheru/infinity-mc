import Modal from '../modal/modal.jsx'
import data from '../../../../data.json'

import { useState } from 'react'

import './button.css'


export default function Button({ item }) {
    const [modalActive, setModalActive] = useState(false)
    const [modalType, setModalType] = useState('') 
    
    return (
        <div className="main__item-navbar">
            <button className='main__item-button about' onClick={() => {setModalActive(true); setModalType('about')}}>?</button>
            <button className='main__item-button buy' onClick={() => {setModalActive(true); setModalType('buy')}}>{ data[item].cost }</button>
            <Modal active={modalActive} setActive={setModalActive} modalType={modalType} item={item}/>
        </div>
    )
}