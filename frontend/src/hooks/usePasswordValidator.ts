import {ChangeEvent, useState} from "react";

export default function usePasswordValidator() {
    const [password, setPassword] = useState<string>('');
    const [passwordValid, setPasswordValid] = useState<boolean>(false);
    const [passwordError, setPasswordError] = useState<string>('');

    const validatePassword = (password: string) => {
        if (!password) {
            setPasswordValid(false);
            setPasswordError('Укажите пароль!')
        }

        if (password.length < 8 || password.length > 24) {
            setPasswordValid(false);
            setPasswordError('Пароль должен быть в пределах 8-24 символов')
        } else {
            setPasswordValid(true)
            setPasswordError('')
        }
    }

    const passwordHandler = (e: ChangeEvent<HTMLInputElement>) => {
        setPassword(e.target.value)
        validatePassword(e.target.value)
    }

    return {
        password,
        passwordValid,
        passwordError,
        passwordHandler
    }
}