import React, { useState, useEffect } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import { fetchCardById, fetchDeckById, deleteCard } from '../services/api';
import { updateCardStats } from '../services/api';

import {
  Container,
  Typography,
  Button,
  Box,
  IconButton,
  Dialog,
  DialogTitle,
  DialogContent,
  DialogActions,
  DialogContentText,
  Snackbar,
  CircularProgress,
  Alert,

} from '@mui/material';
import ReactMarkdown from 'react-markdown';
import seedrandom from 'seedrandom';
import { Link as RouterLink } from 'react-router-dom';

import DeleteIcon from '@mui/icons-material/Delete'; // Import DeleteIcon
import MuiAlert from '@mui/material/Alert'; // For Snackbar Alert

import HorizontalStatusBar from '../components/HorizontalStatusBar';



const AlertSnackbar = React.forwardRef(function Alert(props, ref) {
  return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
});





const rng = seedrandom(Date.now().toString()); // Use a seed for predictable results

const CardPage = () => {
  const { id } = useParams(); // Get card ID from URL
  const navigate = useNavigate();
  const [card, setCard] = useState(null);
  const [deckId, setDeckId] = useState(null); // Track the deckId
  const [showFront, setShowFront] = useState(true);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [history, setHistory] = useState([]); // History of recently used indices

  const [passCount, setPassCount] = useState(0);
const [skipCount, setSkipCount] = useState(0);
const [failCount, setFailCount] = useState(0);



  // State for Delete Card Dialog
  const [openDeleteDialog, setOpenDeleteDialog] = useState(false);

  // State for Snackbar Notifications
  const [snackbar, setSnackbar] = useState({
    open: false,
    message: '',
    severity: 'success', // 'success' | 'error' | 'warning' | 'info'
  });

  // Fetch the card on component mount or when the ID changes
  useEffect(() => {
    const getCard = async () => {
      try {
        const data = await fetchCardById(id);
        setCard(data);
        setDeckId(data.deck_id); // Set the deckId for future fetches
        setPassCount(data.pass_count);
        setSkipCount(data.skip_count);
        setFailCount(data.fail_count);
        

      } catch (err) {
        setError('Failed to fetch card details. Please try again later.');
        console.error(err);
      } finally {
        setLoading(false);
      }
    };

    getCard();
  }, [id]);

  // Utility function to get a unique random index
  const getUniqueRandomIndex = (deckLength) => {
    if (deckLength <= history.length) {
      console.warn('Deck does not have enough unique cards to avoid repeats.');
      return Math.floor(rng() * deckLength); // Fallback: pick any random index
    }

    let randomIndex;
    do {
      randomIndex = Math.floor(rng() * deckLength);
    } while (history.includes(randomIndex));

    // Update history
    setHistory((prevHistory) => {
      const updatedHistory = [...prevHistory, randomIndex];
      if (updatedHistory.length > 2) {
        updatedHistory.shift(); // Keep only the last two indices
      }
      return updatedHistory;
    });

    return randomIndex;
  };

// Function to handle card actions: pass, fail, skip
  const handleCardAction = async (action) => {
    if (!card) return;

    try {
      // Send the action to the backend

      switch(action) {
        case 'pass':
          await updateCardStats(card.id, 'IncrementPass');
          break;
        case 'skip':
          await updateCardStats(card.id, 'IncrementSkip');

          break;
          case 'fail':
            await updateCardStats(card.id, 'IncrementFail');

            break;
        default:
          // code block
          console.log('other click')
      } 
     

      // Fetch the deck to get all cards
      const deck = await fetchDeckById(deckId);

      if (deck.cards.length === 0) {
        setError('No cards available in the deck.');
        return;
      }

      // Select a new random card
      const randomIndex = getUniqueRandomIndex(deck.cards.length);
      const randomCard = deck.cards[randomIndex];

      setShowFront(true); // Reset to show the front of the new card

      // Navigate to the new card
      navigate(`/card/${randomCard.id}`);
    } catch (err) {
      console.error('Failed to process the card action:', err);
      setSnackbar({
        open: true,
        message: 'Failed to process the action. Please try again.',
        severity: 'error',
      });
    }
  };


  // Function to select a random card from the current deck
  const selectCard = async (action) => {

    if (!deckId) {
      console.log('Deck ID is not available');
      setError('Deck ID is not available.');
      return;
    }

    //if (action == 'pass') {
    // when I call this I get an endless looping
    // of next cards.
    //  await updateCardStats(card.id, 'IncrementPass');
    //};
    //
    // I would like a switch with
    // pass -> IncrementPass
    //fail => incrementFail
    //skip -> incrementSkip
    

    try {
      const deck = await fetchDeckById(deckId);
      const randomIndex = getUniqueRandomIndex(deck.cards.length);
      const randomCard = deck.cards[randomIndex];
      setShowFront(true); // Reset to show the front of the new card
      navigate(`/card/${randomCard.id}`);
    } catch (err) {
      setError('Failed to fetch a random card.');
      console.error(err);
    }
  };


  // Handlers for Delete Card Dialog
  const handleOpenDeleteDialog = () => {
    setOpenDeleteDialog(true);
  };

  const handleCloseDeleteDialog = () => {
    setOpenDeleteDialog(false);
  };

  const handleDeleteCard = async () => {
    if (!card) return;

    try {
      await deleteCard(card.id);
      setSnackbar({
        open: true,
        message: 'Card deleted successfully.',
        severity: 'success',
      });
      // Navigate back to the deck page after deletion
      navigate(`/decks/${deckId}`);
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



  if (loading) {
    return (
      <Container sx={{ mt: 4, textAlign: 'center' }}>
        <CircularProgress />
        <Typography variant="body1">Loading card details...</Typography>
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

  return (
    <Container sx={{ mt: 4 }}>
      {/* Render card text as Markdown */}
      <Box sx={{ mt: 2, p: 2, border: '1px solid #ddd', borderRadius: 2 }}>
        <Typography variant="h6" gutterBottom>
          {showFront ? 'Front' : 'Back'}
        </Typography>
        <ReactMarkdown>{showFront ? card.front.text : card.back.text}</ReactMarkdown>
      </Box>

      <Box sx={{ display: 'flex', gap: 2, mt: 4 }}>
        <Button variant="contained" color="primary" onClick={() => setShowFront(!showFront)}>
          Flip
        </Button>
        <Button variant="contained" color="success" onClick={() => handleCardAction('pass')}>
          Pass
        </Button>
        <Button variant="contained" color="warning" onClick={() => handleCardAction('pass')}>
          Skip
        </Button>
        <Button variant="contained" color="error" onClick={() => handleCardAction('pass')}>
          Fail
        </Button>
        
        <Button
          variant="outlined"
          color="secondary"
          component={RouterLink}
          to={`/card-form/${id}`}
        >
          Edit Card
        </Button>
        <Button
          variant="outlined"
          color="secondary"
          component={RouterLink}
          to={`/card-form/`}
        >
          Add Card
        </Button>

        {/* Delete Card Button */}
        <IconButton
          aria-label="delete"
          color="error"
          onClick={handleOpenDeleteDialog}
          sx={{ ml: 1 }}
        >
          <DeleteIcon />
        </IconButton>
      </Box>

    {/* Horizontal Status Bar */}
    <Box sx={{ mt: 4 }}>
      <HorizontalStatusBar pass={passCount} skip={skipCount} fail={failCount} />
    </Box>

      <Box sx={{ mt: 4 }}>
        <Typography variant="h6" gutterBottom>
          Link
        </Typography>
        {card.link ? (
          <Typography>
            <a href={card.link} target="_blank" rel="noopener noreferrer">
              {card.link}
            </a>
          </Typography>
        ) : (
          <Typography variant="body2" color="textSecondary">
            No link available.
          </Typography>
        )}
      </Box>


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
            Are you sure you want to delete this card? This action cannot be undone.
          </DialogContentText>
        </DialogContent>
        <DialogActions>
          <Button onClick={handleCloseDeleteDialog}>Cancel</Button>
          <Button onClick={handleDeleteCard} color="error" variant="contained">
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
        <AlertSnackbar onClose={handleCloseSnackbar} severity={snackbar.severity} sx={{ width: '100%' }}>
          {snackbar.message}
        </AlertSnackbar>
      </Snackbar>

    </Container>
  );
};

export default CardPage;
