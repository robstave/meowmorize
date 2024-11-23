// src/App.js
import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Navbar from './components/Navbar';
import Dashboard from './pages/Dashboard';
import ImportPage from './pages/ImportPage';
import DeckPage from './pages/DeckPage'; //

function App() {
  return (
    <Router>
      <Navbar />
      <Routes>
        <Route path="/" element={<Dashboard />} />
        <Route path="/import" element={<ImportPage />} />
        <Route path="/decks/:id" element={<DeckPage />} /> {/* **New Route for DeckPage** */}
        {/* Add more routes here as you create them */}
      </Routes>
    </Router>
  );
}

export default App;
