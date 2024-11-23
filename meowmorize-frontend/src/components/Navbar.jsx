// src/components/Navbar.jsx
import React from 'react';
import { AppBar, Toolbar, Typography, Button } from '@mui/material';
import ImportExportIcon from '@mui/icons-material/ImportExport';
import { Link as RouterLink } from 'react-router-dom';

const Navbar = () => {
  return (
    <AppBar position="static">
      <Toolbar>
        {/* App Logo or Icon */}
        <ImportExportIcon sx={{ mr: 2 }} />
        {/* App Name */}
        <Typography variant="h6" component="div" sx={{ flexGrow: 1 }}>
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
