export default function BansBody({ punishments }) {
    return (
        <tbody className={'banlist-table__body flex'}>
            {punishments && punishments.map((ban, index) => (
                <tr className={'banlist-body__tr flex'} key={index}>
                    <td className={'banlist-tr-item flex center'}>{ban.victim_name}</td>
                    <td className={'banlist-tr-item flex center'}>{ban.reason}</td>
                    <td className={'banlist-tr-item flex center'}>{ban.end === 0 ?
                        <p className={'br10 forever'}>Навсегда</p> : `${new Date(ban.end).getDate()}д. ${new Date(ban.end).getHours()}ч. ${new Date(ban.end).getMinutes()}мин.`}</td>
                    <td className={'banlist-tr-item flex center'}>{ban.operator}</td>
                </tr>
            ))}
        </tbody>
    )
}