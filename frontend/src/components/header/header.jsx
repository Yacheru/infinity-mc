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
                        <a className={'header__item-link map'} href={'https://map.infinity-mc.ru'}
                           target={'_blank'}>Карта</a>
                    </li>
                    |
                    <li className={`header__item`}>
                        <a className={'header__item-link docs'} href={'https://docs.infinity-mc.ru'}
                           target={'_blank'}>Документация</a>
                    </li>
                    |
                    <li className={`header__item ${selected === '/punishments' ? 'header-selected' : ''}`}>
                        <a className={'header__item-link punishments'} href={'/punishments?category=bans'}>Наказания</a>
                    </li>
                    |
                    <li className={`header__item ${selected === '/stats' ? 'header-selected' : ''}`}>
                        <a className={'header__item-link punishments'} href={'/stats'}>Статистика</a>
                    </li>
                </ul>
            </nav>
        </header>
    )
}