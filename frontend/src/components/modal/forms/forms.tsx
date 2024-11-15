import { IData, Durations } from "$types/data";
import { IPaymentsService } from "$types/api";
import { IForm } from "$types/modal";

import { useEffect, useRef, useState, ChangeEvent, MouseEvent, FocusEvent } from "react";
import toast from "react-hot-toast";
import { useTranslation } from "react-i18next";

import PaymentsService from "@api/axios/requests/payments.ts";
import data from "@config/data.json";
import { AxiosResponse } from "axios"

import '@styles/components/modal/forms.css';
import '@styles/pages/pages.css';

const typedData: IData = data as IData

const payService: IPaymentsService = new PaymentsService()

export default function Form({ item }: IForm) {
    const { t } = useTranslation()

    const checkboxRef = useRef(null)

    const [placeholder, setPlaceholder] = useState<string>('Продолжить')
    const [checkbox, setCheckbox] = useState<boolean>(true)
    const [emailDirty, setEmailDirty] = useState<boolean>(false)
    const [nicknameDirty, setNicknameDirty] = useState<boolean>(false)
    const [formValid, setFormValid] = useState<boolean>(false)
    const [duration, setDuration] = useState<Durations>("15552000")


    useEffect(() => {
        setFormValid(checkbox)
    }, [checkbox])

    const checkboxHandler = (e: MouseEvent<HTMLInputElement>) => {
        setCheckbox(!e.currentTarget.checked)
    }

    const submitForm = async () => {
        setFormValid(false)
        setPlaceholder("Ожидайте...")

        try {
            const costs = typedData[item].costs;
            const price: string = typedData[item].costs[duration as keyof typeof costs][1]
            const payResponse: AxiosResponse = await payService.createPayment(price, email, item, nickname, duration)

            return window.open(payResponse.data['confirmation']['confirmation_url'])
        } catch (e) {
            setPlaceholder("Ошибка запроса.")

            toast.error("Ошибка выполнения запроса...", {
                icon: "❌",
                style: {
                    color: "whitesmoke",
                    background: '#303030',
                    borderRadius: '10px',
                },
                duration: 1500,
            })
        } finally {
            setTimeout(() => window.location.reload(), 1000)
        }
    }

    return (
        <form className={'modal__form'}>
            <div className={'modal__durations flex'}>
            {Array.from({length: Object.keys(typedData[item].costs).length}, (_, i) => (
                <div
                    className={`modal__duration flex bgc-1 b br10 ${duration === typedData[item].durations[i] ? 'duration-active' : ''}`}
                    key={i}
                    onClick={() => setDuration(typedData[item].durations[i])}>
                    <div className={`modal__duration-checkbox`}></div>
                    <div className={'modal__duration-text flex'}>
                        <p>{ duration }₽</p>
                        <span>/</span>
                        <p>{ duration }</p>
                    </div>
                </div>
            ))}
            </div>
            <div className={'modal__navbuy flex'}>
                <button disabled={!formValid} className={'modal__navbuy-button bgc-1 b br10'} type={'button'} onClick={() => submitForm()}>{placeholder}</button>
                <div className={'modal__checkbox-box flex'}>
                    <div className={'modal__checkbox-item'}>
                        <label className={'modal__checkbox-label flex h100'} htmlFor={'checkbox'}>
                            <input onClick={e => checkboxHandler(e)} ref={checkboxRef} name={'checkbox'} type="checkbox" id={'checkbox'}/>
                        </label>
                    </div>
                    <p className={'modal__checkbox-text'}>
                        Я принимаю условия <a href="/terms" target={'_blank'}> пользовательского соглашения</a>
                    </p>
                </div>
            </div>
        </form>
    )
}