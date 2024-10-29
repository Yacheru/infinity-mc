import React, { useEffect, useState } from "react";

import '@styles/components/header/header.css'


export default function Header() {
    const [selected, setSelected] = useState('/')
    // const [authenticated, setAuthenticated] = useState(false)
    let value = location.pathname

    useEffect(() => {
        setSelected(value)

        // api request
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
                    <li className={`header__item ${selected === '/news' ? 'header-selected' : ''}`}>
                        <a className={'header__item-link news'} href={'/news'}>Новости</a>
                    </li>
                </ul>
            </nav>
            <div className='header__auth'>
                <ul className='header__auth__list'>
                    <li className='header__auth-item header-login'>
                        <a className='header__auth-item-link' href="/login">Войти</a>
                    </li>
                </ul>
            </div>
        </header>
    )
}
