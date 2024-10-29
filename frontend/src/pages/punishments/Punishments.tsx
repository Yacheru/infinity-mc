import React from "react";
import Header from "../../components/header/header.tsx";
import Category from "./category/Category.tsx";
import PunishmentsList from "./punishments list/punishmentsList.tsx";

export default function Punishments() {
    return (
        <div className='container'>
            <Header />
            <Category />
            <PunishmentsList />
        </div>
    )
}