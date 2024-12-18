import React from 'react';
import '../../../../styles/pages/terms/terms.css'
import {useTranslation} from "react-i18next";

export default function Main() {
    const t = useTranslation()

    return (
        <div className={'term__container'}>
            <main className={'term__main bgc-1 b br20 h100'}>
                <div className={"term__header"}>
                        <div className={'term__title'}>Пользовательское соглашение</div>
                        <div className={'term__ordered-preface'}>
                            Официальное пользовательское соглашение проекта INFINITY-MC, где мы описываем
                            правила использования наших сервисов и серверов. Прошу, обрати внимание, что все остальные
                            услуги не относятся к этому соглашению.
                        </div>
                    </div>
                <hr className='modal__hr' size='3' color='whitesmoke'/>
                <ol className={'term__ordered-list'}>
                        <hr/>
                        <li className={'term__ordered-item'}>
                            Пользуясь нашими сервисами, ты автоматически соглашаешься с этим пользовательским
                            соглашением и правилами сервера. Если что-то не нравится, можешь прекратить пользоваться
                            нашими сервисами.
                        </li>
                        <hr/>
                        <li className={'term__ordered-item'}>
                            Оплачивать услуги следует только на официальном сайте проекта, который находится по адресу
                            <a className={'important'} href="https://infinity-mc.ru/"> https://infinity-mc.ru/</a>. Мы принимаем оплату только в
                            российских рублях через агрегатор платежей, который может конвертировать другие валюты.
                        </li>
                        <hr/>
                        <li className={'term__ordered-item'}>
                            Возврат денег невозможен ни при каких обстоятельствах, так что внимательно изучи все условия
                            перед оплатой.
                        </li>
                        <hr/>
                        <li className={'term__ordered-item'}>
                            Решение о наложении ограничений на игрока за нарушение правил или пользовательского
                            соглашения принимается администраторами сервера в каждом случае индивидуально.
                        </li>
                        <hr/>
                        <li className={'term__ordered-item'}>
                            Попытки влиять на администрацию сервера для смягчения ограничений могут привести к
                            ужесточению мер и уменьшению шансов на снятие ограничений.
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
                            Помни, что запрещено распространение контента <a className={'important'} href="https://ru.wikipedia.org/wiki/NSFW">NSFW</a> (Not Safe for Work), шок-контента, а
                            также контента, запрещённого на территории Российской Федерации, в любом формате.
                        </li>
                        <li className={'term__ordered-item'}>
                            Услуга на сервере появится сразу же после оплаты. Если вы столкнулись с трудностями,
                            сообщите нам в
                            <a
                                className={'important'}
                                href="https://discord.gg/infinity-tm-494212272353181726"
                                target={'_blank'}> дискорд
                            </a> через канал
                            <a
                                className={'important'}
                                href="https://discord.com/channels/494212272353181726/1148376245428559913"
                                target={'_blank'}>
                                💌〢поддержка
                            </a>,
                            мы обязательно Вам поможем.
                        </li>
                    </ol>
            </main>
        </div>
    )
}