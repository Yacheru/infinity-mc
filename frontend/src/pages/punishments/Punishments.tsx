import React from "react";
import Header from "../../components/header/header.js";
import Category from "./category/Category.js";
import PunishmentsList from "./punishments list/punishmentsList.js";

export default function Punishments() {
    return (
        <div className='container'>
            <Header />
            <Category />
            <PunishmentsList />
        </div>
    )
}