import './main.css'

export default function Main() {
    return (
        <div>
            <header className={'term__header'}>Пользовательское соглашение</header>
            <hr/>
            <main className={'term__main'}>
                <div className={'term__container'}>
                    <div className={'term__ordered-preface'}>
                        Официальное пользовательское соглашение проекта INFINITY-MC, где мы описываем
                        правила использования наших сервисов и серверов. Прошу, обрати внимание, что все остальные
                        услуги не относятся к этому соглашению.
                        <br/><br/>
                        Что важно знать:
                    </div>
                    <ol className={'term__ordered-list'}>
                        <hr/>
                        <li className={'term__ordered-item'}>
                            Пользуясь нашими сервисами, ты автоматически соглашаешься с этим пользовательским соглашением и правилами сервера. Если что-то не нравится, можешь прекратить пользоваться нашими сервисами.
                        </li>
                        <hr/>
                        <li className={'term__ordered-item'}>
                            Оплачивать услуги следует только на официальном сайте проекта, который находится по адресу
                            <a href="https://infinity-mc.ru/"> https://infinity-mc.ru/</a>. Мы принимаем оплату только в российских рублях через агрегатор платежей, который может конвертировать другие валюты.
                        </li>
                        <hr/>
                        <li className={'term__ordered-item'}>
                            Возврат денег невозможен ни при каких обстоятельствах, так что внимательно изучи все условия перед оплатой.
                        </li>
                        <hr/>
                        <li className={'term__ordered-item'}>
                            Решение о наложении ограничений на игрока за нарушение правил или пользовательского соглашения принимается администраторами сервера в каждом случае индивидуально.
                        </li>
                        <hr/>
                        <li className={'term__ordered-item'}>
                            Попытки влиять на администрацию сервера для смягчения ограничений могут привести к ужесточению мер и уменьшению шансов на снятие ограничений.
                        </li>
                        <hr/>
                        <li className={'term__ordered-item'}>
                            Использование дополнительных аккаунтов для обхода ограничений запрещено.
                        </li>
                        <hr/>
                        <li className={'term__ordered-item'}>
                            Если получишь блокировку доступа, она распространяется на все сервера сети INFINITY-MC.
                        </li>
                        <hr/>
                        <li className={'term__ordered-item'}>
                            Помни, что запрещено распространение контента NSFW (Not Safe for Work), шок-контента, а также контента, запрещённого на территории Российской Федерации, в любом формате.
                        </li>
                    </ol>
                </div>
            </main>
        </div>
    )
}