import { Link, useLocation } from "react-router-dom";

import '../../../styles/pages/punishments/category.css'
import React from "react";

export default function Category() {
    const location = useLocation()
    let locate = location.search.split('=')[1]

    return (
        <div className={'category-box flex w100'}>
            <Link to={'/p?category=bans'} className={`category-box-item category-bans ${locate === 'bans' ? 'selected ' : ''}bgc-1 b br20 flex center`}>
                <svg fill="#fff" width="24px" height="24px" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                    <path d="M17,9V7A5,5,0,0,0,7,7V9a3,3,0,0,0-3,3v7a3,3,0,0,0,3,3H17a3,3,0,0,0,3-3V12A3,3,0,0,0,17,9ZM9,7a3,3,0,0,1,6,0V9H9Zm9,12a1,1,0,0,1-1,1H7a1,1,0,0,1-1-1V12a1,1,0,0,1,1-1H17a1,1,0,0,1,1,1Z"/>
                </svg>
                <p>Баны</p>
            </Link>
            <Link to={'/p?category=mutes'} className={`category-box-item category-mutes ${locate === 'mutes' ? 'selected ' : ''}bgc-1 b br20 flex center`}>
                <svg width="24px" height="24px" viewBox="0 0 1024 1024" className="icon" xmlns="http://www.w3.org/2000/svg">
                    <path fill="#fff"
                          d="M412.16 592.128l-45.44 45.44A191.232 191.232 0 01320 512V256a192 192 0 11384 0v44.352l-64 64V256a128 128 0 10-256 0v256c0 30.336 10.56 58.24 28.16 80.128zm51.968 38.592A128 128 0 00640 512v-57.152l64-64V512a192 192 0 01-287.68 166.528l47.808-47.808zM314.88 779.968l46.144-46.08A222.976 222.976 0 00480 768h64a224 224 0 00224-224v-32a32 32 0 1164 0v32a288 288 0 01-288 288v64h64a32 32 0 110 64H416a32 32 0 110-64h64v-64c-61.44 0-118.4-19.2-165.12-52.032zM266.752 737.6A286.976 286.976 0 01192 544v-32a32 32 0 0164 0v32c0 56.832 21.184 108.8 56.064 148.288L266.752 737.6z"/>
                    <path fill="#fff" d="M150.72 859.072a32 32 0 01-45.44-45.056l704-708.544a32 32 0 0145.44 45.056l-704 708.544z"/>
                </svg>
                <p>Муты</p>
            </Link>
            <Link to={'/p?category=warns'} className={`category-box-item category-warns ${locate === 'warns' ? 'selected ' : ''}bgc-1 b br20 flex center`}>
                <svg width="24px" height="24px" viewBox="0 0 24 24" fill="#fff" xmlns="http://www.w3.org/2000/svg">
                    <path d="M12 14a1 1 0 0 1-1-1v-3a1 1 0 1 1 2 0v3a1 1 0 0 1-1 1zm-1.5 2.5a1.5 1.5 0 1 1 3 0 1.5 1.5 0 0 1-3 0z" fill="#fff"/>
                    <path d="M10.23 3.216c.75-1.425 2.79-1.425 3.54 0l8.343 15.852C22.814 20.4 21.85 22 20.343 22H3.657c-1.505 0-2.47-1.6-1.77-2.931L10.23 3.216zM20.344 20L12 4.147 3.656 20h16.688z" fill="#fff"/>
                </svg>
                <p>Предупреждения</p>
            </Link>
        </div>
    )
}