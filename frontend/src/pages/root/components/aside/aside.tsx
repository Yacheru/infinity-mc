import './aside.css'
import '../../container.css'
import React from 'react'

export default function Aside() {
    return (
        <aside className="aside h100">
            <div className="aside__text flex h100">
                <div className="aside__title">
                    <h1>Покупка услуг</h1>
                </div>
                <div className="aside__items">
                    <p>Покупая донат, вы не только наслаждаетесь игрой, но и активно способствуете <span className={'important'}>нашему развитию!</span></p>
                    <p><span className={'important'}>Благодарим</span> каждого игрока, который поддерживает нас финансово и помогает сделать наш сервер еще <span className={'important'}>лучше!</span></p>
                </div>
            </div>
        </aside>
    )
}