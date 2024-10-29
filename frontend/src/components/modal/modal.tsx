import { IModal } from '$types/modal'
import { IData } from '$types/data'

import React from 'react';
import data from '@config/data.json'

import Form from "./forms/forms.js";
import Description from "./blocks/descriptions.js";

import '@styles/components/modal/modal.css'

const typedData: IData = data as IData

export default function Modal({active, setActive, modalType, item}: IModal): React.ReactElement {
    const Component = modalType === 'about' ? Description : Form
    const title: string = typedData[item].title

    let modalContent = (
        <div>
            <header className={`modal__header ${item}`}>{ title }</header>
            <hr className='modal__hr' />
            <main className='modal__main'>
                <Component item={item}/>
            </main>
        </div>
    );

    return (
        <div className={active ? "modal flex center active" : "modal flex center"} onClick={() => setActive(false)}>
            <div className={active ? "modal__content bgc-2 b active" : "modal__content b"} onClick={e => e.stopPropagation()}>
                {modalContent}
            </div>
        </div>
    )
}