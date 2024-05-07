import './header.css'
import '../../../pages.css'

export default function Header() {
    return (
        <header className="header bgc-1 flex b br20">
            <a className="header__project flex" href={'/'}>INFINITY-MC</a>
            <nav className='header__nav'>
                <ul className='header__items flex'>
                    <li className='header__item'>
                        <a className={'header__item-link'} href={'/'}>Донат</a>
                    </li>
                    |
                    <li className='header__item'>
                        <a className={'header__item-link'} href={'https://map.infinity-mc.ru'} target={'_blank'}>Карта</a>
                    </li>
                    |
                    <li className='header__item'>
                        <a className={'header__item-link'} href={'https://docs.infinity-mc.ru'} target={'_blank'}>Документация</a>
                    </li>
                    |
                    <li className='header__item'>
                        <a className={'header__item-link'} href={'/bans'}>Наказания</a>
                    </li>
                </ul>
            </nav>
        </header>
    )
}