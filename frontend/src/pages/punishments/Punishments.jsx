import Header from "../../components/header/header.jsx";
import Category from "./category/Category.jsx";
import PunishmentsList from "./punishments list/punishmentsList.jsx";

export default function Punishments() {
    return (
        <div className='container'>
            <Header />
            <Category />
            <PunishmentsList />
        </div>
    )
}