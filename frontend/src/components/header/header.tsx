import React, {useContext, useEffect, useState} from "react";
import { Context } from "../../context";

import { mc } from "@config/config.json"

import Auth from "@components/header/b-auth.tsx";
import Profile from "@components/header/b-profile.tsx";

import '@styles/components/header/header.css'

export default function Header() {
    const [selected, setSelected] = useState('/')
    const [open, setOpen] = useState(false);
    let value = location.pathname

    const { auth } = useContext(Context);

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
                        <a className={'header__item-link map'} href={mc.map}>Карта</a>
                    </li>
                    |
                    <li className={`header__item`}>
                        <a className={'header__item-link docs'} href={mc.forum}>Форум</a>
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
            <div className='header__button'>
                <div onClick={() => (setOpen(!open))} className={`header__button-item t-center b bgc-2 w100 h100 ${open && auth.isAuth ? 'open': ''}`}>
                    {auth.isAuth ? <Profile nickname={auth.user.nickname}/> : <Auth />}
                </div>
            </div>
        </header>
)}
