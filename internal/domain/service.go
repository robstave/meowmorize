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

type RTOBLL interface {
	GetAllDecks() ([]types.Deck, error)
	CreateDefaultDeck() error
}

func NewService(logger *slog.Logger, deckRepo repositories.DeckRepository, cardRepo repositories.CardRepository) RTOBLL {
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

func (s *Service) CreateDefaultDeck() error {
	defaultDeck := types.Deck{
		Name: "Default Deck",
		Cards: []types.Card{
			{
				Front: "Sample Front",
				Back:  "Sample Back",
			},
		},
	}

	err := s.deckRepo.CreateDeck(defaultDeck)
	if err != nil {
		s.logger.Error("Failed to create default deck", "error", err)
		return err
	}

	s.logger.Info("Default deck created successfully")
	return nil
}
