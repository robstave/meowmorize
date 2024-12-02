// src/services/api.js
import axios from 'axios';

const API_BASE_URL = 'http://192.168.86.176:8789/api'; // Adjust if different

const api = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    'Content-Type': 'application/json',
  },
});

// Fetch all decks
export const fetchDecks = async () => {
  try {
    const response = await api.get('/decks');
    return response.data;
  } catch (error) {
    throw error;
  }
};

// Import a deck
export const importDeck = async (formData) => {
  try {
    const response = await api.post('/decks/import', formData, {
      headers: {
        'Content-Type': 'multipart/form-data',
      },
    });
    return response.data;
  } catch (error) {
    throw error;
  }
};

// Fetch a single deck by ID
export const fetchDeckById = async (deckId) => {
  try {
    const response = await api.get(`/decks/${deckId}`);
    return response.data;
  } catch (error) {
    throw error;
  }
};

export const fetchCardById = async (cardId) => {
  try {
    const response = await api.get(`/cards/${cardId}`);
    return response.data;
  } catch (error) {
    throw error;
  }
};



export const deleteDeck = async (deckId) => {
  try {
    const response = await api.delete(`/decks/${deckId}`);
    return response.data;
  } catch (error) {
    throw error;
  }
};


// Create a new card
export const createCard = async (cardData) => {
  try {
    const response = await api.post('/cards', cardData);
    return response.data;
  } catch (error) {
    throw error;
  }
};

// Update an existing card
export const updateCard = async (cardId, cardData) => {
  try {
    const response = await api.put(`/cards/${cardId}`, cardData);
    return response.data;
  } catch (error) {
    throw error;
  }
};

// Export a deck
export const exportDeck = async (deckId) => {
  try {
    const response = await api.get(`/decks/export/${deckId}`, {
      responseType: 'blob', // Important for handling binary data
    });
    return response.data;
  } catch (error) {
    throw error;
  }
};



// Delete a card by its ID
export const deleteCard = async (cardId) => {
  try {
    const response = await api.delete(`/cards/${cardId}`);
    return response.data;
  } catch (error) {
    throw error;
  }
};

export const updateCardStats = async (cardId, action) => {
  try {
    const response = await api.post('/cards/stats', { card_id: cardId, action });
    return response.data;
  } catch (error) {
    throw error;
  }
};




export default api;
