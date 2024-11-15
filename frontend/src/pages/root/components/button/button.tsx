import { IButton } from "$types/root";
import { IData } from "$types/data";

import React, { useState } from 'react'

import Modal from '@components/modal/modal.js'
import Description from "@components/modal/blocks/descriptions.tsx";
import Form from "@components/modal/forms/forms.tsx";

import data from '@config/data.json'


import '@styles/pages/root/button.css'


const typedData: IData = data as IData

export default function Button({ item }: IButton) {
    const [modalActive, setModalActive] = useState<boolean>(false)
    const [modalType, setModalType] = useState<string>('')

    const Component = modalType === 'about' ? Description : Form
    const title: string = typedData[item].title

    let children = (
        <>
            <header className={`modal__header ${item}`}>{ title }</header>
            <hr className='modal__hr' />
            <main className='modal__main'>
                <Component item={item}/>
            </main>
        </>
    );

    return (
        <div className="main__item-navbar flex">
            <button className='main__item-button about' onClick={() => {setModalActive(true); setModalType('about')}}>?</button>
            <button className='main__item-button buy' onClick={() => {setModalActive(true); setModalType('buy')}}>{ typedData[item].costPlaceholder }</button>
            <Modal active={modalActive} setActive={setModalActive}>
                { children }
            </Modal>
        </div>
    )
}