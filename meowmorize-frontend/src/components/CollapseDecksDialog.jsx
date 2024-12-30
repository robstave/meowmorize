// src/components/CollapseDecksDialog.jsx

import React, { useState, useContext } from 'react';
import {
  Dialog,
  DialogTitle,
  DialogContent,
  DialogActions,
  Button,
  FormControl,
  InputLabel,
  Select,
  MenuItem,
  CircularProgress,
  Alert,
} from '@mui/material';
import { collapseDecks } from '../services/api'; // We'll create this API function
import { DeckContext } from '../context/DeckContext'; // Import DeckContext

const CollapseDecksDialog = ({ open, handleClose, decks, onSuccess, onError }) => {
  const [sourceDeck, setSourceDeck] = useState('');
  const [targetDeck, setTargetDeck] = useState('');
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState('');

  const { loadDecks } = useContext(DeckContext); // Access loadDecks from DeckContext

  
  const handleCollapse = async () => {
    // Validation
    if (!sourceDeck || !targetDeck) {
      setError('Both Source and Target decks must be selected.');
      return;
    }

    if (sourceDeck === targetDeck) {
      setError('Source and Target decks must be different.');
      return;
    }

    setLoading(true);
    setError('');

    try {
      await collapseDecks(targetDeck, sourceDeck);
      onSuccess(`Successfully collapsed "${sourceDeck}" into "${targetDeck}".`);
      // Reset selections
      setSourceDeck('');
      setTargetDeck('');
      handleClose();
    } catch (err) {
      console.error(err);
      onError('Failed to collapse decks. Please try again.');
      setError('Failed to collapse decks. Please try again.');
    } finally {
      setLoading(false);
    }
  };

  const handleCancel = () => {
    // Reset state on cancel
    setSourceDeck('');
    setTargetDeck('');
    setError('');
    handleClose();
  };

  return (
    <Dialog open={open} onClose={handleCancel} maxWidth="sm" fullWidth>
      <DialogTitle>Collapse Decks</DialogTitle>
      <DialogContent>
        {error && <Alert severity="error" sx={{ mb: 2 }}>{error}</Alert>}
        <FormControl fullWidth sx={{ mt: 2 }}>
          <InputLabel id="source-deck-label">Source Deck</InputLabel>
          <Select
            labelId="source-deck-label"
            id="source-deck"
            value={sourceDeck}
            label="Source Deck"
            onChange={(e) => setSourceDeck(e.target.value)}
          >
            {decks.map((deck) => (
              <MenuItem key={deck.id} value={deck.id}>
                {deck.name}
              </MenuItem>
            ))}
          </Select>
        </FormControl>

        <FormControl fullWidth sx={{ mt: 2 }}>
          <InputLabel id="target-deck-label">Target Deck</InputLabel>
          <Select
            labelId="target-deck-label"
            id="target-deck"
            value={targetDeck}
            label="Target Deck"
            onChange={(e) => setTargetDeck(e.target.value)}
          >
            {decks.map((deck) => (
              <MenuItem key={deck.id} value={deck.id}>
                {deck.name}
              </MenuItem>
            ))}
          </Select>
        </FormControl>
      </DialogContent>
      <DialogActions>
        <Button onClick={handleCancel} disabled={loading}>
          Cancel
        </Button>
        <Button onClick={handleCollapse} variant="contained" color="primary" disabled={loading}>
          {loading ? <CircularProgress size={24} /> : 'Collapse'}
        </Button>
      </DialogActions>
    </Dialog>
  );
};

export default CollapseDecksDialog;
