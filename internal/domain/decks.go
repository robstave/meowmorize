// domain/decks.go
package domain

import (
	"fmt"

	"github.com/robstave/meowmorize/internal/domain/types"
)

// ExportDeck retrieves the deck by ID for exporting
func (s *Service) ExportDeck(deckID string) (types.Deck, error) {
	deck, err := s.deckRepo.GetDeckByID(deckID)
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

	// Retrieve all cards from the source deck
	cards, err := s.cardRepo.GetCardsByDeckID(sourceDeckID)
	if err != nil {
		s.logger.Error("Failed to retrieve cards from source deck", "sourceDeckID", sourceDeckID, "error", err)
		return err
	}

	// Iterate over each card and move it to the target deck
	for _, card := range cards {
		// Delete the card from the source deck
		if err := s.cardRepo.DeleteCardByID(card.ID); err != nil {
			s.logger.Error("Failed to delete card from source deck", "cardID", card.ID, "error", err)
			return err
		}

		// Create a new card in the target deck with the same data
		newCard := types.Card{
			DeckID:     targetDeckID,
			ID:         card.ID,
			Front:      card.Front,
			Back:       card.Back,
			Link:       card.Link,
			PassCount:  card.PassCount,
			FailCount:  card.FailCount,
			SkipCount:  card.SkipCount,
			StarRating: card.StarRating,
			Retired:    card.Retired,
			// CreatedAt and UpdatedAt will be handled by GORM
		}

		if err := s.cardRepo.CreateCard(newCard); err != nil {
			s.logger.Error("Failed to add card to target deck", "cardID", newCard.ID, "error", err)
			return err
		}

		s.logger.Info("---- Moved card from source to target deck", "cardID", newCard.ID, "sourceDeckID", sourceDeckID, "targetDeckID", targetDeckID)
	}

	if err := s.deckRepo.DeleteDeck(sourceDeckID); err != nil {
		s.logger.Error("Failed to delete source deck", "sourceDeckID", sourceDeckID, "error", err)
		return err
	}

	return nil

}
