import React, {useEffect, useRef, useState} from "react";

import './forms.css';
import '../../../pages/pages.css'
import data from "../../../../data.json";
import cfg from "../../../../config.json";
import axios from "axios";

export default function Form({ item }) {
    const checkboxRef = useRef(null);

    const [nickname, setNickname] = useState('');
    const [email, setEmail] = useState('');
    const [checkbox, setCheckbox] = useState(true)
    const [emailDirty, setEmailDirty] = useState(false);
    const [nicknameDirty, setNicknameDirty] = useState(false);
    const [nicknameError, setNicknameError] = useState('Поле никнейм не может быть пустым');
    const [emailError, setEmailError] = useState('Поле почты не может быть пустым');
    const [formValid, setFormValid] = useState(false)
    const [duration, setDuration] = useState(0);

    useEffect(() => {
        if (emailError || nicknameError || checkbox) {
            setFormValid(false)
        } else {
            setFormValid(true)
        }
    }, [emailError, nicknameError, checkbox])

    const indexHandler = (i) => {
        setDuration(i)
    }

    const nicknameHandler = (e) => {
        setNickname(e.target.value)
        if (e.target.value.length < 3) {
            setNicknameError('Никнейм не может быть короче 3 символов')
            if (!e.target.value) {
                setNicknameError('Поле никнейм не может быть пустым')
            }
            e.target.classList.add('required')
        } else {
            e.target.classList.remove('required')
            setNicknameError('')
        }
    }

    const emailHandler = (e) => {
        setEmail(e.target.value)
        const re = /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|.(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/
        if (!re.test(String(e.target.value).toLowerCase())) {
            setEmailError('Почта введена некорректно')
            e.target.classList.add('required')
        } else {
            e.target.classList.remove('required')
            setEmailError('')
        }
    }

    const checkboxHandler = (e) => {
        if (e.target.checked) {
            setCheckbox(false)
        } else {
            setCheckbox(true)
        }
    }

    const blurHandler = (e) => {
        switch (e.target.name) {
            case 'email':
                return setEmailDirty(true)
            case 'nickname':
                return setNicknameDirty(true)
        }
    }

    const submitForm = (e) => {
        const amount = {
            0: '169',
            1: '369',
            2: '690',
        }
        let url

        switch (cfg['status']) {
            case 'local':
                url = `http://localhost/v1/payment/?nickname=${nickname}&email=${email}&amount=${amount[duration]}&donat=${item}`
                break
            case 'prod':
                url = `https://api.infinity-mc.ru/v1/payment/?nickname=${nickname}&email=${email}&amount=${amount[duration]}&donat=${item}`
                break
        }

        axios.get(url, {
            auth: {
                username: cfg['user'],
                password: cfg['pass']
            }
        }).then((res) => {
            return window.open(res.data['confirmation']['confirmation_url'])
        });
    }

    return (
        <form className={'modal__form'}>
            <fieldset className={'modal__fieldset'}>
                <label>
                    {(nicknameDirty && nicknameError) && <div style={{color: "red"}}>{nicknameError}</div>}
                    <div className={'input-container nickname'}>
                        <input onChange={e => nicknameHandler(e)} value={nickname} onBlur={e => blurHandler(e)} className={'modal__input b'} name={'nickname'} type="text" placeholder={'Введите ваш никнейм'} id={'nickname'}/>
                    </div>
                    {(emailDirty && emailError) && <div style={{color: "red"}}>{emailError}</div>}
                    <div className={'input-container email'}>
                        <input onChange={e => emailHandler(e)} value={email} onBlur={e => blurHandler(e)} className={'modal__input b'} name={'email'} type="text" placeholder={'Введите вашу почту'} id={'email'}/>
                    </div>
                </label>
            </fieldset>
            <div className={'modal__durations flex'}>
            {Array.from({length: Object.keys(data[item].costs).length}, (_, i) => (
                    <div className={`modal__duration flex bgc-1 b br10 ${duration === i ? 'duration-active' : ''}`}
                         key={i} onClick={() => indexHandler(i)}>
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
                <button disabled={!formValid} className={'modal__navbuy-button bgc-1 b br10'} type={'button'} onClick={e => submitForm(e)}>Продолжить</button>
                <div className={'modal__checkbox-box flex'}>
                    <div className={'modal__checkbox-item'}>
                        <label className={'modal__checkbox-label flex'} htmlFor={'checkbox'}>
                            <input onClick={e => checkboxHandler(e)} ref={checkboxRef} name={'checkbox'} type="checkbox" id={'checkbox'}/>
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