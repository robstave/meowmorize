package domain

import (
	"errors"
	"testing"

	"github.com/robstave/meowmorize/internal/domain/types"
	"github.com/robstave/meowmorize/internal/logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateDefaultDeck_WithData(t *testing.T) {
	cardRepo, userRepo, deckRepo, sessionRepo := setupRepositories()
	llmRepo := setupLLMRepository()

	userRepo.On("GetUserByUsername", "meow").Return(&types.User{ID: "dummy", Username: "meow"}, nil)

	var savedDeck types.Deck
	deckRepo.On("CreateDeck", mock.AnythingOfType("types.Deck")).Return(nil).Run(func(args mock.Arguments) {
		savedDeck = args.Get(0).(types.Deck)
	})

	s := NewService(logger.InitializeLogger(), deckRepo, cardRepo, userRepo, sessionRepo, llmRepo)
	deck, err := s.CreateDefaultDeck(true, "user1")
	assert.NoError(t, err)
	assert.Equal(t, "user1", deck.UserID)
	assert.Len(t, deck.Cards, 2)
	assert.Equal(t, deck.Name, "Default Deck")
	assert.Equal(t, savedDeck.UserID, "user1")
	assert.Len(t, savedDeck.Cards, 2)
	deckRepo.AssertExpectations(t)
}

func TestCreateDefaultDeck_NoData(t *testing.T) {
	cardRepo, userRepo, deckRepo, sessionRepo := setupRepositories()
	llmRepo := setupLLMRepository()
	userRepo.On("GetUserByUsername", "meow").Return(&types.User{ID: "dummy", Username: "meow"}, nil)

	var savedDeck types.Deck
	deckRepo.On("CreateDeck", mock.AnythingOfType("types.Deck")).Return(nil).Run(func(args mock.Arguments) {
		savedDeck = args.Get(0).(types.Deck)
	})

	s := NewService(logger.InitializeLogger(), deckRepo, cardRepo, userRepo, sessionRepo, llmRepo)
	deck, err := s.CreateDefaultDeck(false, "user2")
	assert.NoError(t, err)
	assert.Equal(t, "user2", deck.UserID)
	assert.Len(t, deck.Cards, 0)
	assert.Len(t, savedDeck.Cards, 0)
	deckRepo.AssertExpectations(t)
}

func TestCreateDefaultDeck_Error(t *testing.T) {
	cardRepo, userRepo, deckRepo, sessionRepo := setupRepositories()
	llmRepo := setupLLMRepository()
	userRepo.On("GetUserByUsername", "meow").Return(&types.User{ID: "dummy", Username: "meow"}, nil)

	deckRepo.On("CreateDeck", mock.AnythingOfType("types.Deck")).Return(errors.New("fail"))

	s := NewService(logger.InitializeLogger(), deckRepo, cardRepo, userRepo, sessionRepo, llmRepo)
	deck, err := s.CreateDefaultDeck(false, "user3")
	assert.Error(t, err)
	assert.Empty(t, deck.ID)
	deckRepo.AssertExpectations(t)
}
