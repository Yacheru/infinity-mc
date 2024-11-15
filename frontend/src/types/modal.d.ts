import { CSSProperties, ReactNode } from 'react'

export interface IModal {
    active: boolean;
    setActive: (active: boolean) => void;
    children?: ReactNode;
    width?: number
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
