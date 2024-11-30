import React, { useState } from 'react';
import {
  Dialog,
  DialogTitle,
  DialogContent,
  DialogActions,
  Button,
  Typography,
  LinearProgress,
  Alert,
} from '@mui/material';
import { styled } from '@mui/system';
import { parseMarkdownToCards } from '../utils/markdownParser'; // We'll create this utility
import { createCard } from '../services/api';
import { v4 as uuidv4 } from 'uuid'; // To generate unique IDs for frontend previews

const Input = styled('input')({
  display: 'none',
});

const ImportMarkdownDialog = ({ open, handleClose, deckId, onImportSuccess }) => {
  const [selectedFile, setSelectedFile] = useState(null);
  const [parsedCards, setParsedCards] = useState([]);
  const [importing, setImporting] = useState(false);
  const [error, setError] = useState('');
  const [previewReady, setPreviewReady] = useState(false);

  const handleFileChange = async (event) => {
    const file = event.target.files[0];
    setError('');
    setParsedCards([]);
    setPreviewReady(false);

    console.log('imporing');
    if (file) {
        const allowedExtensions = ['.md', '.markdown', '.txt'];
        const fileName = file.name.toLowerCase();
        console.log(fileName);
        const hasValidExtension = allowedExtensions.some(ext => fileName.endsWith(ext));
        console.log(hasValidExtension);

        
        if (!hasValidExtension) {
          setError('Please upload a valid Markdown file with .md, .markdown, or .txt extension.');
          return;
        }

      try {
        const text = await file.text();
        const cards = parseMarkdownToCards(text);
        setParsedCards(cards);
        setPreviewReady(true);
      } catch (err) {
        console.error(err);
        setError('Failed to parse the Markdown file. Please check the format.');
      }
    }
  };

  const handleImport = async () => {
    setImporting(true);
    setError('');

    try {
      for (const card of parsedCards) {
        // Assign a new UUID if not present (backend should handle this ideally)
        const newCard = {
          ...card,
          id: uuidv4(), // Temporary ID for frontend; backend should generate the actual ID
          deck_id: deckId,
        };
        await createCard(newCard);
      }
      setImporting(false);
      onImportSuccess(parsedCards.length);
      handleClose();
    } catch (err) {
      console.error(err);
      setError('Failed to import some or all cards. Please try again.');
      setImporting(false);
    }
  };

  const handleDialogClose = () => {
    if (!importing) {
      setSelectedFile(null);
      setParsedCards([]);
      setError('');
      setPreviewReady(false);
      handleClose();
    }
  };

  return (
    <Dialog open={open} onClose={handleDialogClose} maxWidth="sm" fullWidth>
      <DialogTitle>Import Cards from Markdown</DialogTitle>
      <DialogContent dividers>
        {error && <Alert severity="error" sx={{ mb: 2 }}>{error}</Alert>}
        {!parsedCards.length && (
          <>
            <label htmlFor="markdown-upload">
              <Input accept=".md,.markdown,.txt" id="markdown-upload" type="file" onChange={handleFileChange} />
              <Button variant="contained" component="span">
                Choose Markdown File
              </Button>
            </label>
          </>
        )}
        {previewReady && (
          <Typography variant="body1">
            {parsedCards.length} card{parsedCards.length > 1 ? 's' : ''} detected for import.
          </Typography>
        )}
        {importing && <LinearProgress sx={{ mt: 2 }} />}
      </DialogContent>
      <DialogActions>
        <Button onClick={handleDialogClose} disabled={importing}>Cancel</Button>
        {previewReady && (
          <Button onClick={handleImport} variant="contained" color="primary" disabled={importing}>
            {importing ? 'Importing...' : 'Accept'}
          </Button>
        )}
      </DialogActions>
    </Dialog>
  );
};

export default ImportMarkdownDialog;
