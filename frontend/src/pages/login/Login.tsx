import {useContext, useState} from "react";

import Header from "@components/header/header.tsx";

import LoginForm from "./login-form.tsx";
import RegForm from "./reg-form.tsx";

import '@styles/pages/login/login.css'
import {Toaster} from "react-hot-toast";
import {Context} from "../../context.tsx";


export default function Login() {
    const [isReg, setIsReg] = useState<boolean>(false);

    const { auth } = useContext(Context);

    if (auth.isAuth) {
        window.location.assign("/");
    }

    return (
        <div className='container'>
            <Header />
            <Toaster position="top-right" />
            <div className='login-container flex center'>
                <div className='login-block flex b br20 bgc-1'>
                    <div className='inner-login-block'>
                        <div className='login-title'>INFINITY-MC</div>
                        <div className='login-forms-box flex col'>
                            { isReg ? <RegForm /> : <LoginForm /> }
                        </div>
                    </div>
                </div>
                <span className='footer-text' onClick={() => {setIsReg(!isReg)}}>
                    { isReg ? 'Войти?' : 'Зарегистрироваться?' }
                </span>
            </div>
        </div>
    )
}