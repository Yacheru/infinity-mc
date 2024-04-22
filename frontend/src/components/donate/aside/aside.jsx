import './aside.css'
import '../container.css'

export default function Aside() {
    return (
        <aside className="aside">
            <div className="aside__text">
                <div className="aside__title">
                    <h1>Покупка услуг</h1>
                </div>
                <div className="aside__items">
                    <p>Покупая донат, вы не только наслаждаетесь игрой, но и активно способствуете нашему развитию!</p>
                    <p>Благодарим каждого игрока, который поддерживает нас финансово и помогает сделать наш сервер еще лучше!</p>
                </div>
            </div>
        </aside>
    )
}