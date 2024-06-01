import {useEffect, useState} from "react";

import BansBody from "./bansBody.jsx";
import NoBans from "./noBans.jsx";
import { useLocation } from "react-router-dom";

import axios from "axios";

import './banlist.css'


export default function Banlist() {
    const [punishments, setPunishments] = useState(null);
    const location = useLocation()
    let value = location.search.split('=')[1]

    useEffect(() => {
        axios.get(`http://localhost/v1/mc/punishments?limit=10&type=${value}`)
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
                {punishments ? <BansBody punishments={punishments.data}/> : <NoBans location={location} />}
            </table>
        </main>
    )
}