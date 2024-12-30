// src/pages/DeckPage.jsx
import React, { useEffect, useState, useContext } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import { fetchDeckById, exportDeck, deleteCard, startSession, getNextCard } from '../services/api';
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
  IconButton,
  TableRow,
  Box,
  Paper,
  Button,
  Snackbar,
  Dialog,
  DialogActions,
  DialogContent,
  DialogContentText,
  DialogTitle,
  Collapse,

  Radio,
  RadioGroup,
  FormControlLabel,
  FormControl,
  FormLabel,
  TextField,
  Select,
  MenuItem,

} from '@mui/material';
import { Link as RouterLink } from 'react-router-dom';
import DeleteIcon from '@mui/icons-material/Delete'; // Import DeleteIcon
//import { useNavigate } from 'react-router-dom'; // Import useNavigate
import GetAppIcon from '@mui/icons-material/GetApp'; // Icon for export button
import { saveAs } from 'file-saver'; // To save files on client-side
import ImportMarkdownDialog from '../components/ImportMarkdownDialog'; // Import the dialog
import MuiAlert from '@mui/material/Alert'; // For Snackbar Alert
import { DeckContext } from '../context/DeckContext'; // Import DeckContext


const AlertSnackbar = React.forwardRef(function Alert(props, ref) {
  return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
});



