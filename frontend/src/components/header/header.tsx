import { useEffect, useState } from "react";

import './header.css'

export default function Header() {
    const [selected, setSelected] = useState('/')
    let value = location.pathname

    useEffect(() => {
        setSelected(value)
    }, [value]);

    return (
        <header className='header bgc-1 flex b br20'>
            <a className="header__project flex" href={'/'}>INFINITY-MC</a>
            <nav className='header__nav'>
                <ul className='header__items flex'>
                    <li className={`header__item ${selected === '/' ? 'header-selected' : ''}`}>
                        <a className={'header__item-link donat'} href={'/'}>Донат</a>
                    </li>
                    |
                    <li className={`header__item`}>
                        <a className={'header__item-link map'} href={'https://map.infinity-mc.ru'}>Карта</a>
                    </li>
                    |
                    <li className={`header__item`}>
                        <a className={'header__item-link docs'} href={'https://forum.infinity-mc.ru'}>Форум</a>
                    </li>
                    |
                    <li className={`header__item ${selected === '/p' ? 'header-selected' : ''}`}>
                        <a className={'header__item-link punishments'} href={'/p?category=bans'}>Наказания</a>
                    </li>
                    |
                    <li className={`header__item ${selected === '/stats' ? 'header-selected' : ''}`}>
                        <a className={'header__item-link stats'} href={'/stats'}>Статистика</a>
                    </li>
                    |
                    <li className={`header__item ${selected === '/news' ? 'header-selected' : ''}`}>
                        <a className={'header__item-link news'} href={'/news'}>Новости</a>
                    </li>
                </ul>
            </nav>
        </header>
    )
}
