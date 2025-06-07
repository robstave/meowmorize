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
	cardRepo, userRepo, deckRepo, sessionRepo := setupRepositories()
	llmRepo := setupLLMRepository()
	userRepo.On("GetUserByUsername", "meow").Return(&types.User{ID: "dummy", Username: "meow"}, nil)

	deck := types.Deck{ID: "deck1", Cards: []types.Card{{ID: "card1", PassCount: 2, FailCount: 1, SkipCount: 1}}}
	deckRepo.On("GetDeckByID", "deck1").Return(deck, nil)
	cardRepo.On("UpdateCard", mock.MatchedBy(func(c types.Card) bool {
		return c.ID == "card1" && c.PassCount == 0 && c.FailCount == 0 && c.SkipCount == 0
	})).Return(nil)

	s := NewService(logger.InitializeLogger(), deckRepo, cardRepo, userRepo, sessionRepo, llmRepo)
	s.sessions["deck1"] = &types.Session{
		DeckID:    "deck1",
		CardStats: []types.CardStats{{CardID: "card1", Viewed: true, Skipped: true}},
		Stats:     types.SessionStats{TotalCards: 1, ViewedCount: 1, Remaining: 0, CurrentIndex: 1},
	}

	err := s.ClearDeckStats("deck1", true, true)
	assert.NoError(t, err)
	sess := s.sessions["deck1"]
	assert.False(t, sess.CardStats[0].Viewed)
	assert.False(t, sess.CardStats[0].Skipped)
	assert.Equal(t, 0, sess.Stats.ViewedCount)
	assert.Equal(t, 1, sess.Stats.Remaining)
	cardRepo.AssertExpectations(t)
	deckRepo.AssertExpectations(t)
}

func TestClearDeckStats_GetDeckError(t *testing.T) {
	cardRepo, userRepo, deckRepo, sessionRepo := setupRepositories()
	llmRepo := setupLLMRepository()
	userRepo.On("GetUserByUsername", "meow").Return(&types.User{ID: "dummy", Username: "meow"}, nil)

	deckRepo.On("GetDeckByID", "bad").Return(types.Deck{}, errors.New("not found"))

	s := NewService(logger.InitializeLogger(), deckRepo, cardRepo, userRepo, sessionRepo, llmRepo)
	err := s.ClearDeckStats("bad", true, true)
	assert.Error(t, err)
	deckRepo.AssertExpectations(t)
}
