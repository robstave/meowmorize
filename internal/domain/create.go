package domain

import (
	"github.com/robstave/meowmorize/internal/domain/types"
)

func (s *Service) CreateDeck(deck types.Deck) error {

	for _, card := range deck.Cards {
		s.logger.Info("Imported Card",
			"uuid", card.ID,
			"front", card.Front.Text,
			"back", card.Back.Text)
	}

	err := s.flashcardRepo.CreateDeck(deck)
	if err != nil {
		s.logger.Error("Failed to create deck", "error", err)
		return err
	}

	s.logger.Info("Deck created successfully")
	return nil
}
func (s *Service) DeleteDeck(deckID string) error {

	s.logger.Error("Deleting deck domain", "deckID", deckID)

	err := s.flashcardRepo.DeleteDeck(deckID)
	if err != nil {
		s.logger.Error("Failed to delete deck", "deckID", deckID, "error", err)
		return err
	}
	s.logger.Info("Deck deleted successfully", "deckID", deckID)
	return nil
}

// GetDeckByID retrieves a deck by its ID
func (s *Service) GetDeckByID(deckID string) (types.Deck, error) {
	deck, err := s.flashcardRepo.GetDeckByID(deckID)
	if err != nil {
		s.logger.Error("Failed to get deck by ID", "deck_id", deckID, "error", err)
		return types.Deck{}, err
	}
	return deck, nil
}

func (s *Service) GetAllDecks(userID string) ([]types.Deck, error) {
	return s.flashcardRepo.GetAllDecksByUser(userID)
}

func (s *Service) UpdateDeck(deck types.Deck) error {
	err := s.flashcardRepo.UpdateDeck(deck)
	if err != nil {
		s.logger.Error("Failed to update deck", "error", err)
		return err
	}
	s.logger.Info("Deck updated successfully", "deck_id", deck.ID)
	return nil
}
