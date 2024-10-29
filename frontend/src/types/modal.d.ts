export interface IModal {
    active: boolean;
    setActive: (active: boolean) => void;
    modalType: string
    item: Items
}

export interface IDescription {
    item: Items
}

export interface IForm {
    item: Items
}

export interface INavBuy {
    valid: boolean
}

export type Items = "hronon" | "nickname" | "badge"
