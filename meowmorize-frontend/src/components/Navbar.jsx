// src/components/Navbar.jsx
import React from 'react';
import { AppBar, Toolbar, Typography, Button } from '@mui/material';
import { Link as RouterLink } from 'react-router-dom';
import Logo from '../logo512.png'; // Adjust the path if necessary
const Navbar = () => {
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
        {/* Add more navigation links here as needed */}
      </Toolbar>
    </AppBar>
  );
};

export default Navbar;
