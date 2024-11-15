import { useState, ChangeEvent } from "react";

const re = /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|.(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/

export default function useEmailValidator() {
    const [email, setEmail] = useState<string>('');
    const [emailValid, setEmailValid] = useState<boolean>(false);
    const [emailError, setEmailError] = useState<string>('')

    const validateEmail = (email: string) => {
        if (!email) {
            setEmailValid(false)
            setEmailError('Укажите почту')
        }

        if (!re.test(String(email).toLowerCase())) {
            setEmailValid(false)
            setEmailError('Почта введена некорректно')
        } else {
            setEmailValid(true)
            setEmailError('')
        }
    }

    const emailHandler = (e: ChangeEvent<HTMLInputElement>) => {
        setEmail(e.target.value)
        validateEmail(e.target.value)
    }

    return {
        email,
        emailValid,
        emailError,
        emailHandler
    }
}