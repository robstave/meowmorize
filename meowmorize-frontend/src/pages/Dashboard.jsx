// src/pages/Dashboard.jsx
import React, { useState, useContext } from 'react';
import { deleteDeck, updateDeck } from '../services/api';
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
  TableSortLabel,
} from '@mui/material';
import DeleteIcon from '@mui/icons-material/Delete';
import EditIcon from '@mui/icons-material/Edit';
import { Link as RouterLink } from 'react-router-dom';
import { DeckContext } from '../context/DeckContext';
import { formatLastAccessed } from '../utils/dateUtils';

function descendingComparator(a, b, orderBy) {
  let valueA, valueB;
  switch (orderBy) {
    case 'name':
      valueA = a.name.toLowerCase();
      valueB = b.name.toLowerCase();
      break;
    case 'description':
      valueA = a.description ? a.description.toLowerCase() : '';
      valueB = b.description ? b.description.toLowerCase() : '';
      break;
    case 'last_accessed':
      valueA = new Date(a.last_accessed);
      valueB = new Date(b.last_accessed);
      break;
    case 'cards':
      valueA = a.cards.length;
      valueB = b.cards.length;
      break;
    default:
      return 0;
  }
  if (valueB < valueA) return -1;
  if (valueB > valueA) return 1;
  return 0;
}

function getComparator(order, orderBy) {
  return order === 'desc'
    ? (a, b) => descendingComparator(a, b, orderBy)
    : (a, b) => -descendingComparator(a, b, orderBy);
}

function stableSort(array, comparator) {
  const stabilizedThis = array.map((el, index) => [el, index]);
  stabilizedThis.sort((a, b) => {
    const orderResult = comparator(a[0], b[0]);
    if (orderResult !== 0) return orderResult;
    return a[1] - b[1];
  });
  return stabilizedThis.map((el) => el[0]);
}

const Dashboard = () => {
  const { decks, setDecks, loading, error } = useContext(DeckContext);
  const [openEditDialog, setOpenEditDialog] = useState(false);
  const [selectedDeck, setSelectedDeck] = useState(null);
  const [editName, setEditName] = useState('');
  const [editDescription, setEditDescription] = useState('');
  const [openDeleteDialog, setOpenDeleteDialog] = useState(false);
  const [deckToDelete, setDeckToDelete] = useState(null);
  const [snackbar, setSnackbar] = useState({
    open: false,
    message: '',
    severity: 'success',
  });

  // Sorting state
  const [order, setOrder] = useState('asc');
  const [orderBy, setOrderBy] = useState('name');

  const handleRequestSort = (property) => {
    const isAsc = orderBy === property && order === 'asc';
    setOrder(isAsc ? 'desc' : 'asc');
    setOrderBy(property);
  };

  const handleOpenEditDialog = (deck) => {
    setSelectedDeck(deck);
    setEditName(deck.name);
    setEditDescription(deck.description || '');
    setOpenEditDialog(true);
  };

  const handleCloseEditDialog = () => {
    setSelectedDeck(null);
    setEditName('');
    setEditDescription('');
    setOpenEditDialog(false);
  };

  const handleEditDeck = async () => {
    if (!selectedDeck) return;

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

      const response = await updateDeck(selectedDeck.id, updatedDeck);

      setDecks((prevDecks) =>
        prevDecks.map((deck) =>
          deck.id === selectedDeck.id ? response : deck
        )
      );

      setSnackbar({
        open: true,
        message: 'Deck updated successfully!',
        severity: 'success',
      });

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

  const handleOpenDeleteDialog = (deck) => {
    setDeckToDelete(deck);
    setOpenDeleteDialog(true);
  };

  const handleCloseDeleteDialog = () => {
    setDeckToDelete(null);
    setOpenDeleteDialog(false);
  };

  const handleDeleteDeck = async () => {
    if (!deckToDelete) return;

    try {
      await deleteDeck(deckToDelete.id);
      setDecks((prevDecks) =>
        prevDecks.filter((deck) => deck.id !== deckToDelete.id)
      );
      setSnackbar({
        open: true,
        message: `Deck "${deckToDelete.name}" deleted successfully.`,
        severity: 'success',
      });
    } catch (err) {
      setSnackbar({
        open: true,
        message: 'Failed to delete the deck. Please try again.',
        severity: 'error',
      });
      console.error(err);
    } finally {
      handleCloseDeleteDialog();
    }
  };

  const handleCloseSnackbar = (event, reason) => {
    if (reason === 'clickaway') return;
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

  // Sort decks using the current sort parameters
  const sortedDecks = stableSort(decks, getComparator(order, orderBy));

  return (
    <Container sx={{ mt: 4 }}>
      <Typography variant="h4" gutterBottom>
        All Decks
      </Typography>
      <TableContainer component={Paper}>
        <Table aria-label="decks table">
          <TableHead>
            <TableRow>
              <TableCell>
                <TableSortLabel
                  active={orderBy === 'name'}
                  direction={orderBy === 'name' ? order : 'asc'}
                  onClick={() => handleRequestSort('name')}
                >
                  Name
                </TableSortLabel>
              </TableCell>
              <TableCell>
                <TableSortLabel
                  active={orderBy === 'description'}
                  direction={orderBy === 'description' ? order : 'asc'}
                  onClick={() => handleRequestSort('description')}
                >
                  Description
                </TableSortLabel>
              </TableCell>
              <TableCell align="right">
                <TableSortLabel
                  active={orderBy === 'last_accessed'}
                  direction={orderBy === 'last_accessed' ? order : 'asc'}
                  onClick={() => handleRequestSort('last_accessed')}
                >
                  Last Session
                </TableSortLabel>
              </TableCell>
              <TableCell align="right">
                <TableSortLabel
                  active={orderBy === 'cards'}
                  direction={orderBy === 'cards' ? order : 'asc'}
                  onClick={() => handleRequestSort('cards')}
                >
                  Cards
                </TableSortLabel>
              </TableCell>
              <TableCell align="center">Actions</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {sortedDecks.map((deck) => (
              <TableRow key={deck.id}>
                <TableCell component="th" scope="row">
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
                <TableCell align="right">{formatLastAccessed(deck.last_accessed)}</TableCell>
                <TableCell align="right">{deck.cards.length}</TableCell>
                <TableCell align="center">
                  <IconButton
                    aria-label="edit"
                    color="primary"
                    onClick={() => handleOpenEditDialog(deck)}
                  >
                    <EditIcon />
                  </IconButton>
                  <IconButton
                    aria-label="delete"
                    color="error"
                    onClick={() => handleOpenDeleteDialog(deck)}
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

      {/* Delete Confirmation Dialog */}
      <Dialog
        open={openDeleteDialog}
        onClose={handleCloseDeleteDialog}
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
          <Button onClick={handleCloseDeleteDialog} color="primary">
            Cancel
          </Button>
          <Button onClick={handleDeleteDeck} color="error" variant="contained" autoFocus>
            Delete
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
        <Alert onClose={handleCloseSnackbar} severity={snackbar.severity} sx={{ width: '100%' }}>
          {snackbar.message}
        </Alert>
      </Snackbar>
    </Container>
  );
};

export default Dashboard;
