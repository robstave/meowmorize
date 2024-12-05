// src/theme.js
import { createTheme } from '@mui/material/styles';

// Define the light theme
export const lightTheme = createTheme({
  palette: {
    mode: 'light',
    primary: {
      main: '#1976d2', // You can customize this color
    },
    background: {
      default: '#f5f5f5',
    },
  },
});

// Define the dark theme
export const darkTheme = createTheme({
  palette: {
    mode: 'dark',
    primary: {
      main: '#90caf9', // You can customize this color
    },
    background: {
      default: '#121212',
    },
  },
});
