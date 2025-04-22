package domain

import (
	"log/slog"
	"sync"

	"github.com/robstave/meowmorize/internal/adapters/repositories"
	"github.com/robstave/meowmorize/internal/domain/types"
)

type Service struct {
	logger         *slog.Logger
	deckRepo       repositories.DeckRepository
	cardRepo       repositories.CardRepository
	userRepo       repositories.UserRepository
	sessionLogRepo repositories.SessionLogRepository
	llmRepo        repositories.LLMRepository
	sessions       map[string]*types.Session
	sessionsMu     sync.RWMutex
}

type MeowDomain interface {
	// Deck methods
	CreateDeck(deck types.Deck) error
	GetAllDecks(userID string) ([]types.Deck, error)
	GetDeckByID(deckID string) (types.Deck, error)
	UpdateDeck(deck types.Deck) error
	DeleteDeck(deckID string) error
	ExportDeck(deckID string) (types.Deck, error)
	CollapseDecks(targetDeckID string, sourceDeckID string) error
	CreateDefaultDeck(defaultData bool, userID string) (types.Deck, error)

	// Card methods
	GetCardByID(cardID string) (*types.Card, error)
	CreateCard(card types.Card, deckID string) (*types.Card, error)
	UpdateCard(card types.Card) error
	DeleteCardByID(cardID string) error
	CloneCardToDeck(cardID string, targetDeckID string) (*types.Card, error)
	UpdateCardStats(cardID string, action types.CardAction, value *int, deckID string, userID string) error

	// LLM methods
	GetExplanation(prompt string) (string, error)

	// Session Management
	StartSession(deckID string, count int, method types.SessionMethod, userID string) error
	AdjustSession(deckID string, cardID string, action types.CardAction, value int, userID string) error
	GetNextCard(deckID string) (string, error)
	ClearSession(deckID string) error
	GetSessionStats(deckID string) (types.SessionStats, error)

	// Clear Deck Statistics
	ClearDeckStats(deckID string, clearSession bool, clearStats bool) error

	// User-related methods
	GetUserByUsername(username string) (*types.User, error)
	CreateUser(user types.User) error
	SeedUser() error

	GetSessionLogsBySessionID(sessionID string) ([]types.SessionLog, error)
	GetSessionLogIdsByUser(userID, deckID string) ([]string, error)
	GetSessionOverview(userID string, deckID string) ([]types.SessionOverview, error)
}

func NewService(logger *slog.Logger,
	deckRepo repositories.DeckRepository,
	cardRepo repositories.CardRepository,
	userRepo repositories.UserRepository,
	sessionLogRepo repositories.SessionLogRepository,
	llmRepo repositories.LLMRepository) MeowDomain {

	service := &Service{
		logger:         logger,
		deckRepo:       deckRepo,
		cardRepo:       cardRepo,
		userRepo:       userRepo,
		sessionLogRepo: sessionLogRepo,
		llmRepo:        llmRepo,
		sessions:       make(map[string]*types.Session),
		sessionsMu:     sync.RWMutex{},
	}

	// Seed the initial user. This is called on every startup, but will only create the user if it doesn't already exist
	// To reset the app, just delete the database file (assuming you're using the default sqlite3 database)
	if err := service.SeedUser(); err != nil {
		logger.Error("Failed to seed initial user", "error", err)
	}

	return service
}
