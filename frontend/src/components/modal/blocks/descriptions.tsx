import { IData } from "$types/data";
import { IDescription } from "$types/modal";

import data from "../../../config/data.json";
import React from "react";

import '../../../styles/components/modal/description.css'

const typeData: IData = data

export default function Description({ item }: IDescription) {
    return (
        <p className='modal__description'>
            {typeData[item].description.split('\n').map((line, index) => (
                <React.Fragment key={index}>
                    {line}
                    <br/>
                </React.Fragment>
            ))}
        </p>
    )
}