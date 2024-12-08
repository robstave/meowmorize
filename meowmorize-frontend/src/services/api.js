// src/services/api.js
import axios from 'axios';

//const API_BASE_URL = 'http://192.168.86.176:8789/api'; // Adjust if different
// Use the environment variable for the API base URL
const API_BASE_URL = process.env.REACT_APP_API_BASE_URL || 'http://192.168.86.176:8789/api';


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

export const updateCardStats = async (cardId, action, deckId) => {
  try {
    const response = await api.post('/cards/stats', { card_id: cardId, action, deck_id: deckId, });
    return response.data;
  } catch (error) {
    throw error;
  }
};



// Start a new session
export const startSession = async (deckId, count = -1, method = 'Random') => {
  try {
    const response = await api.post('/sessions/start', {
      deck_id: deckId,
      count,
      method,
    });
    return response.data;
  } catch (error) {
    throw error;
  }
};

// Get the next card in the session
export const getNextCard = async (deckId) => {
  try {
    const response = await api.get('/sessions/next', {
      params: { deck_id: deckId },
    });
    return response.data.card_id;
  } catch (error) {
    throw error;
  }
};

// Get session statistics
export const getSessionStats = async (deckId) => {
  try {
    const response = await api.get('/sessions/stats', {
      params: { deck_id: deckId },
    });
    return response.data;
  } catch (error) {
    throw error;
  }
};


// Update an existing deck
export const updateDeck = async (deckId, deckData) => {
  try {
    const response = await api.put(`/decks/${deckId}`, deckData);
    return response.data;
  } catch (error) {
    throw error;
  }
};


// Create an empty deck
export const createEmptyDeck = async () => {
  try {
    const response = await api.post('/decks/default', { defaultData: false });
    return response.data;
  } catch (error) {
    throw error;
  }
};




export default api;
