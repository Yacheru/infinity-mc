import './Category.css'

export default function Category() {
    return (
        <div className={'category-box flex'}>
            <div className={'category-box-item category-bans bgc-1 b br20 flex center'}>Баны</div>
            <div className={'category-box-item category-mutes bgc-1 b br20 flex center'}>Муты</div>
            <div className={'category-box-item category-warns bgc-1 b br20 flex center'}>Предупреждения</div>
        </div>

    )
}