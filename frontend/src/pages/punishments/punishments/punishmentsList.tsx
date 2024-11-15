import { useLocation } from "react-router-dom";
import { useEffect, useState } from "react";

import React from "react";

import BansBody from "./punishmentsItem.tsx";
import NoBans from "./no-punishments.tsx";
import Loading from '../../../lazyLoad.tsx';

import {IPunishmentsService} from "$types/api";
import PunishmentsService from "@api/axios/requests/punishments.ts";

import '@styles/pages/punishments/punishments.css';

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
                setPunishments(resp.data.data)
            } catch (error: any) {
                setStatus(error.response?.status)
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
            <div className={'punishment-box__table'}>
                <div className={'punishment-table__header'}>
                <div className={'punishment-header__tr flex'}>
                    <div className={'punishment-tr-item b bgc-2 br10 flex center'}>Нарушитель</div>
                    <div className={'punishment-tr-item b bgc-2 br10 flex center'}>Причина</div>
                    <div className={'punishment-tr-item b bgc-2 br10 flex center'}>Срок</div>
                    <div className={'punishment-tr-item b bgc-2 br10 flex center'}>Администратор</div>
                </div>
                </div>
                {
                    loading ? <Loading /> : status === 200 ? <BansBody punishments={punishments}/> : <NoBans location={location} status={status}/>
                }
            </div>
        </main>
    )
}
