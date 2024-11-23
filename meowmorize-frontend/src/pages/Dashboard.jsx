// src/pages/Dashboard.jsx
import React, { useEffect, useState } from 'react';
import { fetchDecks } from '../services/api';
import { Container, Typography, Table, TableBody, TableCell, TableContainer, TableHead, TableRow, Paper, CircularProgress, Alert } from '@mui/material';

const Dashboard = () => {
  const [decks, setDecks] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

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
            </TableRow>
          </TableHead>
          <TableBody>
            {decks.map((deck) => (
              <TableRow key={deck.id}>
                <TableCell component="th" scope="row">
                  {deck.name}
                </TableCell>
                <TableCell>{deck.description || 'No description provided.'}</TableCell>
                <TableCell align="right">{deck.cards.length}</TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </TableContainer>
    </Container>
  );
};

export default Dashboard;
