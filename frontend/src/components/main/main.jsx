import './main.css'
import '../container.css'

import Button from '../button/button'

export default function Main() {
    return (
        <main className="main">
            <section className="main__items">
                <article className="main__item">
                    <img className='main__item-image' src="./donats/png/Tower.png"/>
                    <div className="main__item-skeleton">
                        <p className="main__item-title">Хронон</p>
                        <Button item={'Хронон'}/>
                    </div>
                </article>

                <article className="main__item">
                    <img className='main__item-image' src="./donats/png/Portal.png"/>
                    <div className="main__item-skeleton">
                        <div className="main__item-title">Никнейм</div>
                        <Button item={'Никнейм'}/>
                    </div>
                </article>

                <article className="main__item">
                    <img className='main__item-image' src="./donats/png/Lake.png"/>
                    <div className="main__item-skeleton">
                        <div className="main__item-title">Значок</div>
                        <Button item={'Значок'}/>
                    </div>
                </article>

            </section>
      </main>
    )
}