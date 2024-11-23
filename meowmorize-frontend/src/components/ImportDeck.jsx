// src/pages/ImportPage.jsx
import React, { useState } from 'react';
import { Container, Typography, Button, Box, Alert, CircularProgress } from '@mui/material';
import UploadFileIcon from '@mui/icons-material/UploadFile';
import { importDeck } from '../services/api';

const ImportPage = () => {
  const [selectedFile, setSelectedFile] = useState(null);
  const [message, setMessage] = useState({ type: '', text: '' });
  const [loading, setLoading] = useState(false);

  const handleFileChange = (event) => {
    setSelectedFile(event.target.files[0]);
    setMessage({ type: '', text: '' });
  };

  const handleImport = async () => {
    if (!selectedFile) {
      setMessage({ type: 'error', text: 'Please select a JSON file to import.' });
      return;
    }

    const formData = new FormData();
    formData.append('deck_file', selectedFile);

    try {
      setLoading(true);
      const response = await importDeck(formData);
      setMessage({ type: 'success', text: `Deck "${response.name}" imported successfully!` });
      setSelectedFile(null);
    } catch (error) {
      setMessage({ type: 'error', text: 'Failed to import deck. Please check the file and try again.' });
      console.error(error);
    } finally {
      setLoading(false);
    }
  };

  return (
    <Container maxWidth="sm" sx={{ mt: 4 }}>
      <Typography variant="h4" gutterBottom>
        Import Deck
      </Typography>
      <Box sx={{ display: 'flex', flexDirection: 'column', gap: 2 }}>
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
        {selectedFile && (
          <Typography variant="body1">
            Selected File: {selectedFile.name}
          </Typography>
        )}
        {message.text && (
          <Alert severity={message.type}>{message.text}</Alert>
        )}
        <Button
          variant="contained"
          color="primary"
          onClick={handleImport}
          disabled={loading}
        >
          {loading ? <CircularProgress size={24} /> : 'Import Deck'}
        </Button>
      </Box>
    </Container>
  );
};

export default ImportPage;
