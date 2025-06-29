import React, { useState, useEffect, useContext } from 'react';
import { useParams, useNavigate, Link as RouterLink } from 'react-router-dom';
// API imports
import { 
  fetchCardById, 
  deleteCard, 
  explainCard, 
  getLLMStatus,
  updateCardStats, 
  getSessionStats, 
  getNextCard, 
  fetchDeckById 
} from '../services/api';

// Material UI imports
import {
  Container, Typography, Button, Box, IconButton,
  Dialog, DialogTitle, DialogContent, DialogActions, DialogContentText,
  Snackbar, CircularProgress, Alert, Rating, Link,
  Menu, MenuItem, Fade, Backdrop
} from '@mui/material';
import MuiAlert from '@mui/material/Alert';
import MoreVertIcon from '@mui/icons-material/MoreVert';

// Markdown related imports
import ReactMarkdown from 'react-markdown';
import remarkGfm from 'remark-gfm';
import rehypeHighlight from 'rehype-highlight';
import rehypeSanitize from 'rehype-sanitize';

// Components
import PieStatusChart from '../components/CatStatusChart';
import CardStatsBar from '../components/CardStatsBar';
import { ThemeContext } from '../context/ThemeContext';

const AlertSnackbar = React.forwardRef(function Alert(props, ref) {
  return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
});

