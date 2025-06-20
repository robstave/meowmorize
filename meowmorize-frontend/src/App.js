// src/App.js

import React, { useMemo, useContext, useEffect } from 'react';
import {   useNavigate } from 'react-router-dom';

import { HashRouter as Router, Routes, Route, Navigate } from 'react-router-dom';
import { ThemeProvider, CssBaseline } from '@mui/material';
import Navbar from './components/Navbar';
import Dashboard from './pages/Dashboard';
import ImportPage from './pages/ImportPage';
import DeckPage from './pages/DeckPage'; //
import CardPage from './pages/CardPage'; //
import CardForm from './pages/CardForm'; //
import UserManagement from './pages/UserManagement'; // Import UserManagement page
import ChangePassword from './pages/ChangePassword';
import { lightTheme, darkTheme } from './theme';
import LoginPage from './pages/LoginPage'; // Import LoginPage
import { AuthContext } from './context/AuthContext'; // Import AuthContext
import { ThemeContext } from './context/ThemeContext'; // Import ThemeContext
import 'highlight.js/styles/atom-one-dark.css'; // or choose your preferred theme



// New component for redirecting to Swagger UI
function SwaggerRedirect() {
  const navigate = useNavigate();

  useEffect(() => {
    navigate("/swagger/index.html#/", { replace: true });
  }, [navigate]);

  return null; // This component doesn't render anything
}

function App() {

  const { isDarkMode } = useContext(ThemeContext); // Access ThemeContext

  // Memoize the MUI theme to optimize performance
  const theme = useMemo(() => (isDarkMode ? darkTheme : lightTheme), [isDarkMode]);


  const { auth } = useContext(AuthContext);

  return (
    <ThemeProvider theme={theme}>
      {/* CssBaseline to apply global styles */}
      <CssBaseline />
      <Router>
        <Navbar />
        <Routes>
          {/* Public Route */}
          <Route path="/login" element={!auth.token ? <LoginPage /> : <Navigate to="/" />} />

          {/* Protected Routes */}
          <Route path="/" element={auth.token ? <Dashboard /> : <Navigate to="/login" />} />
          <Route path="/import" element={auth.token ? <ImportPage /> : <Navigate to="/login" />} />
          <Route path="/decks/:id" element={auth.token ? <DeckPage /> : <Navigate to="/login" />} />
          <Route path="/decks/:deckId/card/:id" element={auth.token ? <CardPage /> : <Navigate to="/login" />} />
          <Route path="/card-form" element={auth.token ? <CardForm /> : <Navigate to="/login" />} />
          <Route path="/card-form/:id" element={auth.token ? <CardForm /> : <Navigate to="/login" />} />
          <Route
            path="/admin/users"
            element={
              auth.token && auth.user?.role === 'admin' ? (
                <UserManagement />
              ) : (
                <Navigate to="/" />
              )
            }
          />
          <Route path="/password" element={auth.token ? <ChangePassword /> : <Navigate to="/login" />} />


          {/* Route to handle Swagger UI redirect */}
          <Route path="/swagger" element={<SwaggerRedirect />} />


          {/* Redirect any unknown routes to login */}
          <Route path="*" element={<Navigate to="/login" />} />
        </Routes>

      </Router>
    </ThemeProvider>
  );
}

export default App;
