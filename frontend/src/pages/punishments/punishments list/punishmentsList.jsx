import { useLocation } from "react-router-dom";
import { useEffect, useState } from "react";

import BansBody from "./punishmentsBody.jsx";
import NoBans from "./noPunishments.jsx";
import Loading from '../../../Load.jsx'

import * as axios from '../../../api/axios/requests';

import './punishments.css'

const LIMIT = 10

export default function Punishment() {
    const [loading, setLoading] = useState(true)
    const [status, setStatus] = useState(null)
    const [punishments, setPunishments] = useState([])
    const type = useLocation().search.split('=')[1]

    useEffect(() => {
        async function getPunishments() {
            try {
                const punishmentsResponse = await axios.getPunishments(LIMIT, type)
                setStatus(punishmentsResponse.status)
                setPunishments(punishmentsResponse.data)
            } catch (error) {
                setStatus(error.response ? error.response.status : 500)
            } finally {
                setLoading(false)
            }
        }

        if (type === "bans" || type === "warns" || type === "mutes") {
            getPunishments()
        } else {
            setLoading(false)
        }
    }, [type]);

    return (
        <main className={'punishment-box b bgc-1 br20'}>
            <table className={'punishment-box__table'}>
                <thead className={'punishment-table__header'}>
                <tr className={'punishment-header__tr flex'}>
                    <td className={'punishment-tr-item b bgc-2 br10 flex center'}>Нарушитель</td>
                    <td className={'punishment-tr-item b bgc-2 br10 flex center'}>Причина</td>
                    <td className={'punishment-tr-item b bgc-2 br10 flex center'}>Срок</td>
                    <td className={'punishment-tr-item b bgc-2 br10 flex center'}>Администратор</td>
                </tr>
                </thead>
                {
                    loading ? <Loading /> :
                        status === 200 ? <BansBody punishments={punishments}/> : <NoBans location={location} status={status}/>
                }
            </table>
        </main>
    )
}
