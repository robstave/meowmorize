// internal/domain/session_test.go
package domain

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/robstave/meowmorize/internal/domain/types"
	"github.com/robstave/meowmorize/internal/logger"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestStartSession_Success(t *testing.T) {
	deckID := uuid.New().String()
	card1 := types.Card{
		ID:    "card1",
		Front: types.CardFront{Text: "Q1"},
		Back:  types.CardBack{Text: "A1"},
	}
	card2 := types.Card{
		ID:    "card2",
		Front: types.CardFront{Text: "Q2"},
		Back:  types.CardBack{Text: "A2"},
	}
	deck := types.Deck{
		ID:    deckID,
		Name:  "Test Deck",
		Cards: []types.Card{card1, card2},
	}

	flashRepo, sessionRepo := setupRepositories()
	llmRepo := setupLLMRepository()

	// Expectations for SeedUser and deck retrieval/update.
	flashRepo.UserRepository.On("GetUserByUsername", "meow").Return(&types.User{ID: "dummy", Username: "meow"}, nil)
	flashRepo.DeckRepository.On("GetDeckByID", deckID).Return(deck, nil)
	flashRepo.DeckRepository.On("UpdateDeck", mock.AnythingOfType("types.Deck")).Return(nil)

	s := NewService(logger.InitializeLogger(), flashRepo, sessionRepo, llmRepo)
	err := s.StartSession(deckID, -1, types.RandomMethod, "meow")
	assert.NoError(t, err)

	stats, err := s.GetSessionStats(deckID)
	assert.NoError(t, err)
	assert.Equal(t, 2, stats.TotalCards)
	assert.Equal(t, 0, stats.ViewedCount)

	flashRepo.DeckRepository.AssertExpectations(t)
	flashRepo.UserRepository.AssertExpectations(t)
}

func TestStartSession_Failure_GetDeck(t *testing.T) {
	deckID := uuid.New().String()
	flashRepo, sessionRepo := setupRepositories()
	llmRepo := setupLLMRepository()

	flashRepo.UserRepository.On("GetUserByUsername", "meow").Return(&types.User{ID: "dummy", Username: "meow"}, nil)

	flashRepo.DeckRepository.On("GetDeckByID", deckID).Return(types.Deck{}, errors.New("deck not found"))
	s := NewService(logger.InitializeLogger(), flashRepo, sessionRepo, llmRepo)
	err := s.StartSession(deckID, 1, types.RandomMethod, "meow")
	assert.Error(t, err)

	flashRepo.DeckRepository.AssertExpectations(t)
	flashRepo.UserRepository.AssertExpectations(t)
}

func TestStartSession_Failure_UpdateDeck(t *testing.T) {
	deckID := uuid.New().String()
	card1 := types.Card{
		ID:    "card1",
		Front: types.CardFront{Text: "Q1"},
		Back:  types.CardBack{Text: "A1"},
	}
	deck := types.Deck{
		ID:    deckID,
		Name:  "Test Deck",
		Cards: []types.Card{card1},
	}
	flashRepo, sessionRepo := setupRepositories()
	llmRepo := setupLLMRepository()

	flashRepo.UserRepository.On("GetUserByUsername", "meow").Return(&types.User{ID: "dummy", Username: "meow"}, nil)

	flashRepo.DeckRepository.On("GetDeckByID", deckID).Return(deck, nil)
	flashRepo.DeckRepository.On("UpdateDeck", mock.AnythingOfType("types.Deck")).Return(errors.New("update failed"))

	s := NewService(logger.InitializeLogger(), flashRepo, sessionRepo, llmRepo)
	err := s.StartSession(deckID, 1, types.RandomMethod, "meow")
	assert.Error(t, err)

	flashRepo.DeckRepository.AssertExpectations(t)
	flashRepo.UserRepository.AssertExpectations(t)
}

