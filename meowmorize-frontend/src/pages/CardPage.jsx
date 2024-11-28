import React, { useState, useEffect } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import { fetchCardById, fetchDeckById } from '../services/api';
import {
  Container,
  Typography,
  Button,
  Box,
  CircularProgress,
  Alert,

} from '@mui/material';
import ReactMarkdown from 'react-markdown';
import seedrandom from 'seedrandom';

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

  // Fetch the card on component mount or when the ID changes
  useEffect(() => {
    const getCard = async () => {
      try {
        const data = await fetchCardById(id);
        setCard(data);
        setDeckId(data.deck_id); // Set the deckId for future fetches
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


  // Function to select a random card from the current deck
  const selectCard = async () => {

    if (!deckId) {
      console.log('Deck ID is not available');
      setError('Deck ID is not available.');
      return;
    }

    try {
      const deck = await fetchDeckById(deckId);
            const randomIndex = getUniqueRandomIndex(deck.cards.length);
      //const randomIndex = Math.floor(rng() * deck.cards.length);
      const randomCard = deck.cards[randomIndex];
      setShowFront(true); // Reset to show the front of the new card
      navigate(`/card/${randomCard.id}`);
    } catch (err) {
      setError('Failed to fetch a random card.');
      console.error(err);
    }
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
        <Typography variant="h5" gutterBottom>
          {showFront ? 'Front' : 'Back'}
        </Typography>
        <ReactMarkdown>{showFront ? card.front.text : card.back.text}</ReactMarkdown>
      </Box>

      <Box sx={{ display: 'flex', gap: 2, mt: 4 }}>
        <Button variant="contained" color="primary" onClick={() => setShowFront(!showFront)}>
          Flip
        </Button>
        <Button variant="contained" color="success" onClick={selectCard}>
          Pass
        </Button>
        <Button variant="contained" color="warning" onClick={selectCard}>
          Skip
        </Button>
        <Button variant="contained" color="error" onClick={selectCard}>
          Fail
        </Button>
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
    </Container>
  );
};

export default CardPage;