const CardPage = () => {
  const { deckId, id } = useParams();
  const navigate = useNavigate();
  const { isDarkMode } = useContext(ThemeContext);

  // Card data states
  const [card, setCard] = useState(null);
  const [deckName, setDeckName] = useState('');
  const [showFront, setShowFront] = useState(true);
  const [stars, setStars] = useState(0);
  
  // Stats states
  const [passCount, setPassCount] = useState(0);
  const [skipCount, setSkipCount] = useState(0);
  const [failCount, setFailCount] = useState(0);
  const [sessionStats, setSessionStats] = useState({
    total_cards: 0,
    viewed_count: 0,
    remaining: 0,
    current_index: 0,
  });

  // UI states
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [isLLMAvailable, setIsLLMAvailable] = useState(false);
  const [openDeleteDialog, setOpenDeleteDialog] = useState(false);
  const [anchorEl, setAnchorEl] = useState(null);
  const openMenu = Boolean(anchorEl);
  const [explanation, setExplanation] = useState('');
  const [explaining, setExplaining] = useState(false);
  const [snackbar, setSnackbar] = useState({
    open: false,
    message: '',
    severity: 'success',
  });

  // Menu handling functions
  const handleMenuClick = (event) => {
    setAnchorEl(event.currentTarget);
  };

  const handleMenuClose = () => {
    setAnchorEl(null);
  };

  // Fetch the card data
  useEffect(() => {
    const fetchCardData = async () => {
      try {
        setLoading(true);
        setExplanation('');
        
        // Fetch card data and LLM status in parallel
        const [cardData, llmStatus] = await Promise.all([
          fetchCardById(id),
          getLLMStatus()
        ]);
        
        // Update card state
        setCard(cardData);
        setIsLLMAvailable(llmStatus);
        setShowFront(true);
        
        // Update card stats
        setPassCount(cardData.pass_count);
        setSkipCount(cardData.skip_count);
        setFailCount(cardData.fail_count);
        setStars(cardData.star_rating || 0);

        // Fetch session stats and deck info
        const [sessionData, deckData] = await Promise.all([
          getSessionStats(deckId),
          fetchDeckById(deckId)
        ]);
        
        setSessionStats(sessionData);
        setDeckName(deckData.name);
      } catch (err) {
        console.error('Error fetching card data:', err);
        setError('Failed to fetch card details. Please try again later.');
      } finally {
        setLoading(false);
      }
    };

    fetchCardData();
  }, [id, deckId]);

  // Function to handle card actions: pass, fail, skip
  const handleCardAction = async (action, value) => {
    if (!card) return;

    try {
      // Handle star rating separately
      if (action === 'SetStars' && value !== null) {
        setStars(value);
        await updateCardStats(card.id, 'SetStars', deckId, value);
        return;
      }
      
      setLoading(true);

      // Handle card progression actions
      const actionMap = {
        'pass': 'IncrementPass',
        'skip': 'IncrementSkip',
        'fail': 'IncrementFail'
      };
      
      const apiAction = actionMap[action];
      if (apiAction) {
        await updateCardStats(card.id, apiAction, deckId);
      }

      // Fetch the next card and progress
      const nextCardId = await getNextCard(deckId);
      if (nextCardId) {
        navigate(`/decks/${deckId}/card/${nextCardId}`);
      } else {
        setSnackbar({
          open: true,
          message: 'No more cards in the session.',
          severity: 'info',
        });
      }

      // Update session stats
      const stats = await getSessionStats(deckId);
      setSessionStats(stats);
    } catch (err) {
      console.error('Failed to process the card action:', err);
      setSnackbar({
        open: true,
        message: 'Failed to process the action. Please try again.',
        severity: 'error',
      });
    } finally {
      setLoading(false);
    }
  };

  const handleExplain = async () => {
    if (!card || showFront) return;

    try {
      setExplaining(true);
      const result = await explainCard(card.id, 'explain the answer');
      setExplanation(result.explanation);
    } catch (err) {
      console.error('Failed to get explanation:', err);
      setSnackbar({
        open: true,
        message: 'Failed to get explanation. Please try again.',
        severity: 'error',
      });
    } finally {
      setExplaining(false);
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


  // Display error message if card fetch failed

  if (error) {
    return (
      <Container sx={{ mt: 4 }}>
        <Alert severity="error">{error}</Alert>
      </Container>
    );
  }

  return (
    <Container sx={{ mt: 4 }}>
      {/* Deck Title as a Link */}
      <Typography variant="h5" gutterBottom>
        <Link component={RouterLink} to={`/decks/${deckId}`} underline="hover">
          {deckName}
        </Link>
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
            <ReactMarkdown rehypePlugins={[rehypeHighlight, rehypeSanitize]} remarkPlugins={[remarkGfm]}>
              {showFront ? card.front.text : card.back.text}
            </ReactMarkdown>
          </Box>

          {/* Pie Status Chart */}
          <Box
            sx={{
              flex: '1 1 20%',
              minWidth: '80px',
              display: 'flex',
              justifyContent: 'center',
              alignItems: 'center',
              mt: { xs: 2, sm: 0 },
            }}
          >
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

        {/* Only show explain button and menu when viewing the back */}
        {!showFront && (
          <>
            {isLLMAvailable && (
              <Button 
                variant="outlined" 
                color="info" 
                onClick={handleExplain}
                disabled={explaining}
              >
                {explaining ? 'Explaining...' : 'Explain'}
              </Button>
            )}

            <IconButton
              aria-label="more"
              aria-controls={openMenu ? 'action-menu' : undefined}
              aria-haspopup="true"
              aria-expanded={openMenu ? 'true' : undefined}
              onClick={handleMenuClick}
            >
              <MoreVertIcon />
            </IconButton>
          </>
        )}
      </Box>

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

      {/* Show explanation if available */}
      {!showFront && explanation && (
        <Box sx={{ mt: 3, p: 2, bgcolor: 'background.paper', borderRadius: 1, boxShadow: 1 }}>
          <Typography variant="h6" gutterBottom>
            Explanation
          </Typography>
          <Typography variant="body1" component="div">
            <ReactMarkdown rehypePlugins={[rehypeHighlight, rehypeSanitize]} remarkPlugins={[remarkGfm]}>
              {explanation}
            </ReactMarkdown>
          </Typography>
        </Box>
      )}

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
          Reference Link
        </Typography>
        {card.link ? (
          <Link href={card.link} target="_blank" rel="noopener noreferrer" color="primary">
            {card.link}
          </Link>
        ) : (
          <Typography variant="body2" color="text.secondary">
            No reference link available
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

      {/* Loading overlay to smooth card transitions */}
      <Backdrop
        sx={{ color: '#fff', zIndex: (theme) => theme.zIndex.drawer + 1 }}
        open={loading}
      >
        <CircularProgress color="inherit" />
      </Backdrop>
    </Container>
  );
};

export default CardPage;