func TestAdjustSession_Success(t *testing.T) {
	deckID := uuid.New().String()
	card := types.Card{
		ID:         "card1",
		Front:      types.CardFront{Text: "Q1"},
		Back:       types.CardBack{Text: "A1"},
		StarRating: 3,
	}
	deck := types.Deck{
		ID:    deckID,
		Name:  "Test Deck",
		Cards: []types.Card{card},
	}

	flashRepo, sessionRepo := setupRepositories()
	llmRepo := setupLLMRepository()

	// Expect SeedUser call.
	flashRepo.UserRepository.On("GetUserByUsername", "meow").Return(&types.User{ID: "dummy", Username: "meow"}, nil)

	// Expect deck retrieval and update.
	flashRepo.DeckRepository.On("GetDeckByID", deckID).Return(deck, nil)
	flashRepo.DeckRepository.On("UpdateDeck", mock.AnythingOfType("types.Deck")).Return(nil)

	// Expect session log creation
	sessionRepo.On("CreateLog", mock.MatchedBy(func(log types.SessionLog) bool {
		return log.DeckID == deckID && log.CardID == "card1" && log.Action == string(types.IncrementPass)
	})).Return(nil)

	// Initialize service.
	s := NewService(logger.InitializeLogger(), flashRepo, sessionRepo, llmRepo)

	// Start session.
	err := s.StartSession(deckID, 1, types.RandomMethod, "meow")
	assert.NoError(t, err)

	// Adjust session using IncrementPass action.
	err = s.AdjustSession(deckID, "card1", types.IncrementPass, 0, "meow")
	assert.NoError(t, err)

	// Retrieve session stats and verify the card is marked as Viewed and Passed.
	stats, err := s.GetSessionStats(deckID)
	assert.NoError(t, err)
	var found bool
	for _, cs := range stats.CardStats {
		if cs.CardID == "card1" {
			found = true
			assert.True(t, cs.Viewed, "Card should be marked as viewed")
			assert.True(t, cs.Passed, "Card should be marked as passed")
			break
		}
	}
	assert.True(t, found, "Card stat not found in session")

	flashRepo.DeckRepository.AssertExpectations(t)
	flashRepo.UserRepository.AssertExpectations(t)
	// No expectations were set on flashRepo.CardRepository, so no assertion needed for it.
}

func TestAdjustSession_InvalidCard(t *testing.T) {
	deckID := uuid.New().String()
	card := types.Card{
		ID:    "card1",
		Front: types.CardFront{Text: "Q1"},
		Back:  types.CardBack{Text: "A1"},
	}
	deck := types.Deck{
		ID:    deckID,
		Name:  "Test Deck",
		Cards: []types.Card{card},
	}
	flashRepo, sessionRepo := setupRepositories()
	llmRepo := setupLLMRepository()

	flashRepo.UserRepository.On("GetUserByUsername", "meow").Return(&types.User{ID: "dummy", Username: "meow"}, nil)

	flashRepo.DeckRepository.On("GetDeckByID", deckID).Return(deck, nil)
	flashRepo.DeckRepository.On("UpdateDeck", mock.AnythingOfType("types.Deck")).Return(nil)

	s := NewService(logger.InitializeLogger(), flashRepo, sessionRepo, llmRepo)
	err := s.StartSession(deckID, 1, types.RandomMethod, "meow")
	assert.NoError(t, err)

	// Do not set up GetCardByID for a non-existent card.
	err = s.AdjustSession(deckID, "non-existent", types.IncrementPass, 0, "meow")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "card not found in session")

	flashRepo.DeckRepository.AssertExpectations(t)
	flashRepo.UserRepository.AssertExpectations(t)
	flashRepo.CardRepository.AssertExpectations(t)
}

