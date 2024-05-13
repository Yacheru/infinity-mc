import { useLocation } from "react-router-dom";

export default function NoBans() {
    const location = useLocation()
    const values = {
        'bans': 'банов',
        'mutes': 'мутов',
        'warns': 'предупреждений'
    }
    const handleError = () => {
        let value = location.search.split('=')[1]

        if (values[value] === undefined) return `Неверно указанная категория😕`

        return `не удалось получить список ${ values[value] }😥`
    }

    return (
        <tbody className={'banlist-table__body no-push flex'}>
        <tr className={'banlist-body__tr no-push flex'}>
            <td className={'banlist-tr-item no-push flex center'}>{handleError()}</td>
        </tr>
        </tbody>
    )
}
