import React from "react";

export default function NavBuy({ valid }) {
    return (
        <div className={'modal__navbuy flex'}>
            <button disabled={!valid} className={'modal__navbuy-button bgc-1 b br10'} type={'button'}>Продолжить</button>
            <div className={'modal__checkbox-box flex'}>
                <div className={'modal__checkbox-item'}>
                    <label className={'modal__checkbox-label flex'} htmlFor="checkbox">
                        <input type="checkbox" id={'checkbox'}/>
                    </label>
                </div>
                <p className={'modal__checkbox-text'}>
                    Я принимаю условия
                    <a href="/terms" target={'_blank'}> пользовательского соглашения</a>
                </p>
            </div>
        </div>
    )
}