func TestGetNextCard_Success(t *testing.T) {
	deckID := uuid.New().String()
	card1 := types.Card{
		ID:    "card1",
		Front: types.CardFront{Text: "Q1"},
		Back:  types.CardBack{Text: "A1"},
	}
	card2 := types.Card{
		ID:    "card2",
		Front: types.CardFront{Text: "Q2"},
		Back:  types.CardBack{Text: "A2"},
	}
	deck := types.Deck{
		ID:    deckID,
		Name:  "Test Deck",
		Cards: []types.Card{card1, card2},
	}
	flashRepo, sessionRepo := setupRepositories()
	llmRepo := setupLLMRepository()

	flashRepo.UserRepository.On("GetUserByUsername", "meow").Return(&types.User{ID: "dummy", Username: "meow"}, nil)
	flashRepo.DeckRepository.On("GetDeckByID", deckID).Return(deck, nil)
	flashRepo.DeckRepository.On("UpdateDeck", mock.AnythingOfType("types.Deck")).Return(nil)

	s := NewService(logger.InitializeLogger(), flashRepo, sessionRepo, llmRepo)
	err := s.StartSession(deckID, -1, types.RandomMethod, "meow")
	assert.NoError(t, err)

	nextCardID, err := s.GetNextCard(deckID)
	assert.NoError(t, err)
	assert.NotEmpty(t, nextCardID)
	// Check that the returned card ID is one of the deck's cards.
	assert.Contains(t, []string{"card1", "card2"}, nextCardID)

	flashRepo.DeckRepository.AssertExpectations(t)
	flashRepo.UserRepository.AssertExpectations(t)
}

func TestGetNextCard_NoSession(t *testing.T) {
	flashRepo, sessionRepo := setupRepositories()
	llmRepo := setupLLMRepository()

	flashRepo.UserRepository.On("GetUserByUsername", "meow").Return(&types.User{ID: "dummy", Username: "meow"}, nil)

	s := NewService(logger.InitializeLogger(), flashRepo, sessionRepo, llmRepo)
	nextCardID, err := s.GetNextCard("non-existent-deck")
	assert.Error(t, err)
	assert.Empty(t, nextCardID)

	flashRepo.UserRepository.AssertExpectations(t)
}

func TestClearSession_Success(t *testing.T) {
	deckID := uuid.New().String()
	card := types.Card{
		ID:    "card1",
		Front: types.CardFront{Text: "Q1"},
		Back:  types.CardBack{Text: "A1"},
	}
	deck := types.Deck{
		ID:    deckID,
		Name:  "Test Deck",
		Cards: []types.Card{card},
	}
	flashRepo, sessionRepo := setupRepositories()
	llmRepo := setupLLMRepository()

	flashRepo.UserRepository.On("GetUserByUsername", "meow").Return(&types.User{ID: "dummy", Username: "meow"}, nil)
	flashRepo.DeckRepository.On("GetDeckByID", deckID).Return(deck, nil)
	flashRepo.DeckRepository.On("UpdateDeck", mock.AnythingOfType("types.Deck")).Return(nil)

	s := NewService(logger.InitializeLogger(), flashRepo, sessionRepo, llmRepo)
	err := s.StartSession(deckID, 1, types.RandomMethod, "meow")
	assert.NoError(t, err)

	err = s.ClearSession(deckID)
	assert.NoError(t, err)

	// After clearing, GetSessionStats should return empty stats.
	stats, err := s.GetSessionStats(deckID)
	assert.NoError(t, err)
	assert.Equal(t, 0, stats.TotalCards)

	flashRepo.UserRepository.AssertExpectations(t)
	flashRepo.DeckRepository.AssertExpectations(t)
}

func TestGetSessionStats_NoSession(t *testing.T) {
	flashRepo, sessionRepo := setupRepositories()
	llmRepo := setupLLMRepository()

	flashRepo.UserRepository.On("GetUserByUsername", "meow").Return(&types.User{ID: "dummy", Username: "meow"}, nil)

	s := NewService(logger.InitializeLogger(), flashRepo, sessionRepo, llmRepo)
	stats, err := s.GetSessionStats("non-existent-deck")
	assert.NoError(t, err)
	assert.Equal(t, 0, stats.TotalCards)

	flashRepo.UserRepository.AssertExpectations(t)
}
