import Button from '../button/button.jsx'

import './main.css'
import '../../../pages.css'

export default function Main() {
    return (
        <main className="main">
            <section className="main__items h100">
                <article className="main__item">
                    <img className='main__item-image br20 h100 w100' src="/root/Tower.webp" alt="Tower"></img>
                    <p className="main__item-title">Хронон</p>
                    <Button item={'hronon'}/>
                </article>

                <article className="main__item">
                    <img className='main__item-image br20 h100 w100' src="/root/Portal.webp" alt="Portal"/>
                    <div className="main__item-title">Никнейм</div>
                    <Button item={'nickname'}/>
                </article>

                <article className="main__item">
                    <img className='main__item-image br20 h100 w100' src="/root/Lake.webp" alt="Lake"/>
                    <div className="main__item-title">Значок</div>
                    <Button item={'badge'}/>
                </article>

            </section>
        </main>
    )
}