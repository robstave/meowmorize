// src/context/ThemeContext.js

import React, { createContext, useState, useEffect } from 'react';

export const ThemeContext = createContext();

export const ThemeProvider = ({ children }) => {
  // Initialize theme based on localStorage or default to 'light'
  const storedTheme = localStorage.getItem('theme') || 'light';
  const [isDarkMode, setIsDarkMode] = useState(storedTheme === 'dark');

  // Update localStorage whenever the theme changes
  useEffect(() => {
    localStorage.setItem('theme', isDarkMode ? 'dark' : 'light');
  }, [isDarkMode]);

  // Function to toggle theme
  const toggleTheme = () => {
    setIsDarkMode((prevMode) => !prevMode);
  };

  return (
    <ThemeContext.Provider value={{ isDarkMode, toggleTheme }}>
      {children}
    </ThemeContext.Provider>
  );
};
