// internal/domain/service.go

package domain

import (
	"fmt"

	"github.com/robstave/meowmorize/internal/domain/types"
)

// ClearDeckStats clears the statistics for a given deck.
// Parameters:
// - deckID: The ID of the deck.
// - clearSession: If true, resets the session statistics for the deck.
// - clearStats: If true, resets the pass, fail, and skip counts for all cards in the deck.
func (s *Service) ClearDeckStats(deckID string, clearSession bool, clearStats bool) error {
	// Retrieve the deck to ensure it exists
	deck, err := s.flashcardRepo.GetDeckByID(deckID)
	if err != nil {
		s.logger.Error("Failed to retrieve deck  lear", "deck_id", deckID, "error", err)
		return err
	}

	// Check if the deck was found by verifying the ID
	if deck.ID == "" {
		s.logger.Warn("Deck not found", "deck_id", deckID)
		return fmt.Errorf("deck with ID %s not found", deckID)
	}

	// Clear session stats if requested
	if clearSession {
		s.sessionsMu.Lock()
		session, exists := s.sessions[deckID]
		if exists {
			// Reset session statistics
			s.logger.Info("Reset session statistics-----------", "deck_id", deckID)

			//session.mu.Lock()

			for i := range session.CardStats {
				session.CardStats[i].Viewed = false
				session.CardStats[i].Skipped = false
			}
			session.Index = 0
			session.Stats = types.SessionStats{
				TotalCards:   len(session.CardStats),
				ViewedCount:  0,
				Remaining:    len(session.CardStats),
				CurrentIndex: 0,
			}
			//session.mu.Unlock()
			s.logger.Info("Session stats reset", "deck_id", deckID)
		} else {
			s.logger.Info("No active session to reset", "deck_id", deckID)
		}
		s.sessionsMu.Unlock()
	}

	// Clear card statistics if requested
	if clearStats {
		s.logger.Info("clear card statistics---++--------", "deck_id", deckID)
		for _, card := range deck.Cards {
			card.PassCount = 0
			card.FailCount = 0
			card.SkipCount = 0

			if err := s.flashcardRepo.UpdateCard(card); err != nil {
				s.logger.Error("Failed to update card stats", "card_id", card.ID, "error", err)
				return err
			}
			s.logger.Info("Card stats reset", "card_id", card.ID)
		}
		s.logger.Info("All card statistics have been reset for deck", "deck_id", deckID)
	}

	return nil
}
