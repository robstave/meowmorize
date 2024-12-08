// src/pages/Dashboard.jsx
import React, { useEffect, useState } from 'react';
import { fetchDecks, deleteDeck, updateDeck } from '../services/api';
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
  TextField,
} from '@mui/material';
import DeleteIcon from '@mui/icons-material/Delete';
import EditIcon from '@mui/icons-material/Edit'; // Import EditIcon
import { Link as RouterLink } from 'react-router-dom'; // **Import RouterLink**

const Dashboard = () => {
  const [decks, setDecks] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  // State for Edit Dialog
  const [openEditDialog, setOpenEditDialog] = useState(false);
  const [selectedDeck, setSelectedDeck] = useState(null);
  const [editName, setEditName] = useState('');
  const [editDescription, setEditDescription] = useState('');


  // **State for Delete Confirmation Dialog**
  const [openDialog, setOpenDialog] = useState(false);
  const [deckToDelete, setDeckToDelete] = useState(null);
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

  // Function to open the edit dialog
  const handleOpenEditDialog = (deck) => {
    setSelectedDeck(deck);
    setEditName(deck.name);
    setEditDescription(deck.description || '');
    setOpenEditDialog(true);
  };

  // Function to close the edit dialog
  const handleCloseEditDialog = () => {
    setSelectedDeck(null);
    setEditName('');
    setEditDescription('');
    setOpenEditDialog(false);
  };

  const handleEditDeck = async () => {
    if (!selectedDeck) return;

    // Basic validation
    if (editName.trim() === '') {
      setSnackbar({
        open: true,
        message: 'Deck name cannot be empty.',
        severity: 'error',
      });
      return;
    }

    try {
      const updatedDeck = {
        name: editName,
        description: editDescription,
      };

      // Call the API to update the deck
      const response = await updateDeck(selectedDeck.id, updatedDeck);

      // Update the decks state
      setDecks((prevDecks) =>
        prevDecks.map((deck) =>
          deck.id === selectedDeck.id ? response : deck
        )
      );

      // Show success notification
      setSnackbar({
        open: true,
        message: 'Deck updated successfully!',
        severity: 'success',
      });

      // Close the dialog
      handleCloseEditDialog();
    } catch (err) {
      console.error(err);
      setSnackbar({
        open: true,
        message: 'Failed to update deck. Please try again.',
        severity: 'error',
      });
    }
  };


  // **Handle Open Delete Confirmation Dialog**
  const handleOpenDialog = (deck) => {
    setDeckToDelete(deck);
    setOpenDialog(true);
  };

  // **Handle Close Delete Confirmation Dialog**
  const handleCloseDialog = () => {
    setDeckToDelete(null);
    setOpenDialog(false);
  };

  // **Handle Delete Deck**
  const handleDeleteDeck = async () => {
    if (!deckToDelete) return;

    try {
      await deleteDeck(deckToDelete.id);
      // Remove the deleted deck from the state
      setDecks((prevDecks) => prevDecks.filter((deck) => deck.id !== deckToDelete.id));
      // Show success notification
      setSnackbar({
        open: true,
        message: `Deck "${deckToDelete.name}" deleted successfully.`,
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

                  <IconButton
                    aria-label="edit"
                    color="primary"
                    onClick={() => handleOpenEditDialog(deck)}
                  >
                    <EditIcon />
                  </IconButton>


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


      {/* Edit Deck Dialog */}
      <Dialog
        open={openEditDialog}
        onClose={handleCloseEditDialog}
        aria-labelledby="edit-deck-dialog-title"
        aria-describedby="edit-deck-dialog-description"
      >
        <DialogTitle id="edit-deck-dialog-title">Edit Deck</DialogTitle>
        <DialogContent>
          <DialogContentText id="edit-deck-dialog-description">
            Update the name and description of the deck.
          </DialogContentText>
          <TextField
            autoFocus
            margin="dense"
            label="Deck Name"
            type="text"
            fullWidth
            variant="standard"
            value={editName}
            onChange={(e) => setEditName(e.target.value)}
          />
          <TextField
            margin="dense"
            label="Description"
            type="text"
            fullWidth
            variant="standard"
            value={editDescription}
            onChange={(e) => setEditDescription(e.target.value)}
            multiline
            rows={4}
          />
        </DialogContent>
        <DialogActions>
          <Button onClick={handleCloseEditDialog}>Cancel</Button>
          <Button onClick={handleEditDeck} color="primary">
            Save
          </Button>
        </DialogActions>
      </Dialog>


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
            Are you sure you want to delete the deck "{deckToDelete?.name}"? This action cannot be undone.
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
