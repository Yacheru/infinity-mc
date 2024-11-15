import {useContext, useState} from "react";
import { Context } from "../../context";

export default function Profile({ nickname }: {nickname: string}) {
    const [open, setOpen] = useState(false);

    const { auth } = useContext(Context);

    return (
        <ul className='profile-list' onClick={() => setOpen(!open)}>
            <li className='profile-list-item'>
                <div className='profile-list-item-link profile t-center'>
                    {nickname}
                </div>
                <ul className={`dropdown-list t-center${open ? ' open bgc-2 b br10' : ''}`}>
                    <li className='dropdown-list-item'>
                        <a className='dropdown-list-item-link t-center' href={`/profile/${nickname}`}>
                            Профиль
                        </a>
                    </li>
                    <li className='dropdown-list-item'>
                        <a className='dropdown-list-item-link t-center' href={`/profile/${nickname}/settings`}>
                            Настройки
                        </a>
                    </li>
                    {
                        auth.isAdmin ? <li className='dropdown-list-item'>
                            <a className='dropdown-list-item-link t-center' href={`/admin`}>
                                Управление
                            </a>
                        </li> : ''
                    }
                    <li className='dropdown-list-item'>
                        <div className='dropdown-list-item-link t-center' onClick={() => auth.logout()}>
                            Выйти
                        </div>
                    </li>
                </ul>
            </li>
        </ul>
    )
}