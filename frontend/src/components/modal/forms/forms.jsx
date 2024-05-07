import data from "../../../../data.json";
import React, {useState} from "react";
// import axios from "axios";

import Inputs from "./inputs.jsx";

import './forms.css';
import '../../../pages/pages.css'

export default function Form({item}) {
    const [activeIndex, setActiveIndex] = useState(null);

    const [formData, setFormData] = useState({
        nickname: '',
        mail: '',
        color: '',
        badge: ''
    });

    const handleChange = (event) => {
        const { name, value } = event.target;
        setFormData({ ...formData, [name]: value });
    }
    const handleDurationClick = (index) => {setActiveIndex(index === activeIndex ? null : index)}

    const handleGetForm = async () => {
        const { nickname, mail, color, badge } = formData;
        if (nickname === '') {
            return alert('Введи никнейм!!!');
        }
        if (mail === '') {
            return alert('Введи почту!!!');
        }
        if (color === '') {
            return alert('Введи цвет!!!');
        }
        if (badge === '') {
            return alert('Введи значок!!!');
        }

        // return alert(`${nicknameValue}, ${mailValue}, ${colorValue}, ${badgeValue}`)

        //
        // if (activeIndex !== null) {
        //     let amount = "6900"
        //     let donat = "hronon"
        //     let color = "#ffffff"
        //     let nickname = "yacheru"
        //     const response = await axios.post(`http://localhost:8000/v1/mc/payments?amount=${amount}&donat=${donat}&color=${color}&nickname=${nickname}`)
        //     return window.open(response.data['confirmation']['confirmation_url'])
        // } else {
        //     return console.log("Див не выбран");
        // }
    };

    let costsLen = Object.keys(data[item].costs).length

    return (
        <form className={'modal__form'} action="">
            <fieldset className={'modal__fieldset'}>
                <label>
                    <input
                        className={'modal__input b'}
                        name={'nickname'}
                        type="text"
                        onChange={handleChange}
                        value={formData.nickname}
                        placeholder={'Введите ваш никнейм'}
                        id={'nickname'}
                    />
                    <input
                        className={'modal__input b'}
                        name={'mail'}
                        type="text"
                        onChange={handleChange}
                        value={formData.mail}
                        placeholder={'Введите вашу почту'}
                        id={'mail'}
                    />
                    <Inputs item={item} handleChange={handleChange} colorValue={formData.color} badgeValue={formData.badge} />
                </label>
            </fieldset>
            <div className={'modal__durations flex'}>
                {Array.from({length: costsLen}, (_, i) => (
                    <div
                        className={`modal__duration flex bgc-1 b br10 ${activeIndex === i || costsLen === 1 ? 'duration-active' : ''}`}
                        key={i} onClick={() => handleDurationClick(i)}>
                        <div className={`modal__duration-checkbox`}></div>
                        <div className={'modal__duration-text flex'}>
                            <p>{ data[item].costs[`${i + 1}`][1] }</p>
                            <span>/</span>
                            <p>{ data[item].costs[`${i + 1}`][0] }</p>
                        </div>
                    </div>
                ))}
            </div>
            <div className={'modal__navbuy flex'}>
                <button className={'modal__navbuy-button bgc-1 b br10'} onClick={handleGetForm} type={'button'}>Продолжить</button>
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
        </form>
    )
}