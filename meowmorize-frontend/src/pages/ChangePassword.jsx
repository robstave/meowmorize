import React, { useState } from 'react';
import { Container, Typography, TextField, Button, Box, Snackbar, Alert } from '@mui/material';
import { changePassword } from '../services/api';

const ChangePassword = () => {
  const [passwordData, setPasswordData] = useState({ password: '', confirm: '' });
  const [snackbar, setSnackbar] = useState({ open: false, message: '', severity: 'success' });

  const handleChange = (e) => {
    setPasswordData({ ...passwordData, [e.target.name]: e.target.value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    if (passwordData.password !== passwordData.confirm) {
      setSnackbar({ open: true, message: 'Passwords do not match', severity: 'error' });
      return;
    }
    try {
      await changePassword(passwordData.password);
      setSnackbar({ open: true, message: 'Password updated', severity: 'success' });
      setPasswordData({ password: '', confirm: '' });
    } catch (err) {
      setSnackbar({ open: true, message: 'Failed to update password', severity: 'error' });
    }
  };

  const handleClose = () => {
    setSnackbar({ ...snackbar, open: false });
  };

  return (
    <Container maxWidth="sm" sx={{ mt: 8 }}>
      <Typography variant="h4" gutterBottom>
        Change Password
      </Typography>
      <Box component="form" onSubmit={handleSubmit} sx={{ display: 'flex', flexDirection: 'column', gap: 2 }}>
        <TextField
          label="New Password"
          name="password"
          type="password"
          value={passwordData.password}
          onChange={handleChange}
          required
        />
        <TextField
          label="Confirm Password"
          name="confirm"
          type="password"
          value={passwordData.confirm}
          onChange={handleChange}
          required
        />
        <Button variant="contained" color="primary" type="submit">
          Update Password
        </Button>
      </Box>
      <Snackbar
        open={snackbar.open}
        autoHideDuration={6000}
        onClose={handleClose}
        anchorOrigin={{ vertical: 'bottom', horizontal: 'center' }}
      >
        <Alert onClose={handleClose} severity={snackbar.severity} sx={{ width: '100%' }}>
          {snackbar.message}
        </Alert>
      </Snackbar>
    </Container>
  );
};

export default ChangePassword;
