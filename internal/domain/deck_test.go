// services/deck_test.go
package domain

import (
	"errors"
	"testing"

	"github.com/robstave/meowmorize/internal/domain/types"
	"github.com/robstave/meowmorize/internal/logger"

	"github.com/stretchr/testify/assert"
)

func TestCreateDeck_Success(t *testing.T) {
	cardRepo, userRepo, deckRepo, sessionRepo := setupRepositories()

	testDeck := types.Deck{
		ID:          "deck1",
		Name:        "Test Deck",
		Description: "A test deck",
		UserID:      "user1",
		Cards: []types.Card{
			{
				ID: "card1",
				Front: types.CardFront{
					Text: "Front text",
				},
				Back: types.CardBack{
					Text: "Back text",
				},
			},
		},
	}

	// Expect the deck repository to be called for deck creation.
	deckRepo.On("CreateDeck", testDeck).Return(nil)
	userRepo.On("GetUserByUsername", "meow").Return(&types.User{ID: "dummy", Username: "meow"}, nil)

	s := NewService(logger.InitializeLogger(), deckRepo, cardRepo, userRepo, sessionRepo)
	err := s.CreateDeck(testDeck)
	assert.NoError(t, err)
	deckRepo.AssertExpectations(t)
}

func TestCreateDeck_Failure(t *testing.T) {
	cardRepo, userRepo, deckRepo, sessionRepo := setupRepositories()

	testDeck := types.Deck{
		ID:          "deck2",
		Name:        "Failing Deck",
		Description: "This deck creation should fail",
		UserID:      "user1",
	}

	expectedErr := errors.New("creation failed")
	deckRepo.On("CreateDeck", testDeck).Return(expectedErr)
	userRepo.On("GetUserByUsername", "meow").Return(&types.User{ID: "dummy", Username: "meow"}, nil)

	s := NewService(logger.InitializeLogger(), deckRepo, cardRepo, userRepo, sessionRepo)
	err := s.CreateDeck(testDeck)
	assert.Error(t, err)
	assert.Equal(t, expectedErr, err)
	deckRepo.AssertExpectations(t)
}

func TestDeleteDeck_Success(t *testing.T) {
	cardRepo, userRepo, deckRepo, sessionRepo := setupRepositories()

	deckID := "deck1"
	deckRepo.On("DeleteDeck", deckID).Return(nil)
	userRepo.On("GetUserByUsername", "meow").Return(&types.User{ID: "dummy", Username: "meow"}, nil)

	s := NewService(logger.InitializeLogger(), deckRepo, cardRepo, userRepo, sessionRepo)
	err := s.DeleteDeck(deckID)
	assert.NoError(t, err)
	deckRepo.AssertExpectations(t)
}

func TestDeleteDeck_Failure(t *testing.T) {
	cardRepo, userRepo, deckRepo, sessionRepo := setupRepositories()

	deckID := "deck2"
	expectedErr := errors.New("delete failed")
	deckRepo.On("DeleteDeck", deckID).Return(expectedErr)
	userRepo.On("GetUserByUsername", "meow").Return(&types.User{ID: "dummy", Username: "meow"}, nil)

	s := NewService(logger.InitializeLogger(), deckRepo, cardRepo, userRepo, sessionRepo)
	err := s.DeleteDeck(deckID)
	assert.Error(t, err)
	assert.Equal(t, expectedErr, err)
	deckRepo.AssertExpectations(t)
}

func TestGetDeckByID_Success(t *testing.T) {
	cardRepo, userRepo, deckRepo, sessionRepo := setupRepositories()

	expectedDeck := types.Deck{
		ID:          "deck1",
		Name:        "Test Deck",
		Description: "A test deck",
		UserID:      "user1",
	}

	deckRepo.On("GetDeckByID", "deck1").Return(expectedDeck, nil)
	userRepo.On("GetUserByUsername", "meow").Return(&types.User{ID: "dummy", Username: "meow"}, nil)

	s := NewService(logger.InitializeLogger(), deckRepo, cardRepo, userRepo, sessionRepo)
	deck, err := s.GetDeckByID("deck1")
	assert.NoError(t, err)
	assert.Equal(t, expectedDeck, deck)
	deckRepo.AssertExpectations(t)
}

func TestGetDeckByID_Failure(t *testing.T) {
	cardRepo, userRepo, deckRepo, sessionRepo := setupRepositories()

	expectedErr := errors.New("not found")
	deckRepo.On("GetDeckByID", "nonexistent").Return(types.Deck{}, expectedErr)
	userRepo.On("GetUserByUsername", "meow").Return(&types.User{ID: "dummy", Username: "meow"}, nil)

	s := NewService(logger.InitializeLogger(), deckRepo, cardRepo, userRepo, sessionRepo)
	deck, err := s.GetDeckByID("nonexistent")
	assert.Error(t, err)
	assert.Equal(t, expectedErr, err)
	assert.Equal(t, types.Deck{}, deck)
	deckRepo.AssertExpectations(t)
}

func TestGetAllDecks_Success(t *testing.T) {
	cardRepo, userRepo, deckRepo, sessionRepo := setupRepositories()

	expectedDecks := []types.Deck{
		{
			ID:          "deck1",
			Name:        "Deck One",
			Description: "First deck",
			UserID:      "user1",
		},
		{
			ID:          "deck2",
			Name:        "Deck Two",
			Description: "Second deck",
			UserID:      "user1",
		},
	}

	deckRepo.On("GetAllDecksByUser", "user1").Return(expectedDecks, nil)
	userRepo.On("GetUserByUsername", "meow").Return(&types.User{ID: "dummy", Username: "meow"}, nil)

	s := NewService(logger.InitializeLogger(), deckRepo, cardRepo, userRepo, sessionRepo)
	decks, err := s.GetAllDecks("user1")
	assert.NoError(t, err)
	assert.Equal(t, expectedDecks, decks)
	deckRepo.AssertExpectations(t)
}

func TestUpdateDeck_Success(t *testing.T) {
	cardRepo, userRepo, deckRepo, sessionRepo := setupRepositories()

	testDeck := types.Deck{
		ID:          "deck1",
		Name:        "Updated Deck",
		Description: "Updated description",
		UserID:      "user1",
	}

	deckRepo.On("UpdateDeck", testDeck).Return(nil)
	userRepo.On("GetUserByUsername", "meow").Return(&types.User{ID: "dummy", Username: "meow"}, nil)

	s := NewService(logger.InitializeLogger(), deckRepo, cardRepo, userRepo, sessionRepo)
	err := s.UpdateDeck(testDeck)
	assert.NoError(t, err)
	deckRepo.AssertExpectations(t)
}

func TestUpdateDeck_Failure(t *testing.T) {
	cardRepo, userRepo, deckRepo, sessionRepo := setupRepositories()

	testDeck := types.Deck{
		ID:          "deck2",
		Name:        "Failing Deck",
		Description: "This update should fail",
		UserID:      "user1",
	}

	expectedErr := errors.New("update failed")
	deckRepo.On("UpdateDeck", testDeck).Return(expectedErr)
	userRepo.On("GetUserByUsername", "meow").Return(&types.User{ID: "dummy", Username: "meow"}, nil)

	s := NewService(logger.InitializeLogger(), deckRepo, cardRepo, userRepo, sessionRepo)
	err := s.UpdateDeck(testDeck)
	assert.Error(t, err)
	assert.Equal(t, expectedErr, err)
	deckRepo.AssertExpectations(t)
}
