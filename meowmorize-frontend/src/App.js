// src/App.js
import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import DecksDashboard from './components/DecksDashboard';
// Import other components as you create them

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<DecksDashboard />} />
        {/* Add more routes here */}
      </Routes>
    </Router>
  );
}

export default App;
