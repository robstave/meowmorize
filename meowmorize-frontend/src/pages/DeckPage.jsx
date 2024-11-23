// src/pages/DeckPage.jsx
import React, { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import { fetchDeckById } from '../services/api';
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
  TableRow,
  Paper,
  Button,
  Collapse,
} from '@mui/material';
import { Link as RouterLink } from 'react-router-dom';

const DeckPage = () => {
  const { id } = useParams(); // Extract the deck ID from the URL
  const [deck, setDeck] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  // State to manage the visibility of the cards list
  const [showCards, setShowCards] = useState(false);

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

  // Handler to toggle the visibility of the cards list
  const handleToggleCards = () => {
    setShowCards((prev) => !prev);
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

      {/* Preview Button to Toggle Cards List */}
      <Button
        variant="contained"
        color="primary"
        onClick={handleToggleCards}
        sx={{ mt: 2 }}
      >
        {showCards ? 'Hide Preview' : 'Show Preview'}
      </Button>

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
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </TableContainer>
      </Collapse>

      {/* Back to Dashboard Button */}
      <Button
        component={RouterLink}
        to="/"
        variant="outlined"
        color="secondary"
        sx={{ mt: 2 }}
      >
        Back to Dashboard
      </Button>
    </Container>
  );
};

export default DeckPage;
