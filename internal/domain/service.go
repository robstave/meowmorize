package domain

import (
	"log/slog"
	"sync"

	"github.com/robstave/meowmorize/internal/adapters/repositories"
	"github.com/robstave/meowmorize/internal/domain/types"
)

type Service struct {
	logger         *slog.Logger
	flashcardRepo  repositories.FlashcardsRepository
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
	CreateCard(card types.Card, deckID string, userID string) (*types.Card, error)
	UpdateCard(card types.Card) error
	DeleteCardByID(cardID string) error
	CloneCardToDeck(cardID string, targetDeckID string) (*types.Card, error)
	UpdateCardStats(cardID string, action types.CardAction, value *int, deckID string, userID string) error

	// LLM methods
	GetExplanation(prompt string) (string, error)
	IsLLMAvailable() bool

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
	GetAllUsers() ([]types.User, error)
	DeleteUser(userID string) error
	UpdateUserPassword(userID string, password string) error
	SeedUser() error

	GetSessionLogsBySessionID(sessionID string) ([]types.SessionLog, error)
	GetSessionLogIdsByUser(userID, deckID string) ([]string, error)
	GetSessionOverview(userID string, deckID string) ([]types.SessionOverview, error)
}

func NewService(logger *slog.Logger,
	flashcardRepo repositories.FlashcardsRepository,
	sessionLogRepo repositories.SessionLogRepository,
	llmRepo repositories.LLMRepository) MeowDomain {

	service := &Service{
		logger:         logger,
		flashcardRepo:  flashcardRepo,
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

// backfillCardOwners sets the user_id on any cards within the deck that do not have an owner
func (s *Service) backfillCardOwners(deck types.Deck, userID string) error {
	for _, card := range deck.Cards {
		if card.UserID == "" {
			card.UserID = userID
			if err := s.flashcardRepo.UpdateCard(card); err != nil {
				return err
			}
		}
	}
	return nil
}
