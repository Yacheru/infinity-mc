import '../../styles/pages/stats/stats.css'
import Header from "../../components/header/header.js";
import React from 'react';

export default function Stats() {
    return (
        <>
            <div className={'container'}>
                <Header />
                <div className={'flex center in-dev'}>
                    IN DEV
                </div>
            </div>
        </>
    )
}