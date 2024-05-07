import React from 'react';
import data from '../../../data.json'

import Form from "./forms/forms.jsx";
import Description from "./blocks/descriptions.jsx";

import './modal.css'

export default function Modal({active, setActive, modalType, item}) {
    const Component = modalType === 'about' ? Description : Form
    let modalContent = '';
    
    modalContent = (
        <div>
            <header className={`modal__header ${item}`}>{ data[item].title }</header>
            <hr className='modal__hr' size='3' color='whitesmoke'/>
            <main className='modal__main'>
                <Component item={item}/>
            </main>
        </div>);

    return (
        <div className={active ? "modal flex center active" : "modal flex center"} onClick={() => setActive(false)}>
            <div className={active ? "modal__content active" : "modal__content"} onClick={e => e.stopPropagation()}>
                {modalContent}
            </div>
        </div>
    )
}