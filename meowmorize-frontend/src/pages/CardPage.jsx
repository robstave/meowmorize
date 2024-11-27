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
  List,
  ListItem,
} from '@mui/material';

const CardPage = () => {
  const { id } = useParams(); // Get card ID from URL
  const navigate = useNavigate();
  const [card, setCard] = useState(null);
  const [deckId, setDeckId] = useState(null); // Track the deckId
  const [showFront, setShowFront] = useState(true);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  // Fetch the card on component mount or when the ID changes
  useEffect(() => {
    const getCard = async () => {
      try {
        console.log('getting card');
        const data = await fetchCardById(id);
        console.log(card);
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

  // Function to select a random card from the current deck
  const selectCard = async () => {
    console.log('select');
    if (!deckId) {
        console.log('Deck ID is not available');
      setError('Deck ID is not available.');
      return;
    }
    console.log(deckId);

    try {
      const deck = await fetchDeckById(deckId);
      const randomCard = deck.cards[Math.floor(Math.random() * deck.cards.length)];
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
      <Typography variant="h4" gutterBottom>
        {showFront ? card.front.text : card.back.text}
      </Typography>

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
          Links
        </Typography>
        <List>
          {card.links && card.links.length > 0 ? (
            card.links.map((link, index) => (
              <ListItem key={index}>
                <a href={link} target="_blank" rel="noopener noreferrer">
                  {link}
                </a>
              </ListItem>
            ))
          ) : (
            <Typography variant="body2" color="textSecondary">
              No links available.
            </Typography>
          )}
        </List>
      </Box>
    </Container>
  );
};

export default CardPage;
