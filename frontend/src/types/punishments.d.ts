export interface INoBans {
    location: Location,
    status: number | null,
}

export interface IBansBody {
    punishments: IPunishment[]
}

export interface IPunishment {
    id: number,
    victim: {
        uuid: string,
        name: string
    },
    reason: string,
    time: {
        start: number,
        end: number,
    },
    operator: {
        uuid: string,
        name: string
    }
}

export type Categories = "bans" | "mutes" | "warns"