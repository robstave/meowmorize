// src/pages/CardForm.jsx
import React, { useState, useEffect } from 'react';
import { useNavigate, useParams } from 'react-router-dom';
import {
  Container,
  Typography,
  TextField,
  Button,
  Box,
  Alert,
  CircularProgress,
} from '@mui/material';
import { createCard, updateCard, fetchCardById } from '../services/api';
import { Link as RouterLink } from 'react-router-dom';
import { fetchDecks } from '../services/api'; // Import fetchDecks

const CardForm = () => {
  const { id } = useParams(); // If 'id' exists, it's edit mode
  const navigate = useNavigate();

  const isEditMode = Boolean(id);

  const [formData, setFormData] = useState({
    deck_id: '', // You might need to handle deck selection differently
    front: { text: '' },
    back: { text: '' },
    link: '',
  });

  const [loading, setLoading] = useState(isEditMode); // Only load if editing
  const [error, setError] = useState('');
  const [success, setSuccess] = useState('');
  const [decks, setDecks] = useState([]);


  // Fetch decks on component mount
useEffect(() => {
    const getDecks = async () => {
      try {
        const data = await fetchDecks();
        setDecks(data);
      } catch (err) {
        setError('Failed to fetch decks.');
        console.error(err);
      }
    };
  
    getDecks();
  }, []);


  useEffect(() => {
    const getDecks = async () => {
        try {
          const data = await fetchDecks();
          setDecks(data);
        } catch (err) {
          setError('Failed to fetch decks.');
          console.error(err);
        }
      };

    if (isEditMode) {
      // Fetch existing card data
      const getCard = async () => {
        try {
          const data = await fetchCardById(id);
          setFormData({
            deck_id: data.deck_id,
            front: { text: data.front.text },
            back: { text: data.back.text },
            link: data.link || '',
          });
        } catch (err) {
          setError('Failed to fetch card details.');
          console.error(err);
        } finally {
          setLoading(false);
        }
      };

      getCard();
    }
    getDecks();
  }, [id, isEditMode]);

  const handleChange = (e) => {
    const { name, value } = e.target;

    // Handle nested fields
    if (name === 'front' || name === 'back') {
      setFormData((prev) => ({
        ...prev,
        [name]: { text: value },
      }));
    } else {
      setFormData((prev) => ({
        ...prev,
        [name]: value,
      }));
    }
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    // Basic validation
    if (!formData.deck_id) {
      setError('Deck ID is required.');
      return;
    }
    if (!formData.front.text || !formData.back.text) {
      setError('Front and Back texts are required.');
      return;
    }

    try {
      setError('');
      setSuccess('');
      if (isEditMode) {
        // Update existing card
        await updateCard(id, formData);
        setSuccess('Card updated successfully!');
      } else {
        // Create new card
        const newCard = await createCard(formData);
        setSuccess('Card created successfully!');
        console.log('created')
        console.log(newCard)
        navigate(`/card/${newCard.id}`); // Navigate to the newly created card
        return;
      }

      // Navigate back to the card page after a short delay to show success message
      setTimeout(() => {
        navigate(`/card/${id}`);
      }, 1500);
    } catch (err) {
      setError('Failed to submit the card. Please try again.');
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

  return (
    <Container maxWidth="sm" sx={{ mt: 4 }}>
      <Typography variant="h4" gutterBottom>
        {isEditMode ? 'Edit Card' : 'Add New Card'}
      </Typography>

      {error && <Alert severity="error" sx={{ mb: 2 }}>{error}</Alert>}
      {success && <Alert severity="success" sx={{ mb: 2 }}>{success}</Alert>}

      <Box component="form" onSubmit={handleSubmit} sx={{ display: 'flex', flexDirection: 'column', gap: 2 }}>
        {/* Deck ID Input */}
{/* Deck Selection Dropdown */}
{!isEditMode && (
  <TextField
    select
    label="Select Deck"
    name="deck_id"
    value={formData.deck_id}
    onChange={handleChange}
    required
    SelectProps={{
      native: true,
    }}
    helperText="Choose the deck this card belongs to."
  >
    <option value="" disabled>
      -- Select Deck --
    </option>
    {decks.map((deck) => (
      <option key={deck.id} value={deck.id}>
        {deck.name}
      </option>
    ))}
  </TextField>
)}


        {/* Front Text Input */}
        <TextField
          label="Front"
          name="front"
          value={formData.front.text}
          onChange={handleChange}
          required
          multiline
          rows={4}
          helperText="Enter the front text of the card."
        />

        {/* Back Text Input */}
        <TextField
          label="Back"
          name="back"
          value={formData.back.text}
          onChange={handleChange}
          required
          multiline
          rows={4}
          helperText="Enter the back text of the card."
        />

        {/* Link Input */}
        <TextField
          label="Link (Optional)"
          name="link"
          value={formData.link}
          onChange={handleChange}
          type="text"
          helperText="Enter a relevant link (e.g., reference or resource)."
        />

        {/* Submit Button */}
        <Button variant="contained" color="primary" type="submit">
          {isEditMode ? 'Update Card' : 'Add Card'}
        </Button>
      </Box>
    </Container>
  );
};

export default CardForm;
