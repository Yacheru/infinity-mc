import './header.css'
import '../../../pages.css'

export default function Header() {
    return (
        <header className="header bgc-1 flex b br20">
            <a className="header__project flex" href={'/'}>INFINITY-MC</a>
            <nav className='header__nav'>
                <ul className='header__items flex'>
                    <li className='header__item'>
                        <a className={'header__item-link donat'} href={'/'}>Донат</a>
                    </li>
                    |
                    <li className='header__item'>
                        <a className={'header__item-link map'} href={'https://map.infinity-mc.ru'} target={'_blank'}>Карта</a>
                    </li>
                    |
                    <li className='header__item'>
                        <a className={'header__item-link docs'} href={'https://docs.infinity-mc.ru'} target={'_blank'}>Документация</a>
                    </li>
                    |
                    <li className='header__item'>
                        <a className={'header__item-link punishments'} href={'/bans?category=bans'}>Наказания</a>
                    </li>
                    |
                    <li className='header__item'>
                        <a className={'header__item-link punishments'} href={'/stats'}>Статистика</a>
                    </li>
                </ul>
            </nav>
        </header>
    )
}