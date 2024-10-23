const failed = {
    'bans': ['банов', 'забанен'],
    'mutes': ['мутов', 'замьючен'],
    'warns': ['предупреждений', 'предупреждён']
}

export default function NoBans({ location, status }) {
    const handleError = () => {
        let value = location.search.split('=')[1]

        if (failed[value] === undefined) return `Неверно указанная категория😕`
        if (status === 204) return `Ещё никто не ${ failed[value][1] }😏`

        return `не удалось получить список ${ failed[value][0] }😥`
    }

    return (
        <tbody className={'punishment-table__body flex center h100'}>
            <tr className={'punishment-body__tr-no-push flex'}>
                <td className={'punishment-tr-item-no-push flex center w100'}>{handleError()}</td>
            </tr>
        </tbody>
    )
}
