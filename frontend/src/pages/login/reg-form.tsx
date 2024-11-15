import {ChangeEvent, useContext, useState} from "react";

import { Context } from "../../context.tsx";
import { observer } from "mobx-react-lite";

import useEmailValidator from "@hooks/useEmailValidation.ts";
import usePasswordValidator from "@hooks/usePasswordValidator.ts";
import useNicknameValidator from "@hooks/useNicknameValidator.ts";
import useDoublePasswordValidator from "@hooks/useDoublePasswordValidator.ts";

import Modal from "@components/modal/modal.tsx";

export default observer(function RegForm() {
    const { email, emailValid, emailError, emailHandler } = useEmailValidator();
    const { password, passwordValid, passwordError, passwordHandler } = usePasswordValidator();
    const { nickname, nicknameValid, nicknameError, nicknameHandler } = useNicknameValidator();
    const { confirmPassword, passwordsMatch, passwordsError, passwordsHandler } = useDoublePasswordValidator();

    const [ modalActive, setModalActive ] = useState<boolean>(false)

    const [ registerError, setRegisterError ] = useState<string>('')
    const [ codeError, setCodeError ] = useState<string>('')

    const [code, setCode] = useState('');
    const [isValid, setIsValid] = useState(false);

    const { auth } = useContext(Context);

    const handleInputChange = (e: ChangeEvent<HTMLInputElement>) => {
        const input = e.target.value.replace(/\D/g, '').slice(0, 4);
        setCode(input);
        setIsValid(input.length === 4);
    };

    return (
        <form className="reg-form">
            <fieldset className="reg-fields">
                {!nicknameValid && <p className='input-error'>{nicknameError}</p>}
                <div className='input-wrapper input-user'>
                    <input
                        className='input b br10'
                        type="text"
                        placeholder='Введите свой никнейм'
                        value={nickname}
                        onChange={nicknameHandler}
                    />
                </div>
                {!emailValid && <p className='input-error'>{emailError}</p>}
                <div className='input-wrapper input-email'>
                    <input
                        value={email}
                        onChange={emailHandler}
                        className='input b br10'
                        type="email"
                        placeholder='Введите свою почту'
                    />
                </div>
                {!passwordValid && <p className='input-error'>{passwordError}</p>}
                <div className='input-wrapper input-password'>
                    <input
                        value={password}
                        onChange={passwordHandler}
                        className='input b br10'
                        type="password"
                        placeholder='Введите пароль'
                    />
                </div>
                {!passwordsMatch && <p className='input-error'>{passwordsError}</p>}
                <div className='input-wrapper input-password'>
                    <input
                        value={confirmPassword}
                        onChange={(e: ChangeEvent<HTMLInputElement>) => passwordsHandler(e, password)}
                        className='input b br10'
                        type="password"
                        placeholder='Введите пароль ещё раз'
                    />
                </div>
            </fieldset>
            <button
                disabled={!nicknameValid || !emailValid || !passwordValid || !passwordsMatch}
                type='button'
                className='submit-button b br10 bgc-1 w100'
                onClick={() => auth.codeSend(email, setModalActive)}
            >
                Зарегистрироваться
            </button>
            {codeError !== '' && <p className='submit-error'>{codeError}</p>}
            <Modal width={500} active={modalActive} setActive={setModalActive}>
                <div className={'modal-code'}>
                    <p className={'modal-code-description'}>
                        На почту <b>{email}</b> отправлен 4-значный код.<br/>
                        Для завершения регистрации введите его ниже.
                    </p>
                    <div className={'code-box flex'}>
                        <input
                            className='input input-code b br10'
                            type="text"
                            placeholder='Введите код подтверждения сюда'
                            maxLength={4}
                            value={code}
                            onChange={handleInputChange}
                        />
                        <button type='button' onClick={() => auth.register(nickname, email, password, code)} disabled={!isValid} className={'button-code b bgc-1 br10'}>Отправить</button>
                    </div>
                    {registerError !== '' && <p className='submit-error'>{registerError}</p>}
                </div>
            </Modal>
        </form>
    );
});