// src/pages/ImportPage.jsx
import React, { useState, useContext } from 'react';
import {
  Container,
  Typography,
  Button,
  Box,
  Alert,
  CircularProgress,
  TextField,
  Snackbar,
  Tooltip,
  Accordion,
  AccordionSummary,
  AccordionDetails,
} from '@mui/material';
import UploadFileIcon from '@mui/icons-material/UploadFile';
import ExpandMoreIcon from '@mui/icons-material/ExpandMore';
import MuiAlert from '@mui/material/Alert';
import ReactMarkdown from 'react-markdown';
import {
  importDeck,
  createEmptyDeck,
  createCard,
  updateDeck,
} from '../services/api';
import { DeckContext } from '../context/DeckContext';
import { parseMarkdownToCards } from '../utils/markdownParser';
// Optionally import instructions for markdown/json if needed
import markdownInstructions from './markdownInstructions';
import jsonInstructions from './jsonInstructions';

const AlertSnackbar = React.forwardRef(function Alert(props, ref) {
  return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
});

const ImportPage = () => {
  // New state for deck title, defaulting to "Default Deck"
  const [deckTitle, setDeckTitle] = useState('Default Deck');
  const [selectedFile, setSelectedFile] = useState(null);
  const [message, setMessage] = useState({ type: '', text: '' });
  const [loading, setLoading] = useState(false);
  const [snackbar, setSnackbar] = useState({
    open: false,
    message: '',
    severity: 'success',
  });
  const { loadDecks } = useContext(DeckContext);

  const handleFileSelect = async (event) => {
    const file = event.target.files[0];
    if (!file) return;
    setSelectedFile(file);
    setMessage({ type: '', text: '' });
    setSnackbar({ open: false, message: '', severity: 'success' });
    await handleImport(file);
  };

  const handleImport = async (file = selectedFile) => {
    if (!file) {
      setMessage({ type: 'error', text: 'Please select a file to import.' });
      setSnackbar({
        open: true,
        message: 'Please select a file to import.',
        severity: 'error',
      });
      return;
    }

    const fileName = file.name.toLowerCase();
    setLoading(true);
    setMessage({ type: '', text: '' });

    try {
      // JSON Import: Send the file to your JSON import API
      if (fileName.endsWith('.json')) {
        const formData = new FormData();
        formData.append('deck_file', file);
        const response = await importDeck(formData);
        // If a custom title was provided, update the deck
        if (deckTitle.trim() !== 'Default Deck') {
          await updateDeck(response.id, { name: deckTitle });
        }
        setMessage({
          type: 'success',
          text: `Deck "${deckTitle}" imported successfully!`,
        });
        setSnackbar({
          open: true,
          message: `Deck "${deckTitle}" imported successfully!`,
          severity: 'success',
        });
      }
      // Markdown Import: Parse the file, create an empty deck, then add cards
      else if (
        fileName.endsWith('.md') ||
        fileName.endsWith('.markdown') ||
        fileName.endsWith('.txt')
      ) {
        const text = await file.text();
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
        // Create a new empty deck
        const newDeck = await createEmptyDeck();
        // If the title is different, update the new deck’s name
        if (deckTitle.trim() !== 'Default Deck') {
          await updateDeck(newDeck.id, { name: deckTitle });
        }
        // Create each card for the new deck
        for (const card of cards) {
          card.deck_id = newDeck.id;
          await createCard(card, newDeck.id);
        }
        setMessage({
          type: 'success',
          text: `Deck "${deckTitle}" imported successfully with ${cards.length} cards!`,
        });
        setSnackbar({
          open: true,
          message: `Deck "${deckTitle}" imported successfully with ${cards.length} cards!`,
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
          message:
            'Unsupported file format. Please select a .json or markdown file.',
          severity: 'error',
        });
      }
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
      <Typography variant="h4" gutterBottom>
        Import Deck
      </Typography>
      <Typography variant="body1" sx={{ mb: 2 }}>
        Import a deck from either <strong>Markdown</strong> or{' '}
        <strong>JSON</strong> format.
      </Typography>
      <Box sx={{ display: 'flex', flexDirection: 'column', gap: 2 }}>
        {/* Editable Deck Title Input */}
        <TextField
          label="Deck Title"
          value={deckTitle}
          onChange={(e) => setDeckTitle(e.target.value)}
          fullWidth
        />
        <Tooltip title="Import file from JSON or Markdown format">
          <Button
            variant="contained"
            component="label"
            startIcon={<UploadFileIcon />}
            disabled={loading}
          >
            {loading ? <CircularProgress size={24} /> : 'Import Deck'}
            <input
              type="file"
              accept=".json,.md,.markdown,.txt"
              hidden
              onChange={handleFileSelect}
            />
          </Button>
        </Tooltip>
        {message.text && <Alert severity={message.type}>{message.text}</Alert>}
      </Box>
      {/* Optionally include additional instructions via Accordion */}
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
