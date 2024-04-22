import data from "../../../../../data.json";
import React, {useState} from "react";

import './forms.css'

export default function Form({item}) {
    const [activeIndex, setActiveIndex] = useState(null);
    const [inputValue, setInputValue] = useState('');

    const handleChange = (event) => {
        setInputValue(event.target.value);
    };

    const handleDurationClick = (index) => {
        setActiveIndex(index === activeIndex ? null : index);
    };

    const handleGetForm = () => {
        if (inputValue === '') {
            alert('Введите свой никнейм!!!')
        }

        if (activeIndex !== null) {
            console.log("Выбранный див:", activeIndex, inputValue);
        } else {
            console.log("Див не выбран");
        }
    };

    let costsLen = Object.keys(data[item].costs).length

    return (
        <form className={'modal__form'} action="">
            <fieldset className={'modal__fieldset'}>
                <label htmlFor="nickname">
                    <input className={'modal__input'} name={'nickname'} type="text" onChange={handleChange} value={inputValue} placeholder={'Введите ваш никнейм'} id={'nickname'}/>
                </label>
            </fieldset>
            <div className={'modal__durations'}>
                {Array.from({length: costsLen}, (_, i) => (
                    <div className={`modal__duration${activeIndex === i || costsLen === 1 ? ' duration-active' : ''}`} key={i} onClick={() => handleDurationClick(i)}>
                        <div className={`modal__duration-checkbox${activeIndex === i || costsLen === 1  ? ' checkbox-active' : ''}`}></div>
                        <div className={'modal__duration-text'}>
                            <p>{ data[item].costs[`${i + 1}`][0] }</p>
                            <p>{data[item].costs[`${i + 1}`][1]}</p>
                        </div>
                    </div>
                ))}
            </div>
            <div className={'modal__navbuy'}>
                <button className={'modal__navbuy-button'} onClick={handleGetForm} type={'button'}>Продолжить</button>
                <div className={'modal__checkbox-box'}>
                    <div className={'modal__checkbox-item'}>
                        <label className={'modal__checkbox-label'} htmlFor="">
                            <input type="checkbox"/>
                        </label>
                    </div>
                    <div className={'modal__checkbox-text'}>
                        <span> Я принимаю условия </span>
                        <a href="/terms" target={'_blank'}>пользовательского соглашения</a>
                    </div>
                </div>
            </div>
        </form>
    )
}