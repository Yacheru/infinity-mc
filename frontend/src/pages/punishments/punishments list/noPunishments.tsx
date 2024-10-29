import React from "react";

interface INoBans {
    location: Location,
    status: number | null,
}

type Categories = "bans" | "mutes" | "warns"

export default function NoBans({ location, status }: INoBans) {
    const failed = {
        "bans": ['банов', 'забанен'],
        "mutes": ['мутов', 'замьючен'],
        "warns": ['предупреждений', 'предупреждён']
    }

    const handleError = () => {
        const value: string = location.search.split('=')[1]
        const category = value ? (value as Categories) : null;

        if (!category) return 'Категория не указана😐'
        if (!failed[category]) return `Неверно указанная категория😕`
        if (status === 204) return `Ещё никто не ${ failed[category][1] }😏`

        return `не удалось получить список ${ failed[category][0] }😥`
    }

    return (
        <tbody className={'punishment-table__body flex center h100'}>
            <tr className={'punishment-body__tr-no-push flex'}>
                <td className={'punishment-tr-item-no-push flex center w100'}>{handleError()}</td>
            </tr>
        </tbody>
    )
}
