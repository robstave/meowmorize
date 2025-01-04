// src/components/DecksDashboard.jsx
import React, { useEffect, useState } from 'react';
import { fetchDecks } from '../services/api';
import './DecksDashboard.css'; // Optional: For styling
 
const DecksDashboard = () => {
  // State variables
  const [decks, setDecks] = useState([]);
  const [loading, setLoading] = useState(true); // To handle loading state
  const [error, setError] = useState(null); // To handle errors

  // useEffect to fetch decks on component mount
  useEffect(() => {
    const getDecks = async () => {
      try {
        const data = await fetchDecks(); // Fetch decks from API
        setDecks(data); // Update state with fetched decks
      } catch (err) {
        setError('Failed to fetch decks. Please try again later.'); // Set error message
        console.error(err);
      } finally {
        setLoading(false); // Set loading to false after fetching
      }
    };

    getDecks();
  }, []);

  // Display loading message while fetching
  if (loading) {
    return <div>Loading decks...</div>;
  }

  // Display error message if there's an error
  if (error) {
    return <div>{error}</div>;
  }

  

  // Display message if no decks are available
  if (decks.length === 0) {
    return <div>No decks available. Create a new deck to get started!</div>;
  }

  // Render the decks in a table
  return (
    <div className="decks-dashboard">
      <h1>All Decks</h1>
      <table>
        <thead>
          <tr>
            <th>Name</th>
   
            <th>Cards</th>
          </tr>
        </thead>
        <tbody>
          {decks.map((deck) => (
            <tr key={deck.id}>
              <td>{deck.name}</td>
               <td>{deck.cards.length}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default DecksDashboard;
