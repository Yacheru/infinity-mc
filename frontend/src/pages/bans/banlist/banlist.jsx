import './banlist.css'

export default function Banlist() {
    return (
        <main className={'banlist-box b bgc-1 br20'}>
            <div className={'banlist-box__title flex'}>
                <div className={'banlist__title-item b bgc-2 br10 flex center'}>Нарушитель</div>
                <div className={'banlist__title-item b bgc-2 br10 flex center'}>Причина</div>
                <div className={'banlist__title-item b bgc-2 br10 flex center'}>Срок</div>
            </div>
            <hr/>
        </main>
    )
}