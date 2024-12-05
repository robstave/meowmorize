// src/App.js

import React, { useState, useMemo, useEffect } from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import { ThemeProvider, CssBaseline } from '@mui/material';
import Navbar from './components/Navbar';
import Dashboard from './pages/Dashboard';
import ImportPage from './pages/ImportPage';
import DeckPage from './pages/DeckPage'; //
import CardPage from './pages/CardPage'; //
import CardForm from './pages/CardForm'; //
import { lightTheme, darkTheme } from './theme';

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


  return (
    <ThemeProvider theme={theme}>
      {/* CssBaseline to apply global styles */}
      <CssBaseline />
      <Router>
      <Navbar mode={mode} toggleTheme={toggleTheme} />
              <Routes>
          <Route path="/" element={<Dashboard />} />
          <Route path="/import" element={<ImportPage />} />
          <Route path="/decks/:id" element={<DeckPage />} /> {/* **New Route for DeckPage** */}
          <Route path="/card/:id" element={<CardPage />} />
          <Route path="/card-form" element={<CardForm />} /> {/* Route for creating a new card */}
          <Route path="/card-form/:id" element={<CardForm />} /> {/* Route for editing an existing card */}


          {/* Add more routes here as you create them */}
        </Routes>
      </Router>
    </ThemeProvider>
  );
}

export default App;
