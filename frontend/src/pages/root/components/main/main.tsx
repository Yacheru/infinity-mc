import React from 'react'

import Button from '../button/button.js'

import '@styles/pages/root/main.css'
import '@styles/pages/pages.css'

export default function Main() {
    return (
        <main className="main">
            <section className="main__items h100">
                <article className="main__item">
                    <img className='main__item-image br20 h100 w100' src="../../../../assets/root/Tower.webp" alt="tower"></img>
                    <p className="main__item-title">Хронон</p>
                    <Button item={'hronon'}/>
                </article>

                <article className="main__item">
                    <img className='main__item-image br20 h100 w100' src="../../../../assets/root/Tower.webp" alt="portal"/>
                    <div className="main__item-title">Никнейм</div>
                    <Button item={'nickname'}/>
                </article>

                <article className="main__item">
                    <img className='main__item-image br20 h100 w100' src="../../../../assets/root/Lake.webp" alt="lake"/>
                    <div className="main__item-title">Значок</div>
                    <Button item={'badge'}/>
                </article>

            </section>
        </main>
    )
}