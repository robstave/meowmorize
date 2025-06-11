// src/components/Navbar.jsx
import React, { useContext } from 'react';
import { useEffect, useState } from 'react';
import {
  AppBar, Toolbar, Typography,
  Button, Switch, FormControlLabel,
  Menu,
  MenuItem,
  Snackbar,
  IconButton,
} from '@mui/material';
import { Link as RouterLink, useNavigate } from 'react-router-dom';
import Logo from '../logo512.png'; // Adjust the path if necessary
import MenuIcon from '@mui/icons-material/Menu';

import { createEmptyDeck, fetchDecks } from '../services/api';

import MuiAlert from '@mui/material/Alert'; // For Snackbar Alert
import CollapseDecksDialog from './CollapseDecksDialog'; // Import the new dialog component


import { AuthContext } from '../context/AuthContext'; // Import AuthContext
import { DeckContext } from '../context/DeckContext'; // Import DeckContext
import { ThemeContext } from '../context/ThemeContext'; // Import ThemeContext



// Inside your Navbar component
const Navbar = () => {
  const { isDarkMode, toggleTheme } = useContext(ThemeContext); // Access ThemeContext

  const { decks, setDecks, loadDecks  } = useContext(DeckContext);
  const { auth, logout } = useContext(AuthContext);
  const [anchorEl, setAnchorEl] = React.useState(null);
  const open = Boolean(anchorEl);

  const [collapseDialogOpen, setCollapseDialogOpen] = useState(false);

  const [snackbar, setSnackbar] = React.useState({
    open: false,
    message: '',
    severity: 'success', // 'success' | 'error' | 'warning' | 'info'
  });

  // Inside the Navbar component
  const navigate = useNavigate();


  useEffect(() => {
    const getDecks = async () => {
      try {
        const data = await fetchDecks();
        setDecks(data);
      } catch (err) {
        console.error('Failed to fetch decks:', err);
      }
    };

    if (auth.token) {
      getDecks();
    }
  }, [auth.token]);







  // Handler to open the dropdown menu
  const handleMenuOpen = (event) => {
    setAnchorEl(event.currentTarget);
  };

  // Handler to close the dropdown menu
  const handleMenuClose = () => {
    setAnchorEl(null);
  };

  // Handler for logout
  const handleLogout = () => {
    logout();
    navigate('/login');
  };


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
      // Refresh decks after creation
      loadDecks();
      // Navigate to the new deck's page
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





  const handleOpenCollapseDialog = () => {
    setCollapseDialogOpen(true);
    handleMenuClose();
  };

  const handleCloseCollapseDialog = () => {
    setCollapseDialogOpen(false);
  };



  const handleCollapseSuccess = (message) => {
    setSnackbar({
      open: true,
      message,
      severity: 'success',
    });
    // Refresh decks after collapse
    loadDecks();
  };




  const handleCollapseError = (message) => {
    setSnackbar({
      open: true,
      message,
      severity: 'error',
    });
  };


  return (
    <AppBar position="static">
      <Toolbar>
        {/* App Logo or Icon */}
        <img src={Logo} alt="MeowMorize Logo" style={{ height: '40px', marginRight: '16px' }} />        {/* App Name */}
        <Typography color="inherit" variant="h6" component={RouterLink}
        onClick={loadDecks} 
          to="/" sx={{ flexGrow: 1 }}>
          MeowMorize
        </Typography>
        {/* Navigation Links */}
        {auth.token && (
          <Button color="inherit" component={RouterLink} to="/">
            Dashboard
          </Button>
        )}

        {/* Dropdown Menu for Deck Actions */}
        {auth.token && (
          <div>

            <IconButton
              color="inherit"
              onClick={handleMenuOpen}
              aria-controls={open ? 'deck-menu' : undefined}
              aria-haspopup="true"
              aria-expanded={open ? 'true' : undefined}
            >
              <MenuIcon />
            </IconButton>
            <Menu
              id="deck-menu"
              anchorEl={anchorEl}
              open={open}
              onClose={handleMenuClose}
              MenuListProps={{
                'aria-labelledby': 'deck-menu-button',
              }}
            >
              <MenuItem component={RouterLink} to="/import" onClick={handleMenuClose}>
                Import Deck
              </MenuItem>
              <MenuItem onClick={handleCreateEmptyDeck}>Create Empty Deck</MenuItem>
              <MenuItem onClick={handleOpenCollapseDialog}>Collapse Decks</MenuItem>
              {auth.user && auth.user.role === 'admin' && (
                <MenuItem component={RouterLink} to="/admin/users" onClick={handleMenuClose}>
                  User Management
                </MenuItem>
              )}

            </Menu>
          </div>
        )}


        {/* Theme Toggle Switch */}
        <FormControlLabel
          control={
            <Switch
              checked={isDarkMode}
              onChange={toggleTheme}
              color="default"
            />
          }
          label="Dark Mode"
        />

        {/* Authentication Buttons */}
        {!auth.token ? (
          <Button color="inherit" component={RouterLink} to="/login">
            Login
          </Button>
        ) : (
          <Button color="inherit" onClick={handleLogout}>
            Logout
          </Button>
        )}
        {/* Add more navigation links here as needed */}
      </Toolbar>

      {/* Collapse Decks Dialog */}
      <CollapseDecksDialog
        open={collapseDialogOpen}
        handleClose={handleCloseCollapseDialog}
        decks={decks}
        onSuccess={handleCollapseSuccess}
        onError={handleCollapseError}
      />


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
