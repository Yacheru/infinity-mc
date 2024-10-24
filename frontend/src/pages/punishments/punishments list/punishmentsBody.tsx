import React, {useState} from 'react'
import { calculateTimeLeft, calculateTimeDescription } from './calculateTime.tsx'

const PLAYER_HEAD_URL = 'https://visage.surgeplay.com/face/32/'
const CONSOLE_UUID = '00000000-0000-0000-0000-000000000000'


export default function PunishmentBody({ punishments }) {
    const [isHoveringDur, setIsHoveringDur] = useState(false);
    const [isHoveringNick, setIsHoveringNick] = useState(false)

    const handleMouseOverDur = () => {
        setIsHoveringDur(true);
    };

    const handleMouseOutDur = () => {
        setIsHoveringDur(false);
    };

    const handleMouseOverNick = () => {
        setIsHoveringNick(true)
    }

    const handleMouseOutNick = () => {
        setIsHoveringNick(false)
    }

    return (
        <tbody className={'punishment-table__body flex h100'}>
        {
            punishments.map((punish, index) => {
                const { victim, reason, time, operator } = punish;
                const victimImage = `${PLAYER_HEAD_URL}${victim['uuid']}`;
                const operatorUUID = operator['uuid'] !== CONSOLE_UUID ? operator['uuid'] : 'console';
                const operatorImage = `${PLAYER_HEAD_URL}${operatorUUID}`;

                return (
                    <tr className={'punishment-body__tr flex'} key={index}>
                        <td aria-label={victim['uuid']} onMouseOut={handleMouseOutNick} onMouseEnter={handleMouseOverNick} className={'punishment-tr-item flex center'}>
                            {isHoveringNick &&
                                <div className={`description-box bgc-2 br10 b`}>
                                    {
                                        victim['uuid']
                                    }
                                </div>
                            }

                            <img className={'player-head-left'} src={victimImage} alt={victim['name']} />
                            {
                                victim['name']
                            }
                        </td>
                        <td className={'punishment-tr-item flex center'}>{reason}</td>
                        <td className={'punishment-tr-item flex center'}>
                            {isHoveringDur &&
                                <div className={`description-box bgc-2 br10 b`}>
                                    {
                                        calculateTimeDescription(time['start'], time['end'])
                                    }
                                </div>
                            }
                            <a onMouseOver={handleMouseOverDur} onMouseOut={handleMouseOutDur}
                               className={'punishment-tr-item__time flex center h100 w100'}>
                                {
                                    calculateTimeLeft(time['end'])
                                }
                            </a>
                        </td>
                        <td className={'punishment-tr-item flex center reverse'}>
                            <img className={'player-head-right'} src={operatorImage} alt={'?'} />
                            {operator['name']}
                        </td>
                    </tr>
                )
            })
        }
        </tbody>
    );
}
