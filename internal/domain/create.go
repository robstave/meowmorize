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
