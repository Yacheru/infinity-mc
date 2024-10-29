import i18n from 'i18next';
import backend from 'i18next-http-backend'
import LanguageDetector from 'i18next-browser-languagedetector'
import { initReactI18next } from 'react-i18next'

i18n
    .use(backend)
    .use(LanguageDetector)
    .use(initReactI18next)
    .init({
        fallbackLng: 'ru',
        debug: true,
        interpolation: {
            escapeValue: true
        },
        backend: {
            loadPath: "./src/assets/locales/{{lng}}/translation.json"
        }
})
