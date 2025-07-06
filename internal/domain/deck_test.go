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
	flashRepo, sessionRepo := setupRepositories()
	llmRepo := setupLLMRepository()

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
	flashRepo.DeckRepository.On("CreateDeck", testDeck).Return(nil)
	flashRepo.UserRepository.On("GetUserByUsername", "meow").Return(&types.User{ID: "dummy", Username: "meow"}, nil)

	s := NewService(logger.InitializeLogger(), flashRepo, sessionRepo, llmRepo)
	err := s.CreateDeck(testDeck)
	assert.NoError(t, err)
	flashRepo.DeckRepository.AssertExpectations(t)
}

func TestCreateDeck_Failure(t *testing.T) {
	flashRepo, sessionRepo := setupRepositories()
	llmRepo := setupLLMRepository()

	testDeck := types.Deck{
		ID:          "deck2",
		Name:        "Failing Deck",
		Description: "This deck creation should fail",
		UserID:      "user1",
	}

	expectedErr := errors.New("creation failed")
	flashRepo.DeckRepository.On("CreateDeck", testDeck).Return(expectedErr)
	flashRepo.UserRepository.On("GetUserByUsername", "meow").Return(&types.User{ID: "dummy", Username: "meow"}, nil)

	s := NewService(logger.InitializeLogger(), flashRepo, sessionRepo, llmRepo)
	err := s.CreateDeck(testDeck)
	assert.Error(t, err)
	assert.Equal(t, expectedErr, err)
	flashRepo.DeckRepository.AssertExpectations(t)
}

func TestDeleteDeck_Success(t *testing.T) {
	flashRepo, sessionRepo := setupRepositories()
	llmRepo := setupLLMRepository()

	deckID := "deck1"
	flashRepo.DeckRepository.On("DeleteDeck", deckID).Return(nil)
	flashRepo.UserRepository.On("GetUserByUsername", "meow").Return(&types.User{ID: "dummy", Username: "meow"}, nil)

	s := NewService(logger.InitializeLogger(), flashRepo, sessionRepo, llmRepo)
	err := s.DeleteDeck(deckID)
	assert.NoError(t, err)
	flashRepo.DeckRepository.AssertExpectations(t)
}

func TestDeleteDeck_Failure(t *testing.T) {
	flashRepo, sessionRepo := setupRepositories()
	llmRepo := setupLLMRepository()

	deckID := "deck2"
	expectedErr := errors.New("delete failed")
	flashRepo.DeckRepository.On("DeleteDeck", deckID).Return(expectedErr)
	flashRepo.UserRepository.On("GetUserByUsername", "meow").Return(&types.User{ID: "dummy", Username: "meow"}, nil)

	s := NewService(logger.InitializeLogger(), flashRepo, sessionRepo, llmRepo)
	err := s.DeleteDeck(deckID)
	assert.Error(t, err)
	assert.Equal(t, expectedErr, err)
	flashRepo.DeckRepository.AssertExpectations(t)
}

func TestGetDeckByID_Success(t *testing.T) {
	flashRepo, sessionRepo := setupRepositories()
	llmRepo := setupLLMRepository()

	expectedDeck := types.Deck{
		ID:          "deck1",
		Name:        "Test Deck",
		Description: "A test deck",
		UserID:      "user1",
	}

	flashRepo.DeckRepository.On("GetDeckByID", "deck1").Return(expectedDeck, nil)
	flashRepo.UserRepository.On("GetUserByUsername", "meow").Return(&types.User{ID: "dummy", Username: "meow"}, nil)

	s := NewService(logger.InitializeLogger(), flashRepo, sessionRepo, llmRepo)
	deck, err := s.GetDeckByID("deck1")
	assert.NoError(t, err)
	assert.Equal(t, expectedDeck, deck)
	flashRepo.DeckRepository.AssertExpectations(t)
}

func TestGetDeckByID_Failure(t *testing.T) {
	flashRepo, sessionRepo := setupRepositories()
	llmRepo := setupLLMRepository()

	expectedErr := errors.New("not found")
	flashRepo.DeckRepository.On("GetDeckByID", "nonexistent").Return(types.Deck{}, expectedErr)
	flashRepo.UserRepository.On("GetUserByUsername", "meow").Return(&types.User{ID: "dummy", Username: "meow"}, nil)

	s := NewService(logger.InitializeLogger(), flashRepo, sessionRepo, llmRepo)
	deck, err := s.GetDeckByID("nonexistent")
	assert.Error(t, err)
	assert.Equal(t, expectedErr, err)
	assert.Equal(t, types.Deck{}, deck)
	flashRepo.DeckRepository.AssertExpectations(t)
}

func TestGetAllDecks_Success(t *testing.T) {
	flashRepo, sessionRepo := setupRepositories()
	llmRepo := setupLLMRepository()

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

	flashRepo.DeckRepository.On("GetAllDecksByUser", "user1").Return(expectedDecks, nil)
	flashRepo.UserRepository.On("GetUserByUsername", "meow").Return(&types.User{ID: "dummy", Username: "meow"}, nil)

	s := NewService(logger.InitializeLogger(), flashRepo, sessionRepo, llmRepo)
	decks, err := s.GetAllDecks("user1")
	assert.NoError(t, err)
	assert.Equal(t, expectedDecks, decks)
	flashRepo.DeckRepository.AssertExpectations(t)
}

func TestUpdateDeck_Success(t *testing.T) {
	flashRepo, sessionRepo := setupRepositories()
	llmRepo := setupLLMRepository()

	testDeck := types.Deck{
		ID:          "deck1",
		Name:        "Updated Deck",
		Description: "Updated description",
		UserID:      "user1",
	}

	flashRepo.DeckRepository.On("UpdateDeck", testDeck).Return(nil)
	flashRepo.UserRepository.On("GetUserByUsername", "meow").Return(&types.User{ID: "dummy", Username: "meow"}, nil)

	s := NewService(logger.InitializeLogger(), flashRepo, sessionRepo, llmRepo)
	err := s.UpdateDeck(testDeck)
	assert.NoError(t, err)
	flashRepo.DeckRepository.AssertExpectations(t)
}

func TestUpdateDeck_Failure(t *testing.T) {
	flashRepo, sessionRepo := setupRepositories()
	llmRepo := setupLLMRepository()

	testDeck := types.Deck{
		ID:          "deck2",
		Name:        "Failing Deck",
		Description: "This update should fail",
		UserID:      "user1",
	}

	expectedErr := errors.New("update failed")
	flashRepo.DeckRepository.On("UpdateDeck", testDeck).Return(expectedErr)
	flashRepo.UserRepository.On("GetUserByUsername", "meow").Return(&types.User{ID: "dummy", Username: "meow"}, nil)

	s := NewService(logger.InitializeLogger(), flashRepo, sessionRepo, llmRepo)
	err := s.UpdateDeck(testDeck)
	assert.Error(t, err)
	assert.Equal(t, expectedErr, err)
	flashRepo.DeckRepository.AssertExpectations(t)
}
