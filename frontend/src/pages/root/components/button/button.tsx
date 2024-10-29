import { IButton } from "$types/root";
import { IData } from "$types/data";

import Modal from '@components/modal/modal.js'
import data from '@config/data.json'

import { useState } from 'react'

import '../../../../styles/pages/root/button.css'
import React from 'react'

const typedData: IData = data as IData

export default function Button({ item }: IButton) {
    const [modalActive, setModalActive] = useState<boolean>(false)
    const [modalType, setModalType] = useState<string>('')
    
    return (
        <div className="main__item-navbar flex">
            <button className='main__item-button about' onClick={() => {setModalActive(true); setModalType('about')}}>?</button>
            <button className='main__item-button buy' onClick={() => {setModalActive(true); setModalType('buy')}}>{ typedData[item].costPlaceholder }</button>
            <Modal active={modalActive} setActive={setModalActive} modalType={modalType} item={item}/>
        </div>
    )
}