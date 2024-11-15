import {ChangeEvent, useState} from "react";

export default function useNicknameValidator() {
    const [nickname, setNickname] = useState<string>('');
    const [nicknameValid, setNicknameValid] = useState<boolean>(false);
    const [nicknameError, setNicknameError] = useState<string>('');

    const validatePassword = (password: string) => {
        if (!password) {
            setNicknameValid(false);
            setNicknameError('Укажите свой никнейм!')
        }

        if (password.length < 3 || password.length > 16) {
            setNicknameValid(false);
            setNicknameError('Никнейм должен быть от 3 до 16 символов')
        } else {
            setNicknameValid(true)
            setNicknameError('')
        }
    }

    const nicknameHandler = (e: ChangeEvent<HTMLInputElement>) => {
        setNickname(e.target.value)
        validatePassword(e.target.value)
    }

    return {
        nickname,
        nicknameValid,
        nicknameError,
        nicknameHandler
    }
}