const DeckPage = () => {
  const { id } = useParams(); // Extract the deck ID from the URL
  const { setDecks, loadDecks } = useContext(DeckContext); // Access DeckContext

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
  const [setImportSuccessCount] = useState(0);

  // State for Delete Card Dialog
  const [openDeleteDialog, setOpenDeleteDialog] = useState(false);
  const [cardToDelete, setCardToDelete] = useState(null);

  const [openStartSessionDialog, setOpenStartSessionDialog] = useState(false);
  const [sessionCountOption, setSessionCountOption] = useState('all'); // 'all' or 'count'
  const [sessionCount, setSessionCount] = useState(20); // Default value
  const [sessionMethod, setSessionMethod] = useState('Random'); // Default method

  // State for Snackbar Notifications
  const [snackbar, setSnackbar] = useState({
    open: false,
    message: '',
    severity: 'success', // 'success' | 'error' | 'warning' | 'info'
  });




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

  /*
    // Handler to start a session
    const handleStartSession = async () => {
      try {
        await startSession(id, -1, 'Random'); // Start session with count=-1 and method='Random'
        const nextCardId = await getNextCard(id); // Get the next card ID
        if (nextCardId) {
          navigate(`/card/${nextCardId}`); // Navigate to the card page
        } else {
          setSnackbar({
            open: true,
            message: 'No cards available in this deck.',
            severity: 'info',
          });
        }
      } catch (err) {
        setSnackbar({
          open: true,
          message: 'Failed to start session. Please try again.',
          severity: 'error',
        });
        console.error(err);
      }
    };
    */

  const handleOpenStartSessionDialog = () => {
    setOpenStartSessionDialog(true);
    // Set default session count based on deck size
    if (deck.cards.length <= 20) {
      setSessionCount(deck.cards.length);
    } else {
      setSessionCount(20);
    }
  };

  const handleCloseStartSessionDialog = () => {
    setOpenStartSessionDialog(false);
  };



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
      setSnackbar({
        open: true,
        message: 'Deck exported successfully!',
        severity: 'success',
      });
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
    loadDecks();
    setSnackbar({
      open: true,
      message: `${count} card(s) imported successfully!`,
      severity: 'success',
    });
  };

  // Handlers for Delete Card Dialog
  const handleOpenDeleteDialog = (card) => {
    setCardToDelete(card);
    setOpenDeleteDialog(true);
  };

  const handleCloseDeleteDialog = () => {
    setCardToDelete(null);
    setOpenDeleteDialog(false);
  };

  const handleDeleteCard = async () => {
    if (!cardToDelete) return;

    try {
      await deleteCard(cardToDelete.id);
      // Remove the deleted card from the deck state
      setDeck((prevDeck) => ({
        ...prevDeck,
        cards: prevDeck.cards.filter((card) => card.id !== cardToDelete.id),
      }));
      setSnackbar({
        open: true,
        message: `Card "${cardToDelete.front.text}" deleted successfully.`,
        severity: 'success',
      });
      setDecks((prevDecks) =>
        prevDecks.map((d) =>
          d.id === deck.id ? { ...d, cards: d.cards.filter((c) => c.id !== cardToDelete.id) } : d
        )
      );
    } catch (err) {
      console.error(err);
      setSnackbar({
        open: true,
        message: 'Failed to delete the card. Please try again.',
        severity: 'error',
      });
    } finally {
      handleCloseDeleteDialog();
    }
  };

  // Handler to close the Snackbar
  const handleCloseSnackbar = (event, reason) => {
    if (reason === 'clickaway') {
      return;
    }
    setSnackbar((prev) => ({ ...prev, open: false }));
  };

  const handleStartSessionSubmit = async () => {
    const count = sessionCountOption === 'all' ? -1 : sessionCount;

    try {
      await startSession(id, count, sessionMethod);
      const nextCardId = await getNextCard(id);
      if (nextCardId) {
        navigate(`/card/${nextCardId}`);
      } else {
        setSnackbar({
          open: true,
          message: 'No cards available in this deck.',
          severity: 'info',
        });
      }
    } catch (err) {
      setSnackbar({
        open: true,
        message: 'Failed to start session. Please try again.',
        severity: 'error',
      });
      console.error(err);
    } finally {
      handleCloseStartSessionDialog();
    }
  };


  const handleSessionCountOptionChange = (event) => {
    setSessionCountOption(event.target.value);
  };

  const handleSessionCountChange = (event) => {
    const value = parseInt(event.target.value, 10);
    if (!isNaN(value) && value > 0 && value <= deck.cards.length) {
      setSessionCount(value);
    }
  };

  const handleSessionMethodChange = (event) => {
    setSessionMethod(event.target.value);
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
          onClick={handleOpenStartSessionDialog}
        >
          Start Session
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
                  <TableCell align="center">
                    {/* Delete Button */}
                    <IconButton
                      aria-label="delete"
                      color="error"
                      onClick={() => handleOpenDeleteDialog(card)}
                    >
                      <DeleteIcon />
                    </IconButton>
                  </TableCell>
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


      {/* Delete Card Confirmation Dialog */}
      <Dialog
        open={openDeleteDialog}
        onClose={handleCloseDeleteDialog}
        aria-labelledby="delete-dialog-title"
        aria-describedby="delete-dialog-description"
      >
        <DialogTitle id="delete-dialog-title">Delete Card</DialogTitle>
        <DialogContent>
          <DialogContentText id="delete-dialog-description">
            Are you sure you want to delete the card "{cardToDelete?.front.text}"? This action cannot be undone.
          </DialogContentText>
        </DialogContent>
        <DialogActions>
          <Button onClick={handleCloseDeleteDialog} disabled={!cardToDelete}>
            Cancel
          </Button>
          <Button onClick={handleDeleteCard} color="error" variant="contained" disabled={!cardToDelete}>
            Delete
          </Button>
        </DialogActions>
      </Dialog>

      {/* Start Session Dialog */}
      <Dialog
        open={openStartSessionDialog}
        onClose={handleCloseStartSessionDialog}
        aria-labelledby="start-session-dialog-title"
        aria-describedby="start-session-dialog-description"
      >
        <DialogTitle id="start-session-dialog-title">Start a New Session</DialogTitle>
        <DialogContent>
          <DialogContentText id="start-session-dialog-description">
            Configure the parameters for your new review session.
          </DialogContentText>

          {/* Session Count Option */}
          <FormControl component="fieldset" sx={{ mt: 2 }}>
            <FormLabel component="legend">Number of Cards</FormLabel>
            <RadioGroup
              aria-label="session-count"
              name="session-count"
              value={sessionCountOption}
              onChange={handleSessionCountOptionChange}
            >
              <FormControlLabel value="all" control={<Radio />} label="All Cards" />
              <FormControlLabel value="count" control={<Radio />} label="Specify Count" />
            </RadioGroup>
          </FormControl>

          {/* Session Count Input (Visible only if 'Specify Count' is selected) */}
          {sessionCountOption === 'count' && (
            <TextField
              autoFocus
              margin="dense"
              label="Number of Cards"
              type="number"
              fullWidth
              variant="standard"
              value={sessionCount}
              onChange={handleSessionCountChange}
              inputProps={{
                min: 1,
                max: deck.cards.length,
              }}
              helperText={`Enter a number between 1 and ${deck.cards.length}`}
            />
          )}

          {/* Session Method Selection */}
          <FormControl fullWidth sx={{ mt: 2 }}>
            <FormLabel>Session Method</FormLabel>
            <Select
              value={sessionMethod}
              onChange={handleSessionMethodChange}
              variant="standard"
            >
              <MenuItem value="Random">Random</MenuItem>
              {/* Add more methods here as needed */}
              <MenuItem value="Fails">Fails</MenuItem>
              <MenuItem value="Skips">Skips</MenuItem>
              <MenuItem value="Worst">Worst</MenuItem>
            </Select>
          </FormControl>
        </DialogContent>
        <DialogActions>
          <Button onClick={handleCloseStartSessionDialog} disabled={exporting}>Cancel</Button>
          <Button onClick={handleStartSessionSubmit} variant="contained" color="primary">
            Start Session
          </Button>
        </DialogActions>
      </Dialog>


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

export default DeckPage;
