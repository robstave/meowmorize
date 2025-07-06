package domain

import (
	"errors"
	"testing"

	"github.com/robstave/meowmorize/internal/domain/types"
	"github.com/robstave/meowmorize/internal/logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestClearDeckStats_Success(t *testing.T) {
	flashRepo, sessionRepo := setupRepositories()
	llmRepo := setupLLMRepository()
	flashRepo.UserRepository.On("GetUserByUsername", "meow").Return(&types.User{ID: "dummy", Username: "meow"}, nil)

	deck := types.Deck{ID: "deck1", Cards: []types.Card{{ID: "card1", PassCount: 2, FailCount: 1, SkipCount: 1}}}
	flashRepo.DeckRepository.On("GetDeckByID", "deck1").Return(deck, nil)
	flashRepo.DeckRepository.On("UpdateDeck", mock.MatchedBy(func(d types.Deck) bool {
		return d.ID == "deck1"
	})).Return(nil)
	flashRepo.CardRepository.On("UpdateCard", mock.MatchedBy(func(c types.Card) bool {
		return c.ID == "card1" && c.PassCount == 0 && c.FailCount == 0 && c.SkipCount == 0
	})).Return(nil)

	s := NewService(logger.InitializeLogger(), flashRepo, sessionRepo, llmRepo)

	err := s.StartSession("deck1", -1, types.RandomMethod, "meow")
	assert.NoError(t, err)

	err = s.ClearDeckStats("deck1", true, true)
	assert.NoError(t, err)
	sessStats, err := s.GetSessionStats("deck1")
	assert.NoError(t, err)

	assert.Equal(t, 0, sessStats.ViewedCount)

	//assert.Equal(t, 1, sess.Stats.Remaining)
	flashRepo.CardRepository.AssertExpectations(t)
	flashRepo.DeckRepository.AssertExpectations(t)
}

func TestClearDeckStats_GetDeckError(t *testing.T) {
	flashRepo, sessionRepo := setupRepositories()
	llmRepo := setupLLMRepository()
	flashRepo.UserRepository.On("GetUserByUsername", "meow").Return(&types.User{ID: "dummy", Username: "meow"}, nil)

	flashRepo.DeckRepository.On("GetDeckByID", "bad").Return(types.Deck{}, errors.New("not found"))

	s := NewService(logger.InitializeLogger(), flashRepo, sessionRepo, llmRepo)
	err := s.ClearDeckStats("bad", true, true)
	assert.Error(t, err)
	flashRepo.DeckRepository.AssertExpectations(t)
}
