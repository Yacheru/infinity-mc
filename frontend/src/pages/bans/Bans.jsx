import Header from "../root/components/header/header.jsx";
import Category from "./category/Category.jsx";
import Banlist from "./banlist/banlist.jsx";

export default function Bans() {
    return (
        <div className='container'>
            <Header />
            <Category />
            <Banlist />
        </div>
    )
}