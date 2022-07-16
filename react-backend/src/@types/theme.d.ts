export type IsUserLoggedInStatus = boolean

export type ThemeContextType = {
    isUserLoggedInStatus: boolean
    updateIsLoggedInUserStatus: (userStatus: boolean) => void
}
