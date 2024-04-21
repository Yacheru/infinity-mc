import './header.css'
import '../container.css'

export default function Header() {
    return (
        <header className="header">
            <a className="header__project" href="">INFINITY-MC</a>
            <nav className='header__nav'>
                <ul className='header__items'>
                    <li className='header__item'>
                        <a className='header__item-link' href="#">Донат</a>
                    </li>
                    |
                    <li className='header__item'>
                        <a className='header__item-link' href="https://map.infinity-mc.ru" target='_blank'>Карта</a>
                    </li>
                    |
                    <li className='header__item'>
                        <a className='header__item-link' href="#">Документация</a>
                    </li>
                </ul>
            </nav>
        </header>
    )
}