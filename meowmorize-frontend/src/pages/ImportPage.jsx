// src/pages/ImportPage.jsx
import React, { useState, useContext } from 'react';
import {
  Container,
  Typography,
  Button,
  Box,
  Alert,
  CircularProgress,
  Snackbar,
  Tooltip,
  Paper,
  Accordion,
  AccordionSummary,
  AccordionDetails,
} from '@mui/material';
import UploadFileIcon from '@mui/icons-material/UploadFile';
import ExpandMoreIcon from '@mui/icons-material/ExpandMore';
import MuiAlert from '@mui/material/Alert';
import ReactMarkdown from 'react-markdown';
import { importDeck, createEmptyDeck, createCard } from '../services/api';
import { DeckContext } from '../context/DeckContext';
import { parseMarkdownToCards } from '../utils/markdownParser';

// Import instructions text from separate files
import markdownInstructions from './markdownInstructions';
import jsonInstructions from './jsonInstructions';

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
  const { loadDecks } = useContext(DeckContext);

  const handleFileChange = (event) => {
    const file = event.target.files[0];
    setSelectedFile(file);
    setMessage({ type: '', text: '' });
    setSnackbar({ open: false, message: '', severity: 'success' });
  };

  const handleImport = async () => {
    if (!selectedFile) {
      setMessage({ type: 'error', text: 'Please select a file to import.' });
      setSnackbar({
        open: true,
        message: 'Please select a file to import.',
        severity: 'error',
      });
      return;
    }

    const fileName = selectedFile.name.toLowerCase();
    setLoading(true);
    setMessage({ type: '', text: '' });

    try {
      // JSON Import: Send the file to your JSON import API
      if (fileName.endsWith('.json')) {
        const formData = new FormData();
        formData.append('deck_file', selectedFile);
        const response = await importDeck(formData);
        setMessage({
          type: 'success',
          text: `Deck "${response.name}" imported successfully!`,
        });
        setSnackbar({
          open: true,
          message: `Deck "${response.name}" imported successfully!`,
          severity: 'success',
        });
      }
      // Markdown Import: Parse the file, create an empty deck, then add cards
      else if (
        fileName.endsWith('.md') ||
        fileName.endsWith('.markdown') ||
        fileName.endsWith('.txt')
      ) {
        const text = await selectedFile.text();
        const cards = parseMarkdownToCards(text);
        if (cards.length === 0) {
          setMessage({
            type: 'error',
            text: 'No valid cards found in the markdown file.',
          });
          setSnackbar({
            open: true,
            message: 'No valid cards found in the markdown file.',
            severity: 'error',
          });
          return;
        }
        // Create a new empty deck first
        const newDeck = await createEmptyDeck();
        // For each parsed card, assign the new deck id and create the card
        for (const card of cards) {
          card.deck_id = newDeck.id;
          await createCard(card);
        }
        setMessage({
          type: 'success',
          text: `Deck "${newDeck.name}" imported successfully with ${cards.length} cards!`,
        });
        setSnackbar({
          open: true,
          message: `Deck "${newDeck.name}" imported successfully with ${cards.length} cards!`,
          severity: 'success',
        });
      }
      // Unsupported File Type
      else {
        setMessage({
          type: 'error',
          text: 'Unsupported file format. Please select a .json or markdown file.',
        });
        setSnackbar({
          open: true,
          message: 'Unsupported file format. Please select a .json or markdown file.',
          severity: 'error',
        });
      }
      // Refresh decks list after a successful import
      loadDecks();
      setSelectedFile(null);
    } catch (error) {
      console.error(error);
      setMessage({
        type: 'error',
        text: 'Failed to import deck. Please check the file and try again.',
      });
      setSnackbar({
        open: true,
        message: 'Failed to import deck. Please check the file and try again.',
        severity: 'error',
      });
    } finally {
      setLoading(false);
    }
  };

  const handleCloseSnackbar = (event, reason) => {
    if (reason === 'clickaway') return;
    setSnackbar((prev) => ({ ...prev, open: false }));
  };

  return (
    <Container maxWidth="sm" sx={{ mt: 4 }}>
      {/* Explanation header */}
      <Typography variant="h4" gutterBottom>
        Import Deck
      </Typography>
      <Typography variant="body1" sx={{ mb: 2 }}>
        Import a deck from either <strong>Markdown</strong> or <strong>JSON</strong> format.
      </Typography>

      <Box sx={{ display: 'flex', flexDirection: 'column', gap: 2 }}>
        {/* Tooltip wrapped file input button */}
        <Tooltip title="Import file from JSON or Markdown format">
          <Button variant="contained" component="label" startIcon={<UploadFileIcon />}>
            Choose File
            <input
              type="file"
              accept=".json,.md,.markdown,.txt"
              hidden
              onChange={handleFileChange}
            />
          </Button>
        </Tooltip>

        {selectedFile && (
          <Typography variant="body1">
            Selected File: {selectedFile.name}
          </Typography>
        )}

        {message.text && <Alert severity={message.type}>{message.text}</Alert>}

        <Button
          variant="contained"
          color="primary"
          onClick={handleImport}
          disabled={loading}
        >
          {loading ? <CircularProgress size={24} /> : 'Import Deck'}
        </Button>
      </Box>

      {/* Accordion for instructions */}
      <Box sx={{ mt: 4 }}>
        <Typography variant="h6" gutterBottom>
          Import File Types
        </Typography>

        <Accordion>
          <AccordionSummary expandIcon={<ExpandMoreIcon />}>
            <Typography>Markdown</Typography>
          </AccordionSummary>
          <AccordionDetails>
            <ReactMarkdown>{markdownInstructions}</ReactMarkdown>
          </AccordionDetails>
        </Accordion>

        <Accordion>
          <AccordionSummary expandIcon={<ExpandMoreIcon />}>
            <Typography>JSON</Typography>
          </AccordionSummary>
          <AccordionDetails>
            <ReactMarkdown>{jsonInstructions}</ReactMarkdown>
          </AccordionDetails>
        </Accordion>
      </Box>

      <Snackbar
        open={snackbar.open}
        autoHideDuration={6000}
        onClose={handleCloseSnackbar}
        anchorOrigin={{ vertical: 'bottom', horizontal: 'center' }}
      >
        <AlertSnackbar
          onClose={handleCloseSnackbar}
          severity={snackbar.severity}
          sx={{ width: '100%' }}
        >
          {snackbar.message}
        </AlertSnackbar>
      </Snackbar>
    </Container>
  );
};

export default ImportPage;
