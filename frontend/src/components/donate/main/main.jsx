import './main.css'
import '../container.css'

import Button from '../button/button.jsx'

export default function Main() {
    return (
        <main className="main">
            <section className="main__items">
                <article className="main__item">
                    <img className='main__item-image' src="./donats/png/Tower.png" alt="tower"/>
                    <div className="main__item-skeleton">
                        <p className="main__item-title">Хронон</p>
                        <Button item={'hronon'}/>
                    </div>
                </article>

                <article className="main__item">
                    <img className='main__item-image' src="./donats/png/Portal.png" alt="portal"/>
                    <div className="main__item-skeleton">
                        <div className="main__item-title">Никнейм</div>
                        <Button item={'nickname'}/>
                    </div>
                </article>

                <article className="main__item">
                    <img className='main__item-image' src="./donats/png/Lake.png" alt="Lake"/>
                    <div className="main__item-skeleton">
                        <div className="main__item-title">Значок</div>
                        <Button item={'badge'}/>
                    </div>
                </article>

            </section>
      </main>
    )
}