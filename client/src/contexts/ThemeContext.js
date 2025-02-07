import { createContext, useState, useEffect, useContext } from "react";

const ThemeContext = createContext();

export const ThemeProvider = ({ children }) => {
  const storedMode = localStorage.getItem("themeMode");

  const [isDarkMode, setIsDarkMode] = useState(storedMode === "dark");

  const toggleDarkMode = () => {
    setIsDarkMode((prev) => !prev);
    localStorage.setItem("themeMode", isDarkMode ? "light" : "dark");
  };

  useEffect(() => {
    const storedMode = localStorage.getItem("themeMode");
    setIsDarkMode(storedMode === "dark");
  }, []);

  return (
    <ThemeContext.Provider value={{ isDarkMode, toggleDarkMode }}>
      {children}
    </ThemeContext.Provider>
  );
};

export const useTheme = () => useContext(ThemeContext);
