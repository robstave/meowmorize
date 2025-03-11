// src/pages/DeckPage.jsx
import React, { useEffect, useState, useContext } from 'react';
import { useParams, useNavigate, Link as RouterLink } from 'react-router-dom';
import { getSessionOverview } from '../services/api';
import { format } from 'date-fns';

import {
  fetchDeckById,
  exportDeck,
  deleteCard,
  startSession,
  getNextCard,
  updateDeck,
} from '../services/api';
import { formatLastAccessed } from '../utils/dateUtils';
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
  Tooltip,
  Menu,
} from '@mui/material';
import MoreVertIcon from '@mui/icons-material/MoreVert';
import DeleteIcon from '@mui/icons-material/Delete';
import GetAppIcon from '@mui/icons-material/GetApp';
import { saveAs } from 'file-saver';
import ImportMarkdownDialog from '../components/ImportMarkdownDialog';
import MuiAlert from '@mui/material/Alert';
import { DeckContext } from '../context/DeckContext';

const AlertSnackbar = React.forwardRef(function Alert(props, ref) {
  return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
});

// Utility to truncate a string to n characters, appending "..." if needed.
const truncate = (str, n) =>
  str.length > n ? `${str.substring(0, n)}...` : str;

const DeckPage = () => {
  const { id } = useParams();
  const { setDecks, loadDecks } = useContext(DeckContext);
  const [deck, setDeck] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  // State to manage the visibility of the cards list
  const [showCards, setShowCards] = useState(false);

  // Export Dialog states
  const [openExportDialog, setOpenExportDialog] = useState(false);
  const [exporting, setExporting] = useState(false);
  const [exportError, setExportError] = useState('');

  const navigate = useNavigate();

  // Import Markdown Dialog state
  const [openImportDialog, setOpenImportDialog] = useState(false);
  const [importSuccessCount, setImportSuccessCount] = useState(0);

  // Delete Card Dialog states
  const [openDeleteDialog, setOpenDeleteDialog] = useState(false);
  const [cardToDelete, setCardToDelete] = useState(null);

  // Start Session Dialog states
  const [openStartSessionDialog, setOpenStartSessionDialog] = useState(false);
  const [sessionCountOption, setSessionCountOption] = useState('all'); // 'all' or 'count'
  const [sessionCount, setSessionCount] = useState(20); // Default value
  const [sessionMethod, setSessionMethod] = useState('Random');

  // States for inline editing of deck name
  const [isEditing, setIsEditing] = useState(false);
  const [newDeckName, setNewDeckName] = useState('');

  // States for inline editing of deck description
  const [isEditingDescription, setIsEditingDescription] = useState(false);
  const [newDeckDescription, setNewDeckDescription] = useState('');

  const [sessions, setSessions] = useState([]);


  // Snackbar Notification state
  const [snackbar, setSnackbar] = useState({
    open: false,
    message: '',
    severity: 'success',
  });

  // New state for the More Actions menu (three dots)
  const [menuAnchorEl, setMenuAnchorEl] = useState(null);
  const handleMenuOpen = (event) => {
    setMenuAnchorEl(event.currentTarget);
  };
  const handleMenuClose = () => {
    setMenuAnchorEl(null);
  };



  useEffect(() => {
    const getDeckAndSessions = async () => {
      try {
        const [deckData, sessionData] = await Promise.all([
          fetchDeckById(id),
          getSessionOverview(id),
        ]);
        setDeck(deckData);
        setSessions(sessionData);
      } catch (err) {
        setError('Failed to fetch deck details or session data.');
        console.error(err);
      } finally {
        setLoading(false);
      }
    };

    getDeckAndSessions();
  }, [id]);


  const handleOpenStartSessionDialog = () => {
    setOpenStartSessionDialog(true);
    if (deck.cards.length <= 20) {
      setSessionCount(deck.cards.length);
    } else {
      setSessionCount(20);
    }
  };

  const handleCloseStartSessionDialog = () => {
    setOpenStartSessionDialog(false);
  };

  const handleToggleCards = () => {
    setShowCards((prev) => !prev);
  };

  // Handlers for deck name editing
  const handleDeckNameClick = () => {
    setIsEditing(true);
    setNewDeckName(deck.name);
  };

  const handleDeckNameChange = (event) => {
    setNewDeckName(event.target.value);
  };

  const handleDeckNameSave = async () => {
    if (newDeckName.trim() === '') {
      setSnackbar({
        open: true,
        message: 'Deck name cannot be empty.',
        severity: 'error',
      });
      return;
    }

    try {
      const updatedDeck = await updateDeck(id, { name: newDeckName, description: deck.description });
      setDeck(updatedDeck);
      setDecks((prevDecks) =>
        prevDecks.map((d) => (d.id === id ? updatedDeck : d))
      );
      setSnackbar({
        open: true,
        message: 'Deck name updated successfully!',
        severity: 'success',
      });
      setIsEditing(false);
    } catch (err) {
      console.error(err);
      setSnackbar({
        open: true,
        message: 'Failed to update deck name. Please try again.',
        severity: 'error',
      });
    }
  };

  const handleDeckNameCancel = () => {
    setIsEditing(false);
    setNewDeckName('');
  };

  // Handlers for deck description editing
  const handleDescriptionClick = () => {
    setIsEditingDescription(true);
    setNewDeckDescription(deck.description || '');
  };

  const handleDeckDescriptionChange = (event) => {
    setNewDeckDescription(event.target.value);
  };

  const handleDeckDescriptionSave = async () => {
    try {
      const updatedDeck = await updateDeck(id, { name: deck.name, description: newDeckDescription });
      setDeck(updatedDeck);
      setDecks((prevDecks) =>
        prevDecks.map((d) => (d.id === id ? updatedDeck : d))
      );
      setSnackbar({
        open: true,
        message: 'Deck description updated successfully!',
        severity: 'success',
      });
      setIsEditingDescription(false);
    } catch (err) {
      console.error(err);
      setSnackbar({
        open: true,
        message: 'Failed to update deck description. Please try again.',
        severity: 'error',
      });
    }
  };

  const handleDeckDescriptionCancel = () => {
    setIsEditingDescription(false);
    setNewDeckDescription('');
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

  const handleImportSuccess = async (count) => {
    setImportSuccessCount(count);
    loadDecks(); // Update context if needed
    try {
      const updatedDeck = await fetchDeckById(id);
      setDeck(updatedDeck);
    } catch (error) {
      console.error('Failed to refresh deck after import:', error);
    }
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

  const handleCloseSnackbar = (event, reason) => {
    if (reason === 'clickaway') return;
    setSnackbar((prev) => ({ ...prev, open: false }));
  };

  const handleStartSessionSubmit = async () => {
    const count = sessionCountOption === 'all' ? -1 : sessionCount;
    try {
      await startSession(id, count, sessionMethod);
      const nextCardId = await getNextCard(id);
      if (nextCardId) {
        //navigate(`/card/${nextCardId}`);
        // navigate to card page
        navigate(`/decks/${id}/card/${nextCardId}`);

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

  const getColorForPercentage = (value) => {
    if (value >= 85) return '#4caf50';  // Green for 85%+
    if (value >= 70) return '#ffeb3b';  // Yellow for 70%-85%
    return '#f44336';                 // Red for below 70%
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
      <Box sx={{ display: 'flex', alignItems: 'center' }}>
        {!isEditing ? (
          <Tooltip title="Click to edit deck name">
            <Typography
              variant="h4"
              gutterBottom
              onClick={handleDeckNameClick}
              sx={{ cursor: 'pointer', '&:hover': { textDecoration: 'underline' } }}
            >
              {deck.name}
            </Typography>
          </Tooltip>
        ) : (
          <Box sx={{ display: 'flex', alignItems: 'center' }}>
            <TextField
              variant="standard"
              value={newDeckName}
              onChange={handleDeckNameChange}
              label="Deck Name"
              sx={{ mr: 2 }}
            />
            <Button variant="contained" color="primary" onClick={handleDeckNameSave} sx={{ mr: 1 }}>
              Save
            </Button>
            <Button variant="outlined" color="secondary" onClick={handleDeckNameCancel}>
              Cancel
            </Button>
          </Box>
        )}
      </Box>

      {/* Editable Deck Description */}
      <Box sx={{ mt: 1 }}>
        {!isEditingDescription ? (
          <Tooltip title="Click to edit deck description">
            <Typography
              variant="body1"
              gutterBottom
              onClick={handleDescriptionClick}
              sx={{ cursor: 'pointer', '&:hover': { textDecoration: 'underline' } }}
            >
              {deck.description || 'Click to add description'}
            </Typography>
          </Tooltip>
        ) : (
          <Box sx={{ display: 'flex', alignItems: 'center' }}>
            <TextField
              variant="standard"
              fullWidth
              value={newDeckDescription}
              onChange={handleDeckDescriptionChange}
              label="Deck Description"
            />
            <Button variant="contained" color="primary" onClick={handleDeckDescriptionSave} sx={{ ml: 1 }}>
              Save
            </Button>
            <Button variant="outlined" color="secondary" onClick={handleDeckDescriptionCancel} sx={{ ml: 1 }}>
              Cancel
            </Button>
          </Box>
        )}
      </Box>

      {/* Last Accessed & Number of Cards */}
      <Typography variant="subtitle1" color="textSecondary" gutterBottom>
        Last Session: {formatLastAccessed(deck.last_accessed)}
      </Typography>
      <Typography variant="subtitle1" gutterBottom>
        Number of Cards: {deck.cards.length}
      </Typography>

      {/* Session Overview Table */}
      <Typography variant="h6" sx={{ mt: 4, mb: 2 }}>
        Last 3 Sessions
      </Typography>

 
      {/* Action Buttons */}
      <Box sx={{ display: 'flex', gap: 2, mt: 4,  mb: 2, alignItems: 'center' }}>
        <Tooltip title="Start a new review session using this deck">
          <Button variant="contained" color="primary" onClick={handleOpenStartSessionDialog}>
            Start Session
          </Button>
        </Tooltip>
        <Tooltip title="Toggle preview of deck cards">
          <Button variant="contained" color="primary" onClick={handleToggleCards}>
            {showCards ? 'Hide Preview' : 'Show Preview'}
          </Button>
        </Tooltip>
        <Tooltip title="More actions">
          <IconButton onClick={handleMenuOpen}>
            <MoreVertIcon />
          </IconButton>
        </Tooltip>
      </Box>
      <Menu
        anchorEl={menuAnchorEl}
        open={Boolean(menuAnchorEl)}
        onClose={handleMenuClose}
      >
        <MenuItem component={RouterLink} to="/" onClick={handleMenuClose}>
          Back to Dashboard
        </MenuItem>
        <MenuItem onClick={() => { handleOpenExportDialog(); handleMenuClose(); }}>
          Export Deck
        </MenuItem>
        <MenuItem onClick={() => { handleOpenImportDialog(); handleMenuClose(); }}>
          Import from Markdown
        </MenuItem>
      </Menu>

      {/* Cards List */}
      <Collapse in={showCards}>
        <TableContainer component={Paper} sx={{ mt: 2 }}>
          <Table aria-label="cards table">
            <TableHead>
              <TableRow>
                <TableCell>Front</TableCell>
                <TableCell>Back</TableCell>
                <TableCell>Stars</TableCell>
                <TableCell align="center">Actions</TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {deck.cards.map((card) => (
                <TableRow key={card.id}>
                  <TableCell>
                    <Tooltip title={card.front.text}>
                      <span>{truncate(card.front.text, 150)}</span>
                    </Tooltip>
                  </TableCell>
                  <TableCell>
                    <Tooltip title={card.back.text}>
                      <span>{truncate(card.back.text, 150)}</span>
                    </Tooltip>
                  </TableCell>
                  <TableCell>{card.star_rating || 0}</TableCell>
                  <TableCell align="center">
                    <Tooltip title="Delete this card">
                      <IconButton aria-label="delete" color="error" onClick={() => handleOpenDeleteDialog(card)}>
                        <DeleteIcon />
                      </IconButton>
                    </Tooltip>
                  </TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </TableContainer>
      </Collapse>



      {sessions.length === 0 ? (
        <Typography variant="body2">No recent sessions available.</Typography>
      ) : (
        <TableContainer component={Paper} elevation={2}>
          <Table size="small">
            <TableHead>
              <TableRow>
                <TableCell><strong>Date</strong></TableCell>
                <TableCell align="right"><strong>Initial %</strong></TableCell>
                <TableCell align="right"><strong>Final %</strong></TableCell>
                <TableCell align="right"><strong>Cards Attempted</strong></TableCell>
                <TableCell align="right"><strong>Cards Final</strong></TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {sessions.map((session) => (
                <TableRow key={session.sessionid}>
                  <TableCell>{format(new Date(session.timestamp), 'PPPp')}</TableCell>
                  <TableCell
                    align="right"
                    style={{ backgroundColor: getColorForPercentage(session.percentage) }}
                  >{session.percentage.toFixed(1)}%</TableCell>
                  <TableCell
                    align="right"
                    style={{ backgroundColor: getColorForPercentage(session.percentage_after) }}
                  >{session.percentage_after.toFixed(1)}%</TableCell>
                  <TableCell align="right">{session.cards}</TableCell>
                  <TableCell align="right">{session.cards_after}</TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </TableContainer>
      )}


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
          <Button onClick={handleCloseExportDialog} disabled={exporting}>
            Cancel
          </Button>
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
          <FormControl fullWidth sx={{ mt: 2 }}>
            <FormLabel>Session Method</FormLabel>
            <Select value={sessionMethod} onChange={handleSessionMethodChange} variant="standard">
              <MenuItem value="Random">Random</MenuItem>
              <MenuItem value="Fails">Fails</MenuItem>
              <MenuItem value="Skips">Skips</MenuItem>
              <MenuItem value="Worst">Worst</MenuItem>
              <MenuItem value="Stars">Stars</MenuItem>
              <MenuItem value="Unrated">Unrated</MenuItem>
              <MenuItem value="AdjustedRandom">AdjustedRandom</MenuItem>
            </Select>
          </FormControl>
        </DialogContent>
        <DialogActions>
          <Button onClick={handleCloseStartSessionDialog} disabled={exporting}>
            Cancel
          </Button>
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
