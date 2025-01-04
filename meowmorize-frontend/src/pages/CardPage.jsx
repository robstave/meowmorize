import React, { useState, useEffect, useContext } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import { fetchCardById, deleteCard } from '../services/api';
import { updateCardStats, getSessionStats, getNextCard, fetchDeckById } from '../services/api';
import remarkGfm from 'remark-gfm';
import { Fade } from '@mui/material';
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
  Rating,

} from '@mui/material';
import Menu from '@mui/material/Menu';
import MenuItem from '@mui/material/MenuItem';
import MoreVertIcon from '@mui/icons-material/MoreVert';

import ReactMarkdown from 'react-markdown';

import { Link as RouterLink } from 'react-router-dom';

import DeleteIcon from '@mui/icons-material/Delete'; // Import DeleteIcon
import MuiAlert from '@mui/material/Alert'; // For Snackbar Alert

import PieStatusChart from '../components/CatStatusChart';
import CardStatsBar from '../components/CardStatsBar';
import { ThemeContext } from '../context/ThemeContext';



const AlertSnackbar = React.forwardRef(function Alert(props, ref) {
  return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
});


const CardPage = () => {
  const { id } = useParams(); // Get card ID from URL
  const navigate = useNavigate();
  const [card, setCard] = useState(null);
  const [deckId, setDeckId] = useState(null); // Track the deckId
  const [deckName, setDeckName] = useState(''); // New state for deck name
  const [showFront, setShowFront] = useState(true);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  // const [history, setHistory] = useState([]); // History of recently used indices

  const [passCount, setPassCount] = useState(0);
  const [skipCount, setSkipCount] = useState(0);
  const [failCount, setFailCount] = useState(0);

  const [stars, setStars] = useState(0); // New state for stars


  const [sessionStats, setSessionStats] = useState({
    total_cards: 0,
    viewed_count: 0,
    remaining: 0,
    current_index: 0,

  });




  // State for Delete Card Dialog
  const [openDeleteDialog, setOpenDeleteDialog] = useState(false);

  // State for Snackbar Notifications
  const [snackbar, setSnackbar] = useState({
    open: false,
    message: '',
    severity: 'success', // 'success' | 'error' | 'warning' | 'info'
  });

  // Access ThemeContext to get isDarkMode
  const { isDarkMode } = useContext(ThemeContext); // Ensure ThemeContext is correctly set up

  const [anchorEl, setAnchorEl] = useState(null);
  const openMenu = Boolean(anchorEl);

 

  const handleMenuClick = (event) => {
    setAnchorEl(event.currentTarget);
  };

  const handleMenuClose = () => {
    setAnchorEl(null);
  };

  // Fetch the card on component mount or when the ID changes
  useEffect(() => {
    const getCard = async () => {
      try {
        const data = await fetchCardById(id);
        setCard(data);
        setDeckId(data.deck_id); // Set the deckId for future fetches
        // set card stats
        setPassCount(data.pass_count);
        setSkipCount(data.skip_count);
        setFailCount(data.fail_count);
        setStars(data.star_rating || 0); // Initialize stars
        // set session stats
        const stats = await getSessionStats(data.deck_id);
        setSessionStats(stats);


        // Fetch deck details to get the deck name
        const deckData = await fetchDeckById(data.deck_id);
        setDeckName(deckData.name);


      } catch (err) {
        setError('Failed to fetch card details. Please try again later.');
        console.error(err);
      } finally {
        setLoading(false);
      }
    };

    getCard();
  }, [id]);



  // Function to handle card actions: pass, fail, skip
  const handleCardAction = async (action, value) => {
    if (!card) return;

    try {
      // Send the action to the backend
 

      if (action === 'SetStars' && value !== null) {
        setStars(value); // Update local stars state
        await updateCardStats(card.id, 'SetStars', deckId, value);
        return;
      }
      setLoading(true);

      switch (action) {
        case 'pass':
          await updateCardStats(card.id, 'IncrementPass', deckId);
          break;
        case 'skip':
          await updateCardStats(card.id, 'IncrementSkip', deckId);
          break;
        case 'fail':
          await updateCardStats(card.id, 'IncrementFail', deckId);
          break;
        default:
          // code block
       }


      // Fetch the next card
      const nextCardId = await getNextCard(deckId);

      if (nextCardId) {
       // setShowFront(true); // Reset to show the front of the new card

        navigate(`/card/${nextCardId}`);
      } else {
        setSnackbar({
          open: true,
          message: 'No more cards in the session.',
          severity: 'info',
        });
      }

      // Optionally, refresh session stats
      const stats = await getSessionStats(deckId);
      setSessionStats(stats);





    } catch (err) {
      console.error('Failed to process the card action:', err);
      setSnackbar({
        open: true,
        message: 'Failed to process the action. Please try again.',
        severity: 'error',
      });
    }  finally {
    // Hide loading spinner
    setLoading(false);
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

      {/* Deck Title */}
      <Typography variant="h5" gutterBottom>
        {deckName}
      </Typography>

      {/* Render card text as Markdown */}
      {/* Card Header with Front/Back Label and Pie Chart */}
      <Fade in={!loading}>
      <Box
        sx={{
          display: 'flex',
          justifyContent: 'space-between',
          alignItems: 'flex-start',
          mt: 2,
          p: 2,
          border: '1px solid #ddd',
          borderRadius: 2,
          flexWrap: 'wrap', // To handle smaller screens
        }}
      >
        {/* Front/Back Label and Content */}
        <Box sx={{ flex: '1 1 80%', minWidth: '250px' }}>
          <Typography variant="h6" gutterBottom>
            {showFront ? 'Front' : 'Back'}
          </Typography>
          <ReactMarkdown remarkPlugins={[remarkGfm]}>{showFront ? card.front.text : card.back.text}</ReactMarkdown>
        </Box>

        {/* Pie Status Chart */}
        <Box sx={{ flex: '1 1 20%', minWidth: '80px', display: 'flex', justifyContent: 'center', alignItems: 'center', mt: { xs: 2, sm: 0 } }}>
          <PieStatusChart
            pass={passCount}
            skip={skipCount}
            fail={failCount}
            isDarkMode={isDarkMode} // Pass the dark mode flag
          />
        </Box>
      </Box>
      </Fade>


      {/* Star Rating */}
      <Box sx={{ display: 'flex', alignItems: 'center', mt: 2 }}>
        <Typography variant="h6" sx={{ mr: 2 }}>
          Your Rating:
        </Typography>
        <Rating
          name="card-rating"
          value={stars}
          onChange={(event, newValue) => {
            console.log("beep")
            if (newValue !== null) {
              handleCardAction('SetStars', newValue);
            }
          }}
        />
      </Box>



      <Box sx={{ display: 'flex', gap: 2, mt: 4 }}>
        <Button variant="contained" color="primary" onClick={() => setShowFront(!showFront)}>
          Flip
        </Button>
        <Button variant="contained" color="success" onClick={() => handleCardAction('pass')}>
          Pass
        </Button>
        <Button variant="contained" color="warning" onClick={() => handleCardAction('skip')}>
          Skip
        </Button>
        <Button variant="contained" color="error" onClick={() => handleCardAction('fail')}>
          Fail
        </Button>

      

        {/* Menu Button for Add, Edit, Delete */}
        <IconButton
          aria-label="more"
          aria-controls={openMenu ? 'action-menu' : undefined}
          aria-haspopup="true"
          aria-expanded={openMenu ? 'true' : undefined}
          onClick={handleMenuClick}
        >
          <MoreVertIcon />
        </IconButton>

        <Menu
          id="action-menu"
          anchorEl={anchorEl}
          open={openMenu}
          onClose={handleMenuClose}
          MenuListProps={{
            'aria-labelledby': 'action-button',
          }}
        >
          <MenuItem component={RouterLink} to={`/card-form/`} onClick={handleMenuClose}>
            Add Card
          </MenuItem>
          <MenuItem component={RouterLink} to={`/card-form/${id}`} onClick={handleMenuClose}>
            Edit Card
          </MenuItem>
          <MenuItem onClick={handleOpenDeleteDialog}>Delete Card</MenuItem>
        </Menu>

      </Box>

      {/* Horizontal Status Bar */}
      <Box sx={{ mt: 4 }}>
        <CardStatsBar cards={sessionStats.card_stats} />
      </Box>

      {/* Session Statistics at the Bottom */}
      <Box sx={{ mt: 4 }}>
        <Typography variant="h6" gutterBottom>
          Session Statistics
        </Typography>
        <Typography variant="body1">
          <strong>Total Cards:</strong> {sessionStats.total_cards}
        </Typography>
        <Typography variant="body1">
          <strong>Viewed Count:</strong> {sessionStats.viewed_count}
        </Typography>
        <Typography variant="body1">
          <strong>Remaining:</strong> {sessionStats.remaining}
        </Typography>
        <Typography variant="body1">
          <strong>Current Index:</strong> {sessionStats.current_index}
        </Typography>
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
