import axios from "axios";

import {useEffect, useState} from "react";

import './banlist.css'
import BansBody from "./bansBody.jsx";
import NoBans from "./noBans.jsx";

export default function Banlist() {
    const [punishments, setPunishments] = useState({});

    useEffect(() => {
        axios.get('http://localhost:8000/v1/mc/bans?limit=10')
            .then(result => {
                return setPunishments(result)
            });
    }, [])

    return (
        <main className={'banlist-box b bgc-1 br20'}>
            <table className={'banlist-box__table'}>
                <thead className={'banlist-table__header'}>
                    <tr className={'banlist-header__tr flex'}>
                        <td className={'banlist-tr-item b bgc-2 br10 flex center'}>Нарушитель</td>
                        <td className={'banlist-tr-item b bgc-2 br10 flex center'}>Причина</td>
                        <td className={'banlist-tr-item b bgc-2 br10 flex center'}>Срок</td>
                        <td className={'banlist-tr-item b bgc-2 br10 flex center'}>Администратор</td>
                    </tr>
                </thead>
                {punishments === 200 ? <BansBody punishments={punishments}/> : <NoBans />}
            </table>
        </main>
    )
}