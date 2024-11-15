import useEmailValidator from "@hooks/useEmailValidation.ts";
import usePasswordValidator from "@hooks/usePasswordValidator.ts";

import { useContext, useState } from "react";
import { Context } from "../../context.tsx";
import { observer } from "mobx-react-lite";

export default observer(function LoginForm() {
    const { email, emailValid, emailError, emailHandler } = useEmailValidator();
    const { password, passwordValid, passwordError, passwordHandler } = usePasswordValidator()
    const [ loginError, setLoginError ] = useState('')

    const { auth } = useContext(Context);

    return (
        <form className="login-form">
            <fieldset className="login-fields">
                {!emailValid && <p className='input-error'>{emailError}</p>}
                <div className="input-wrapper input-email">
                    <input
                        value={email}
                        onChange={emailHandler}
                        className='input b br10'
                        type="text"
                        name='email'
                        placeholder='Введите свою почту...'
                    />
                </div>
                {!passwordValid && <p className='input-error'>{passwordError}</p>}
                <div className="input-wrapper input-password">
                    <input
                        value={password}
                        onChange={passwordHandler}
                        className='input input-password b br10'
                        type="password"
                        name='password'
                        placeholder='Введите пароль...'
                    />
                </div>
            </fieldset>
            <button
                disabled={!emailValid || !passwordValid}
                onClick={() => auth.login(email, password)}
                type='button'
                className='submit-button b br10 bgc-1 w100'
            >
                Войти
            </button>
            {loginError && <p className='submit-error'>{loginError}</p>}
        </form>
    )
});