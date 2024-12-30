import React, { createContext, useState, useEffect, useContext } from 'react';
import { fetchDecks } from '../services/api';
import { AuthContext } from './AuthContext';

export const DeckContext = createContext();

export const DeckProvider = ({ children }) => {
  const { auth } = useContext(AuthContext);
  const [decks, setDecks] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  const loadDecks = async () => {
    if (!auth.token) {
      setDecks([]);
      setLoading(false);
      return;
    }
    try {
      setLoading(true);
      const data = await fetchDecks();
      setDecks(data);
      setError(null);
    } catch (err) {
      console.error('Failed to fetch decks:', err);
      setError('Failed to fetch decks.');
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    loadDecks();
  }, [auth.token]);

  return (
    <DeckContext.Provider value={{ decks, setDecks, loadDecks, loading, error }}>
      {children}
    </DeckContext.Provider>
  );
};
