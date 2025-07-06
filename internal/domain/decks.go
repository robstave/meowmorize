// domain/decks.go
package domain

import (
	"fmt"

	"github.com/robstave/meowmorize/internal/domain/types"
)

// ExportDeck retrieves the deck by ID for exporting
func (s *Service) ExportDeck(deckID string) (types.Deck, error) {
	deck, err := s.flashcardRepo.GetDeckByID(deckID)
	if err != nil {
		s.logger.Error("Failed to export deck", "deck_id", deckID, "error", err)
		return types.Deck{}, err
	}
	return deck, nil
}

// CollapseDecks merges all cards from the source deck into the target deck.
// It deletes each card from the source deck and adds it to the target deck.
// Parameters:
// - targetDeckID: ID of the deck to merge into.
// - sourceDeckID: ID of the deck to merge from.
func (s *Service) CollapseDecks(targetDeckID string, sourceDeckID string) error {
	if targetDeckID == "" || sourceDeckID == "" {
		return fmt.Errorf("both targetDeckID and sourceDeckID must be provided")
	}

	// Retrieve all cards associated with the source deck
	cards, err := s.flashcardRepo.GetCardsByDeckID(sourceDeckID)
	if err != nil {
		s.logger.Error("Failed to retrieve cards from source deck", "sourceDeckID", sourceDeckID, "error", err)
		return err
	}

	// For each card, remove the association with the source deck and add it to the target deck
	for _, card := range cards {
		// Remove association from the source deck
		if err := s.flashcardRepo.RemoveCardAssociation(sourceDeckID, card.ID); err != nil {
			s.logger.Error("Failed to remove association from source deck", "cardID", card.ID, "error", err)
			return err
		}

		// Add association to the target deck
		if err := s.flashcardRepo.AddCardAssociation(targetDeckID, card.ID); err != nil {
			s.logger.Error("Failed to add association to target deck", "cardID", card.ID, "error", err)
			return err
		}

		s.logger.Info("Moved card association", "cardID", card.ID, "from", sourceDeckID, "to", targetDeckID)
	}

	// Delete the source deck (which now has no associated cards)
	if err := s.flashcardRepo.DeleteDeck(sourceDeckID); err != nil {
		s.logger.Error("Failed to delete source deck", "sourceDeckID", sourceDeckID, "error", err)
		return err
	}

	return nil
}
