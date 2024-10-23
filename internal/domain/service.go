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
	GetDeckByID(deckID string) (types.Deck, error)
	DeleteDeck(deckID string) error
	UpdateDeck(deck types.Deck) error // New method
}

func NewService(logger *slog.Logger, deckRepo repositories.DeckRepository, cardRepo repositories.CardRepository) BLL {
	return &Service{
		logger:   logger,
		deckRepo: deckRepo,
		cardRepo: cardRepo,
	}
}
