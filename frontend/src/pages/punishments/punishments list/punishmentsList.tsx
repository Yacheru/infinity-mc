import { useLocation } from "react-router-dom";
import { useEffect, useState } from "react";

import React from "react";

import BansBody from "./punishmentsBody.tsx";
import NoBans from "./noPunishments.tsx";
import Loading from '../../../lazyLoad.tsx';

import '@styles/pages/punishments/punishments.css';
import {IPunishmentsService} from "$types/api";
import PunishmentsService from "@api/axios/entities/punishments.ts";

const LIMIT = "10"

const pService: IPunishmentsService = new PunishmentsService();

export default function PunishmentsList() {
    const [loading, setLoading] = useState(true)
    const [status, setStatus] = useState<number | null>(null)
    const [punishments, setPunishments] = useState([])
    const type = useLocation().search.split('=')[1]

    useEffect(() => {
        const getPunishments = async () => {
            try {
                const resp = await pService.getPunishments(LIMIT, type)
                setStatus(resp.status)
                setPunishments(resp.data)
            } catch (error: any) {
                setStatus(error.response ? error.response.status : 500)
            } finally {
                setLoading(false)
            }
        }

        if (type === "bans" || type === "warns" || type === "mutes") {
            getPunishments().then()
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
                    loading ? <Loading /> : status === 200 ? <BansBody punishments={punishments}/> : <NoBans location={location} status={status}/>
                }
            </table>
        </main>
    )
}
