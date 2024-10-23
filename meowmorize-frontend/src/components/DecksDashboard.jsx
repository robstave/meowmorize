// src/components/DecksDashboard.jsx
import React, { useEffect, useState } from 'react';
import { fetchDecks } from '../services/api';
import './DecksDashboard.css'; // Optional: For styling

const DecksDashboard = () => {
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
    return <div>Loading decks...</div>;
  }

  if (error) {
    return <div>{error}</div>;
  }

  if (decks.length === 0) {
    return <div>No decks available. Create a new deck to get started!</div>;
  }

  return (
    <div className="decks-dashboard">
      <h1>All Decks</h1>
      <table>
        <thead>
          <tr>
            <th>Name</th>
            <th>Description</th>
            <th>Number of Cards</th>
          </tr>
        </thead>
        <tbody>
          {decks.map((deck) => (
            <tr key={deck.id}>
              <td>{deck.name}</td>
              <td>{deck.description || 'No description provided.'}</td>
              <td>{deck.cards.length}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default DecksDashboard;
