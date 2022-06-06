import { createI18n } from 'vue-i18n'
import enUS from './locales/en-US.json'

// Type-define 'en-US' as the master schema for the resource
type MessageSchema = typeof enUS

const i18n = createI18n<[MessageSchema], 'en-US'>({
    legacy: false,
    locale: 'en-US',
    messages: {
        'en-US': enUS
    }
})

export default i18n
