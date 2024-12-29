package domain

import (
	"log/slog"
	"sync"

	"github.com/robstave/meowmorize/internal/adapters/repositories"
	"github.com/robstave/meowmorize/internal/domain/types"
)

type Service struct {
	logger     *slog.Logger
	deckRepo   repositories.DeckRepository
	cardRepo   repositories.CardRepository
	userRepo   repositories.UserRepository
	sessions   map[string]*types.Session // Key: DeckID
	sessionsMu sync.RWMutex
}

type MeowDomain interface {
	GetAllDecks() ([]types.Deck, error)
	CreateDefaultDeck(defaultDeck bool) (types.Deck, error)
	CreateDeck(types.Deck) error
	GetDeckByID(deckID string) (types.Deck, error)
	DeleteDeck(deckID string) error
	UpdateDeck(deck types.Deck) error
	GetCardByID(cardID string) (*types.Card, error)
	UpdateCard(card types.Card) error
	CreateCard(card types.Card) (*types.Card, error)
	DeleteCardByID(cardID string) error
	CloneCardToDeck(cardID string, targetDeckID string) (*types.Card, error)
	ExportDeck(deckID string) (types.Deck, error)
	UpdateCardStats(cardID string, action types.CardAction, value *int, deckID string) error

	// Session Management
	StartSession(deckID string, count int, method types.SessionMethod) error
	AdjustSession(deckID string, cardID string, action types.CardAction) error
	GetNextCard(deckID string) (string, error)
	ClearSession(deckID string) error
	GetSessionStats(deckID string) (types.SessionStats, error)

	// Clear Deck Statistics
	ClearDeckStats(deckID string, clearSession bool, clearStats bool) error

	// User-related methods
	GetUserByUsername(username string) (*types.User, error)
	CreateUser(user types.User) error
	SeedUser() error
	CollapseDecks(targetDeckID string, sourceDeckID string) error
}

// internal/domain/service.go
// ...
func NewService(logger *slog.Logger, deckRepo repositories.DeckRepository, cardRepo repositories.CardRepository, userRepo repositories.UserRepository) MeowDomain {
	service := &Service{
		logger:     logger,
		deckRepo:   deckRepo,
		cardRepo:   cardRepo,
		userRepo:   userRepo,
		sessions:   make(map[string]*types.Session),
		sessionsMu: sync.RWMutex{},
	}

	// Seed the initial user
	if err := service.SeedUser(); err != nil {
		logger.Error("Failed to seed initial user", "error", err)
	}

	return service
}
