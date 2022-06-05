import Alpine from 'alpinejs'
import feather from 'feather-icons'

window.feather = feather
window.Alpine = Alpine

const dropdown = () => ({
    open: false,
    trigger() {
        this.open = !this.open
    }
})
Alpine.data('dropdown', dropdown)



const mainMenu = () => ({
    menu: true,
    triggerMenu() {
        this.menu = !this.menu
    }
})
Alpine.data('mainMenu', mainMenu)

Alpine.start()
feather.replace()
