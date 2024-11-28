// domain/decks.go
package domain

import (
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
