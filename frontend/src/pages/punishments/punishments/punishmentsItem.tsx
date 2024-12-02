import { calculateTimeLeft } from '../../../utils/calculateTime.tsx'
import { useContext } from "react";
import { Context } from "../../../context.tsx";

const PLAYER_HEAD_URL = 'https://visage.surgeplay.com/head/32/'
const CONSOLE_UUID = '00000000-0000-0000-0000-000000000000'

export default function BansBody({ punishments }: { punishments: any[] }) {
    const { auth } = useContext(Context)

    return (
        <div className={'punishment-table__body flex h100'}>
        {
            punishments.map((punish, index) => {
                const { victim, reason, time, operator } = punish;
                const victimUUID = victim.uuid !== CONSOLE_UUID ? victim.uuid : 'steve'
                const operatorUUID = operator.uuid !== CONSOLE_UUID ? operator.uuid : 'console';

                const victimImage = `${PLAYER_HEAD_URL}${victimUUID}?y=70`;
                const operatorImage = `${PLAYER_HEAD_URL}${operatorUUID}`;

                const expired = (new Date().getTime() / 1000) > time.end && time.end !== 0

                return (
                    <div className={`punishment-body__tr flex ${expired ? 'p-elapsed' : ''}`} key={index}>
                        <div aria-label={victim.uuid} className={'punishment-tr-item flex center'}>
                            <img className={'player-head-left'} src={victimImage} alt='?' />
                            { victim.name }
                        </div>
                        <div className={'punishment-tr-item flex center'}>{reason}</div>
                        <div className={'punishment-tr-item flex center'}>
                            <a className={'punishment-tr-item__time flex center h100 w100'}>
                                { calculateTimeLeft(time.end) }
                            </a>
                        </div>
                        <div className={'punishment-tr-item flex center reverse'}>
                            <img className={'player-head-right'} src={operatorImage} alt={'?'} />
                            {operator.name}
                            { auth.isAdmin ? <p className={'configure-punishment b bgc-2 br10'}></p> : '' }
                        </div>
                    </div>
                )
            })
        }
        </div>
    );
}
