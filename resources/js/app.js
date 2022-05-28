import Alpine from 'alpinejs'

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
