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
        if (status === 429) return `Не так быстро!🤬`

        return `не удалось получить список ${ failed[category][0] }😥`
    }

    return (
        <div className={'punishment-table__body flex center h100'}>
            <div className={'punishment-body__tr-no-push flex'}>
                <div className={'punishment-tr-item-no-push flex center w100'}>{handleError()}</div>
            </div>
        </div>
    )
}
