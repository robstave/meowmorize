// src/App.js

import React, { useState, useMemo, useEffect, useContext  } from 'react';
import { BrowserRouter as Router, Routes, Route, Navigate } from 'react-router-dom';
import { ThemeProvider, CssBaseline } from '@mui/material';
import Navbar from './components/Navbar';
import Dashboard from './pages/Dashboard';
import ImportPage from './pages/ImportPage';
import DeckPage from './pages/DeckPage'; //
import CardPage from './pages/CardPage'; //
import CardForm from './pages/CardForm'; //
import { lightTheme, darkTheme } from './theme';

 
 
import LoginPage from './pages/LoginPage'; // Import LoginPage
 import { AuthContext } from './context/AuthContext'; // Import AuthContext



function App() {

  // Check localStorage for theme preference
  const storedTheme = localStorage.getItem('theme') || 'light';
  const [mode, setMode] = useState(storedTheme);

  // Update localStorage whenever theme changes
  useEffect(() => {
    localStorage.setItem('theme', mode);
  }, [mode]);

  // Memoize the theme to optimize performance
  const theme = useMemo(() => (mode === 'light' ? lightTheme : darkTheme), [mode]);

  // Function to toggle theme
  const toggleTheme = () => {
    setMode((prevMode) => (prevMode === 'light' ? 'dark' : 'light'));
  };

  const { auth } = useContext(AuthContext);
  
  return (
    <ThemeProvider theme={theme}>
      {/* CssBaseline to apply global styles */}
      <CssBaseline />
      <Router>
      <Navbar mode={mode} toggleTheme={toggleTheme} />
      <Routes>
          {/* Public Route */}
          <Route path="/login" element={!auth.token ? <LoginPage /> : <Navigate to="/" />} />

          {/* Protected Routes */}
          <Route path="/" element={auth.token ? <Dashboard /> : <Navigate to="/login" />} />
          <Route path="/import" element={auth.token ? <ImportPage /> : <Navigate to="/login" />} />
          <Route path="/decks/:id" element={auth.token ? <DeckPage /> : <Navigate to="/login" />} />
          <Route path="/card/:id" element={auth.token ? <CardPage /> : <Navigate to="/login" />} />
          <Route path="/card-form" element={auth.token ? <CardForm /> : <Navigate to="/login" />} />
          <Route path="/card-form/:id" element={auth.token ? <CardForm /> : <Navigate to="/login" />} />

          {/* Redirect any unknown routes to login */}
          <Route path="*" element={<Navigate to="/login" />} />
        </Routes>
        
      </Router>
    </ThemeProvider>
  );
}

export default App;
