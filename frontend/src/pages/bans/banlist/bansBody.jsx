export default function BansBody({ punishments }) {
    return (
        <tbody className={'banlist-table__body flex'}>
            {punishments.map((punish, index) => (
                <tr className={'banlist-body__tr flex'} key={index}>
                    <td className={'banlist-tr-item flex center'}>
                        <img className={'player-head-left'} src={`https://visage.surgeplay.com/face/32/${punish['victim']['uuid']}`} />
                        {punish['victim']['name']}
                    </td>
                    <td className={'banlist-tr-item flex center'}>{punish['reason']}</td>
                    <td className={'banlist-tr-item flex center'}>
                        {punish['time']['end'] === 0 ?
                            <p className={'br10 forever'}>Навсегда</p>
                            : new Date / 1000 > punish['time']['end'] ? <p className={'br10 elapsed'}>Срок истек</p> :
                                `${Math.floor(((punish['time']['end'] - new Date / 1000) / 60) / 60)} час(а/ов) 
                                ${Math.floor(((punish['time']['end'] - new Date / 1000) / 60) % 60)} минут(ы)`}
                    </td>
                    <td className={'banlist-tr-item flex center'}>
                        <img className={'player-head-right'}
                             src={punish['operator']['uuid'] !== '00000000-0000-0000-0000-000000000000' ?
                                 `https://visage.surgeplay.com/face/32/${punish['operator']['uuid']}`
                                 : `https://visage.surgeplay.com/face/32/console`
                        } />
                        {punish['operator']['name']}
                    </td>
                </tr>
            ))}
        </tbody>
    )
}