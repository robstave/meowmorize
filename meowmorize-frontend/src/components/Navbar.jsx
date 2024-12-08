// src/components/Navbar.jsx
import React from 'react';
import { AppBar, Toolbar, Typography, Button, Switch, FormControlLabel, Snackbar } from '@mui/material';
import { Link as RouterLink } from 'react-router-dom';
import Logo from '../logo512.png'; // Adjust the path if necessary

import { useNavigate } from 'react-router-dom';
import { createEmptyDeck } from '../services/api'; // Import the API function
import MuiAlert from '@mui/material/Alert'; // For Snackbar Alert


// Inside your Navbar component
const Navbar = ({ mode, toggleTheme }) => {
  const [snackbar, setSnackbar] = React.useState({
    open: false,
    message: '',
    severity: 'success', // 'success' | 'error' | 'warning' | 'info'
  });

  // Inside the Navbar component
  const navigate = useNavigate();

  // Handler to close the Snackbar
  const handleCloseSnackbar = (event, reason) => {
    if (reason === 'clickaway') {
      return;
    }
    setSnackbar((prev) => ({ ...prev, open: false }));
  };


  // Handler to create an empty deck
  const handleCreateEmptyDeck = async () => {
    try {
      const newDeck = await createEmptyDeck();
      setSnackbar({
        open: true,
        message: `Empty deck "${newDeck.name}" created successfully!`,
        severity: 'success',
      });
      // Optional: Navigate to the new deck's page
      navigate(`/decks/${newDeck.id}`);
    } catch (error) {
      console.error('Failed to create empty deck:', error);
      setSnackbar({
        open: true,
        message: 'Failed to create empty deck. Please try again.',
        severity: 'error',
      });
    }
  };


  return (
    <AppBar position="static">
      <Toolbar>
        {/* App Logo or Icon */}
        <img src={Logo} alt="MeowMorize Logo" style={{ height: '40px', marginRight: '16px' }} />        {/* App Name */}
        <Typography color="inherit" variant="h6" component={RouterLink}
          to="/" sx={{ flexGrow: 1 }}>
          MeowMorize
        </Typography>
        {/* Navigation Links */}
        <Button color="inherit" component={RouterLink} to="/">
          Dashboard
        </Button>
        <Button color="inherit" component={RouterLink} to="/import">
          Import Deck
        </Button>

        {/* Create Empty Deck Button */}
        <Button color="inherit" onClick={handleCreateEmptyDeck}>
          Create Empty Deck
        </Button>


        {/* Theme Toggle Switch */}
        <FormControlLabel
          control={
            <Switch
              checked={mode === 'dark'}
              onChange={toggleTheme}
              color="default"
            />
          }
          label="Dark Mode"
        />

        {/* Add more navigation links here as needed */}
      </Toolbar>
        {/* Snackbar for Notifications */}
        <Snackbar
        open={snackbar.open}
        autoHideDuration={6000}
        onClose={handleCloseSnackbar}
        anchorOrigin={{ vertical: 'bottom', horizontal: 'center' }}
      >
        <MuiAlert onClose={handleCloseSnackbar} severity={snackbar.severity} sx={{ width: '100%' }}>
          {snackbar.message}
        </MuiAlert>
      </Snackbar>
      
    </AppBar>
  );
};

export default Navbar;
