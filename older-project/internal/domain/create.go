package domain

import (
	"github.com/robstave/meowmorize/internal/domain/types"
)

func (s *Service) CreateDeck(deck types.Deck) error {

	for _, card := range deck.Cards {
		s.logger.Info("Imported Card",
			"uuid", card.ID,
			"did", card.DeckID,
			"front", card.Front.Text,
			"back", card.Back.Text)
	}

	err := s.deckRepo.CreateDeck(deck)
	if err != nil {
		s.logger.Error("Failed to create deck", "error", err)
		return err
	}

	s.logger.Info("Deck created successfully")
	return nil
}

// DeleteDeck deletes a deck by its ID
func (s *Service) DeleteDeck(deckID string) error {
	err := s.deckRepo.DeleteDeck(deckID)
	if err != nil {
		s.logger.Error("Failed to delete deck", "deck_id", deckID, "error", err)
		return err
	}

	s.logger.Info("Deck deleted successfully", "deck_id", deckID)
	return nil
}

// GetDeckByID retrieves a deck by its ID
func (s *Service) GetDeckByID(deckID string) (types.Deck, error) {
	deck, err := s.deckRepo.GetDeckByID(deckID)
	if err != nil {
		s.logger.Error("Failed to get deck by ID", "deck_id", deckID, "error", err)
		return types.Deck{}, err
	}
	return deck, nil
}
