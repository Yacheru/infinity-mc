import React from 'react';
import './modal.css'
import data from '../../../data.json'

export default function Modal({active, setActive, modalType, item}) {
    let modalContent = '';
    
    if (modalType === 'about') {
        modalContent = (
        <div>
            <header className='modal__header'>
                Инфо { item }
            </header>
            <hr className='modal__hr' size='3' color='whitesmoke'/>
            <main className='modal__main'>
                <p className='modal__description'>
                    {data[item].split('\n').map((line, index) => (
                        <React.Fragment key={index}>
                            {line}
                            <br/>
                        </React.Fragment>
                    ))}
                </p>
            </main>
        </div>);
    } else if (modalType === 'buy') {
        modalContent = <div>Купить { item }</div>;
    }
    
    return (
        <div className={active ? "modal active" : "modal"} onClick={() => setActive(false)}>
            <div className={active ? "modal__content active" : "modal__content"} onClick={e => e.stopPropagation()}>
                {modalContent}
            </div>
        </div>
    )
}