// src/pages/Dashboard.jsx
import React, { useEffect, useState } from 'react';
import { fetchDecks, deleteDeck } from '../services/api';
import {
  Container,
  Typography,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  Paper,
  CircularProgress,
  Alert,
  IconButton,
  Dialog,
  DialogActions,
  DialogContent,
  DialogContentText,
  DialogTitle,
  Button,
  Snackbar,
} from '@mui/material';
import DeleteIcon from '@mui/icons-material/Delete';
import { Link as RouterLink } from 'react-router-dom'; // **Import RouterLink**

const Dashboard = () => {
  const [decks, setDecks] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  // **State for Delete Confirmation Dialog**
  const [openDialog, setOpenDialog] = useState(false);
  const [selectedDeck, setSelectedDeck] = useState(null);

  // **State for Snackbar Notifications**
  const [snackbar, setSnackbar] = useState({
    open: false,
    message: '',
    severity: 'success', // 'success' | 'error' | 'warning' | 'info'
  });

  useEffect(() => {
    const getDecks = async () => {
      try {
        const data = await fetchDecks();
        setDecks(data);
      } catch (err) {
        setError('Failed to fetch decks. Please try again later.');
        console.error(err);
      } finally {
        setLoading(false);
      }
    };

    getDecks();
  }, []);

  // **Handle Open Delete Confirmation Dialog**
  const handleOpenDialog = (deck) => {
    setSelectedDeck(deck);
    setOpenDialog(true);
  };

  // **Handle Close Delete Confirmation Dialog**
  const handleCloseDialog = () => {
    setSelectedDeck(null);
    setOpenDialog(false);
  };

  // **Handle Delete Deck**
  const handleDeleteDeck = async () => {
    if (!selectedDeck) return;

    try {
      await deleteDeck(selectedDeck.id);
      // Remove the deleted deck from the state
      setDecks((prevDecks) => prevDecks.filter((deck) => deck.id !== selectedDeck.id));
      // Show success notification
      setSnackbar({
        open: true,
        message: `Deck "${selectedDeck.name}" deleted successfully.`,
        severity: 'success',
      });
    } catch (err) {
      // Show error notification
      setSnackbar({
        open: true,
        message: 'Failed to delete the deck. Please try again.',
        severity: 'error',
      });
      console.error(err);
    } finally {
      handleCloseDialog();
    }
  };

  // **Handle Close Snackbar**
  const handleCloseSnackbar = (event, reason) => {
    if (reason === 'clickaway') {
      return;
    }
    setSnackbar((prev) => ({ ...prev, open: false }));
  };

  if (loading) {
    return (
      <Container sx={{ mt: 4, textAlign: 'center' }}>
        <CircularProgress />
        <Typography variant="body1">Loading decks...</Typography>
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

  if (decks.length === 0) {
    return (
      <Container sx={{ mt: 4 }}>
        <Alert severity="info">No decks available. Create a new deck to get started!</Alert>
      </Container>
    );
  }

  return (
    <Container sx={{ mt: 4 }}>
      <Typography variant="h4" gutterBottom>
        All Decks
      </Typography>
      <TableContainer component={Paper}>
        <Table aria-label="decks table">
          <TableHead>
            <TableRow>
              <TableCell>Name</TableCell>
              <TableCell>Description</TableCell>
              <TableCell align="right">Number of Cards</TableCell>
              <TableCell align="center">Actions</TableCell> {/* New Actions Column */}
            </TableRow>
          </TableHead>
          <TableBody>
            {decks.map((deck) => (
              <TableRow key={deck.id}>
                <TableCell component="th" scope="row">
                  {/* **Deck Name as a Link to DeckPage** */}
                  <Typography
                    component={RouterLink}
                    to={`/decks/${deck.id}`}
                    variant="body1"
                    color="primary"
                    sx={{ textDecoration: 'none' }}
                  >
                    {deck.name}
                  </Typography>
                </TableCell>
                <TableCell>{deck.description || 'No description provided.'}</TableCell>
                <TableCell align="right">{deck.cards.length}</TableCell>
                <TableCell align="center">
                  {/* **Delete Button** */}
                  <IconButton
                    aria-label="delete"
                    color="error"
                    onClick={() => handleOpenDialog(deck)}
                  >
                    <DeleteIcon />
                  </IconButton>
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </TableContainer>

      {/* **Delete Confirmation Dialog** */}
      <Dialog
        open={openDialog}
        onClose={handleCloseDialog}
        aria-labelledby="delete-dialog-title"
        aria-describedby="delete-dialog-description"
      >
        <DialogTitle id="delete-dialog-title">Delete Deck</DialogTitle>
        <DialogContent>
          <DialogContentText id="delete-dialog-description">
            Are you sure you want to delete the deck "{selectedDeck?.name}"? This action cannot be undone.
          </DialogContentText>
        </DialogContent>
        <DialogActions>
          <Button onClick={handleCloseDialog} color="primary">
            Cancel
          </Button>
          <Button onClick={handleDeleteDeck} color="error" variant="contained" autoFocus>
            Delete
          </Button>
        </DialogActions>
      </Dialog>

      {/* **Snackbar for Notifications** */}
      <Snackbar
        open={snackbar.open}
        autoHideDuration={6000}
        onClose={handleCloseSnackbar}
        anchorOrigin={{ vertical: 'bottom', horizontal: 'center' }}
      >
        <Alert onClose={handleCloseSnackbar} severity={snackbar.severity} sx={{ width: '100%' }}>
          {snackbar.message}
        </Alert>
      </Snackbar>
    </Container>
  );
};

export default Dashboard;
