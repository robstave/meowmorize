// src/pages/ImportPage.jsx
import React, { useState, useContext } from 'react';
import { Container, Typography, Button, Box, Alert, CircularProgress } from '@mui/material';
import { Snackbar } from '@mui/material';

import UploadFileIcon from '@mui/icons-material/UploadFile';
import { importDeck } from '../services/api';
import { DeckContext } from '../context/DeckContext'; // Import DeckContext

import MuiAlert from '@mui/material/Alert'; // For Snackbar Alert

const AlertSnackbar = React.forwardRef(function Alert(props, ref) {
  return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
});


const ImportPage = () => {
  const [selectedFile, setSelectedFile] = useState(null);
  const [message, setMessage] = useState({ type: '', text: '' });
  const [loading, setLoading] = useState(false);


  const [snackbar, setSnackbar] = useState({
    open: false,
    message: '',
    severity: 'success',
  });

  const { loadDecks } = useContext(DeckContext); // Access DeckContext


  // Handle file selection
  const handleFileChange = (event) => {
    setSelectedFile(event.target.files[0]);
    setMessage({ type: '', text: '' }); // Reset messages on new file selection
    setSnackbar({ open: false, message: '', severity: 'success' }); // Reset messages on new file selection

  };

  // Handle deck import
  const handleImport = async () => {
    if (!selectedFile) {
      setMessage({ type: 'error', text: 'Please select a JSON file to import.' });
      setSnackbar({ open: true, message: 'Please select a JSON file to import.', severity: 'error' });

      return;
    }

    const formData = new FormData();
    formData.append('deck_file', selectedFile);

    try {
      setLoading(true);
      const response = await importDeck(formData);
      setSnackbar({ open: true, message: `Deck "${response.name}" imported successfully!`, severity: 'success' });

      setMessage({ type: 'success', text: `Deck "${response.name}" imported successfully!` });
      setSelectedFile(null); // Reset file input
      loadDecks();
    } catch (error) {
      setMessage({ type: 'error', text: 'Failed to import deck. Please check the file and try again.' });
      console.error(error);
    } finally {
      setLoading(false);
    }
  };

  // Handler to close the Snackbar
  const handleCloseSnackbar = (event, reason) => {
    if (reason === 'clickaway') {
      return;
    }
    setSnackbar((prev) => ({ ...prev, open: false }));
  };


  return (
    <Container maxWidth="sm" sx={{ mt: 4 }}>
      <Typography variant="h4" gutterBottom>
        Import Deck
      </Typography>
      <Box sx={{ display: 'flex', flexDirection: 'column', gap: 2 }}>
        {/* File Upload Button */}
        <Button
          variant="contained"
          component="label"
          startIcon={<UploadFileIcon />}
        >
          Choose JSON File
          <input
            type="file"
            accept=".json"
            hidden
            onChange={handleFileChange}
          />
        </Button>

        {/* Display Selected File Name */}
        {selectedFile && (
          <Typography variant="body1">
            Selected File: {selectedFile.name}
          </Typography>
        )}

        {/* Display Success or Error Message */}
        {message.text && (
          <Alert severity={message.type}>{message.text}</Alert>
        )}

        {/* Import Button */}
        <Button
          variant="contained"
          color="primary"
          onClick={handleImport}
          disabled={loading}
        >
          {loading ? <CircularProgress size={24} /> : 'Import Deck'}
        </Button>
      </Box>

      {/* Snackbar for Notifications */}
      <Snackbar
        open={snackbar.open}
        autoHideDuration={6000}
        onClose={handleCloseSnackbar}
        anchorOrigin={{ vertical: 'bottom', horizontal: 'center' }}
      >
        <AlertSnackbar onClose={handleCloseSnackbar} severity={snackbar.severity} sx={{ width: '100%' }}>
          {snackbar.message}
        </AlertSnackbar>
      </Snackbar>


    </Container>
  );
};

export default ImportPage;
