module.exports = {
    content: [
        "./resources/views/**/*.{html,js}",
    ],
    theme: {
        extend: {},
    },
    plugins: [
        require('@tailwindcss/forms'),
        require('@tailwindcss/aspect-ratio'),
    ],
}
