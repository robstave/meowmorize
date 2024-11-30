// src/pages/DeckPage.jsx
import React, { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import { fetchDeckById, exportDeck } from '../services/api';
import {
  Container,
  Typography,
  CircularProgress,
  Alert,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  Box,
  Paper,
  Button,
  Dialog,
  DialogActions,
  DialogContent,
  DialogContentText,
  DialogTitle,
  Collapse,

} from '@mui/material';
import { Link as RouterLink } from 'react-router-dom';


import { useNavigate } from 'react-router-dom'; // Import useNavigate

import GetAppIcon from '@mui/icons-material/GetApp'; // Icon for export button
import { saveAs } from 'file-saver'; // To save files on client-side
import ImportMarkdownDialog from '../components/ImportMarkdownDialog'; // Import the dialog





const DeckPage = () => {
  const { id } = useParams(); // Extract the deck ID from the URL
  const [deck, setDeck] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  // State to manage the visibility of the cards list
  const [showCards, setShowCards] = useState(false);

  // State for Export Dialog
  const [openExportDialog, setOpenExportDialog] = useState(false);
  const [exporting, setExporting] = useState(false);
  const [exportError, setExportError] = useState('');

  const navigate = useNavigate(); // Initialize navigate

  // State for Import Markdown Dialog
  const [openImportDialog, setOpenImportDialog] = useState(false);
  const [importSuccessCount, setImportSuccessCount] = useState(0);

  // Function to select a random card
  const selectCard = () => {
    console.log('Selecting card');
    if (deck.cards.length === 0) return;

    const randomCard = deck.cards[Math.floor(Math.random() * deck.cards.length)];
    navigate(`/card/${randomCard.id}`);
  };


  useEffect(() => {
    const getDeck = async () => {
      try {
        const data = await fetchDeckById(id);
        setDeck(data);
      } catch (err) {
        setError('Failed to fetch deck details. Please try again later.');
        console.error(err);
      } finally {
        setLoading(false);
      }
    };

    getDeck();
  }, [id]);

  // Handler to toggle the visibility of the cards list
  const handleToggleCards = () => {
    setShowCards((prev) => !prev);
  };


  // Handlers for Export Dialog
  const handleOpenExportDialog = () => {
    setExportError('');
    setOpenExportDialog(true);
  };

  const handleCloseExportDialog = () => {
    if (!exporting) {
      setOpenExportDialog(false);
    }
  };

  // Function to handle exporting the deck
  const handleExportDeck = async () => {
    try {
      setExporting(true);
      const blob = await exportDeck(id);
      const deckNameSanitized = deck.name.replace(/[^a-z0-9]/gi, '_').toLowerCase();
      saveAs(blob, `${deckNameSanitized}.json`);
      setOpenExportDialog(false);
    } catch (err) {
      setExportError('Failed to export deck. Please try again.');
      console.error(err);
    } finally {
      setExporting(false);
    }
  };

  
  // Handlers for Import Markdown Dialog
  const handleOpenImportDialog = () => {
    setOpenImportDialog(true);
  };

  const handleCloseImportDialog = () => {
    setOpenImportDialog(false);
  };

  const handleImportSuccess = (count) => {
    setImportSuccessCount(count);
    // Refresh the deck to include new cards
    const refreshDeck = async () => {
      try {
        const data = await fetchDeckById(id);
        setDeck(data);
      } catch (err) {
        console.error('Failed to refresh deck after import', err);
      }
    };
    refreshDeck();
  };



  if (loading) {
    return (
      <Container sx={{ mt: 4, textAlign: 'center' }}>
        <CircularProgress />
        <Typography variant="body1">Loading deck details...</Typography>
      </Container>
    );
  }

  if (error) {
    return (
      <Container sx={{ mt: 4 }}>
        <Alert severity="error">{error}</Alert>
      </Container>
    );
  }

  if (!deck) {
    return (
      <Container sx={{ mt: 4 }}>
        <Alert severity="info">Deck not found.</Alert>
      </Container>
    );
  }

  return (
    <Container sx={{ mt: 4 }}>
      {/* Deck Title */}
      <Typography variant="h4" gutterBottom>
        {deck.name}
      </Typography>

      {/* Number of Cards */}
      <Typography variant="subtitle1" gutterBottom>
        Number of Cards: {deck.cards.length}
      </Typography>

      {/* Deck Description */}
      {deck.description && (
        <Typography variant="body1" gutterBottom>
          {deck.description}
        </Typography>
      )}

      <Box sx={{ display: 'flex', gap: 2, mt: 4 }}>

        <Button
          variant="contained"
          color="primary"
          onClick={selectCard}

        >
          View Random Card
        </Button>


        {/* Preview Button to Toggle Cards List */}
        <Button
          variant="contained"
          color="primary"
          onClick={handleToggleCards}

        >
          {showCards ? 'Hide Preview' : 'Show Preview'}
        </Button>

        {/* Back to Dashboard Button */}
        <Button
          component={RouterLink}
          to="/"
          variant="outlined"
          color="secondary"

        >
          Back to Dashboard
        </Button>

        {/* Export Deck Button */}
        <Button
          variant="outlined"
          color="secondary"
          startIcon={<GetAppIcon />}
          onClick={handleOpenExportDialog}
        >
          Export Deck
        </Button>

        {/* Import Markdown Button */}
        <Button
          variant="outlined"
          color="primary"
          onClick={handleOpenImportDialog}
        >
          Import from Markdown
        </Button>

      </Box>
      {/* Cards List (Conditionally Rendered) */}
      <Collapse in={showCards}>
        <TableContainer component={Paper} sx={{ mt: 2 }}>
          <Table aria-label="cards table">
            <TableHead>
              <TableRow>
                <TableCell>Front</TableCell>
                <TableCell>Back</TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {deck.cards.map((card) => (
                <TableRow key={card.id}>
                  <TableCell>{card.front.text}</TableCell>
                  <TableCell>{card.back.text}</TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </TableContainer>
      </Collapse>

     {/* Export Deck Confirmation Dialog */}
     <Dialog
        open={openExportDialog}
        onClose={handleCloseExportDialog}
        aria-labelledby="export-dialog-title"
        aria-describedby="export-dialog-description"
      >
        <DialogTitle id="export-dialog-title">Export Deck</DialogTitle>
        <DialogContent>
          <DialogContentText id="export-dialog-description">
            Do you want to export the deck "{deck.name}" as a JSON file?
          </DialogContentText>
          {exportError && <Alert severity="error" sx={{ mt: 2 }}>{exportError}</Alert>}
        </DialogContent>
        <DialogActions>
          <Button onClick={handleCloseExportDialog} disabled={exporting}>Cancel</Button>
          <Button onClick={handleExportDeck} color="secondary" variant="contained" disabled={exporting}>
            {exporting ? 'Exporting...' : 'Export'}
          </Button>
        </DialogActions>
      </Dialog>

      {/* Import Markdown Dialog */}
      <ImportMarkdownDialog
        open={openImportDialog}
        handleClose={handleCloseImportDialog}
        deckId={deck.id}
        onImportSuccess={handleImportSuccess}
      />

      {/* Snackbar for Import Success */}
      {importSuccessCount > 0 && (
        <Alert severity="success" sx={{ mt: 2 }}>
          Successfully imported {importSuccessCount} card{importSuccessCount > 1 ? 's' : ''}.
        </Alert>
      )}

    </Container>
  );
};

export default DeckPage;
