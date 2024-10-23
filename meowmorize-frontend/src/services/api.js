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

// You can add more API methods here as needed

export default api;
