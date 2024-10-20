package domain

import (
	"log/slog"

	"github.com/robstave/meowmorize/internal/adapters/repositories"
	"github.com/robstave/meowmorize/internal/domain/types"
)

type Service struct {
	logger   *slog.Logger
	deckRepo repositories.DeckRepository
	cardRepo repositories.CardRepository
}

type BLL interface {
	GetAllDecks() ([]types.Deck, error)
	CreateDefaultDeck() error
	CreateDeck(types.Deck) error
}

func NewService(logger *slog.Logger, deckRepo repositories.DeckRepository, cardRepo repositories.CardRepository) BLL {
	return &Service{
		logger:   logger,
		deckRepo: deckRepo,
		cardRepo: cardRepo,
	}
}

func (s *Service) GetAllDecks() ([]types.Deck, error) {
	decks, err := s.deckRepo.GetAllDecks()
	if err != nil {
		s.logger.Error("Failed to retrieve decks", "error", err)
		return nil, err
	}
	return decks, nil
}
