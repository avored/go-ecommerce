import * as React from 'react'
import { IsUserLoggedInStatus, ThemeContextType } from '../@types/theme'

export const ThemeContext = React.createContext<ThemeContextType | null>(null);

const ThemeProvider: React.FC<React.ReactNode> = ({ children }) => {
  const [isLoggedInUserStatus, setIsLoggedInUserStatus] = React.useState<IsUserLoggedInStatus>(false)

  return <ThemeContext.Provider value={{
    isUserLoggedInStatus: isLoggedInUserStatus,
    updateIsLoggedInUserStatus: setIsLoggedInUserStatus
  }}>
    {children}
  </ThemeContext.Provider>;
};

export default ThemeProvider;
