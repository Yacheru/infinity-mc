import { ChangeEvent, useState } from "react";

export default function useDoublePasswordValidator() {
    const [confirmPassword, setConfirmPassword] = useState<string>('');
    const [passwordsMatch, setPasswordsMatch] = useState<boolean>(false);
    const [passwordsError, setPasswordsError] = useState<string>('');

    const validatePasswords = (password: string, firstPassword: string) => {
        if (password === firstPassword) {
            setPasswordsMatch(true);
            setPasswordsError('')
        } else {
            setPasswordsMatch(false);
            setPasswordsError('Пароли не совпадают')
        }
    }

    const passwordsHandler = (e: ChangeEvent<HTMLInputElement>, firstPassword: string) => {
        setConfirmPassword(e.target.value)
        validatePasswords(e.target.value, firstPassword);
    }

    return {
        confirmPassword,
        passwordsMatch,
        passwordsError,
        passwordsHandler
    }
}