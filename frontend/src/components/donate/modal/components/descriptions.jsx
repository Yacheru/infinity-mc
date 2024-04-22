import data from "../../../../../data.json";
import React from "react";

import './description.css'

export default function Description({ item }) {
    return (
        <p className='modal__description'>
            {data[item].description.split('\n').map((line, index) => (
                <React.Fragment key={index}>
                    {line}
                    <br/>
                </React.Fragment>
            ))}
        </p>
    )
}