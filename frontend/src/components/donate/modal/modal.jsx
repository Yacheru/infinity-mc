import React from 'react';
import data from '../../../../data.json'

import Form from "./forms/forms.jsx";
import Description from "./components/descriptions.jsx";

import './modal.css'

export default function Modal({active, setActive, modalType, item}) {
    let modalContent = '';
    
    if (modalType === 'about') {
        modalContent = (
        <div>
            <header className={`modal__header ${item}`}>{ data[item].title }</header>
            <hr className='modal__hr' size='3' color='whitesmoke'/>
            <main className='modal__main'>
                <Description item={item}/>
            </main>
        </div>);
    } else if (modalType === 'buy') {
        modalContent = (
            <div>
                <header className={`modal__header ${item}`}>{ data[item].title }</header>
                <hr className='modal__hr' size='3' color='whitesmoke'/>
                <main className='modal__main'>
                    <Form item={item}/>
                </main>
            </div>
        )
    }

    return (
        <div className={active ? "modal active" : "modal"} onClick={() => setActive(false)}>
            <div className={active ? "modal__content active" : "modal__content"} onClick={e => e.stopPropagation()}>
                {modalContent}
            </div>
        </div>
    )
}