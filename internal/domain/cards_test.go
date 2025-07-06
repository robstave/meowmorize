// services/card_test.go
package domain

import (
	"testing"

	"github.com/robstave/meowmorize/internal/domain/types"
	"github.com/robstave/meowmorize/internal/logger"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCardService_GetCardDetails_Success(t *testing.T) {

	flashRepo, sessionRepo := setupRepositories()

	llmRepo := setupLLMRepository()

	expectedCard := &types.Card{
		ID: "card1",
		Front: types.CardFront{
			Text: "Front Text",
		},
		Back: types.CardBack{
			Text: "Back Text",
		},
		Link: "https://example.com/resource1",
	}

	flashRepo.CardRepository.On("GetCardByID", "card1").Return(expectedCard, nil)
	flashRepo.UserRepository.On("GetUserByUsername", "meow").Return(nil, nil)
	flashRepo.UserRepository.On("CreateUser", mock.MatchedBy(func(u types.User) bool {
		return u.Username != ""
	})).Return(nil, nil)

	dm := NewService(logger.InitializeLogger(), flashRepo, sessionRepo, llmRepo)

	card, err := dm.GetCardByID("card1")
	assert.NoError(t, err)
	assert.NotNil(t, card)
	assert.Equal(t, expectedCard.ID, card.ID)
	assert.Equal(t, expectedCard.Front, card.Front)
	assert.Equal(t, expectedCard.Back, card.Back)
	assert.Equal(t, expectedCard.Link, card.Link)

	flashRepo.CardRepository.AssertExpectations(t)
}

func TestCardService_GetCardDetails_NotFound(t *testing.T) {
	flashRepo, sessionRepo := setupRepositories()
	llmRepo := setupLLMRepository()

	// Set up expectation for the SeedUser call in NewService.
	// SeedUser calls GetUserByUsername with the default username "meow".
	flashRepo.UserRepository.On("GetUserByUsername", "meow").Return(&types.User{ID: "dummy", Username: "meow"}, nil)

	// Simulate not finding the card.
	flashRepo.CardRepository.On("GetCardByID", "non-existent").Return(nil, nil)

	dm := NewService(logger.InitializeLogger(), flashRepo, sessionRepo, llmRepo)
	card, err := dm.GetCardByID("non-existent")
	assert.Error(t, err)
	assert.Nil(t, card)
	flashRepo.CardRepository.AssertExpectations(t)
	flashRepo.UserRepository.AssertExpectations(t)
}

// Test creating a new card successfully.
func TestCardService_CreateCard_Success(t *testing.T) {
	flashRepo, sessionRepo := setupRepositories()
	llmRepo := setupLLMRepository()

	// Set up expectations for the SeedUser call.
	// SeedUser will call GetUserByUsername with the default username "meow".
	flashRepo.UserRepository.On("GetUserByUsername", "meow").Return(&types.User{ID: "dummy", Username: "meow"}, nil)

	newCard := types.Card{
		Front: types.CardFront{Text: "Test Front"},
		Back:  types.CardBack{Text: "Test Back"},
		Link:  "https://example.com/test",
	}

	// Expect card creation and association with deck.
	flashRepo.CardRepository.On("CreateCard", mock.AnythingOfType("types.Card")).Return(nil)
	flashRepo.DeckRepository.On("AddCardToDeck", "deck1", mock.AnythingOfType("types.Card")).Return(nil)

	dm := NewService(logger.InitializeLogger(), flashRepo, sessionRepo, llmRepo)
	createdCard, err := dm.CreateCard(newCard, "deck1", "user1")
	assert.NoError(t, err)
	assert.NotNil(t, createdCard)
	// Verify that a new ID was assigned.
	assert.NotEmpty(t, createdCard.ID)

	flashRepo.CardRepository.AssertExpectations(t)
	flashRepo.DeckRepository.AssertExpectations(t)
	flashRepo.UserRepository.AssertExpectations(t)
}

// Test updating an existing card.
func TestCardService_UpdateCard_Success(t *testing.T) {
	flashRepo, sessionRepo := setupRepositories()
	llmRepo := setupLLMRepository()

	// Set up expectation for the SeedUser call.
	// SeedUser will call GetUserByUsername with the default username "meow".
	flashRepo.UserRepository.On("GetUserByUsername", "meow").Return(&types.User{ID: "dummy", Username: "meow"}, nil)

	// Existing card to be updated.
	existingCard := &types.Card{
		ID:    "card123",
		Front: types.CardFront{Text: "Old Front"},
		Back:  types.CardBack{Text: "Old Back"},
		Link:  "https://old.example.com",
	}

	// Expect retrieval of the card and its subsequent update.
	flashRepo.CardRepository.On("GetCardByID", "card123").Return(existingCard, nil)
	flashRepo.CardRepository.On("UpdateCard", mock.AnythingOfType("types.Card")).Return(nil)

	dm := NewService(logger.InitializeLogger(), flashRepo, sessionRepo, llmRepo)

	updatedCard := types.Card{
		ID:    "card123",
		Front: types.CardFront{Text: "New Front"},
		Back:  types.CardBack{Text: "New Back"},
		Link:  "https://new.example.com",
	}
	err := dm.UpdateCard(updatedCard)
	assert.NoError(t, err)

	flashRepo.CardRepository.AssertExpectations(t)
	flashRepo.UserRepository.AssertExpectations(t)
}

func TestCardService_DeleteCardByID_Success(t *testing.T) {
	flashRepo, sessionRepo := setupRepositories()
	llmRepo := setupLLMRepository()

	// Set up expectation for the SeedUser call.
	// SeedUser calls GetUserByUsername with the default username "meow".
	flashRepo.UserRepository.On("GetUserByUsername", "meow").Return(&types.User{ID: "dummy", Username: "meow"}, nil)

	// Expect deletion of the card.
	flashRepo.CardRepository.On("DeleteCardByID", "cardToDelete").Return(nil)

	dm := NewService(logger.InitializeLogger(), flashRepo, sessionRepo, llmRepo)
	err := dm.DeleteCardByID("cardToDelete")
	assert.NoError(t, err)

	flashRepo.CardRepository.AssertExpectations(t)
	flashRepo.UserRepository.AssertExpectations(t)
}

// Test cloning a card to another deck.
func TestCardService_CloneCardToDeck_Success(t *testing.T) {
	flashRepo, sessionRepo := setupRepositories()
	llmRepo := setupLLMRepository()

	clonedCard := &types.Card{
		ID:    "clonedCard1",
		Front: types.CardFront{Text: "Cloned Front"},
		Back:  types.CardBack{Text: "Cloned Back"},
		Link:  "https://example.com/cloned",
	}

	flashRepo.CardRepository.On("CloneCardToDeck", "cardOriginal", "deckTarget").Return(clonedCard, nil)
	flashRepo.UserRepository.On("GetUserByUsername", "meow").Return(&types.User{ID: "dummy", Username: "meow"}, nil)

	dm := NewService(logger.InitializeLogger(), flashRepo, sessionRepo, llmRepo)
	result, err := dm.CloneCardToDeck("cardOriginal", "deckTarget")
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "clonedCard1", result.ID)
	flashRepo.CardRepository.AssertExpectations(t)
}

// Test updating card statistics.
func TestCardService_UpdateCardStats_Success(t *testing.T) {
	flashRepo, sessionRepo := setupRepositories()
	llmRepo := setupLLMRepository()

	card := &types.Card{
		ID:         "cardStats1",
		Front:      types.CardFront{Text: "Front"},
		Back:       types.CardBack{Text: "Back"},
		Link:       "https://example.com/stats",
		PassCount:  0,
		FailCount:  0,
		SkipCount:  0,
		StarRating: 0,
	}
	// Expect retrieval of the card.
	flashRepo.CardRepository.On("GetCardByID", "cardStats1").Return(card, nil)
	flashRepo.UserRepository.On("GetUserByUsername", "meow").Return(&types.User{ID: "dummy", Username: "meow"}, nil)

	// Expect the card to be updated with incremented pass count.
	flashRepo.CardRepository.On("UpdateCard", mock.MatchedBy(func(updatedCard types.Card) bool {
		return updatedCard.ID == "cardStats1" && updatedCard.PassCount == 1
	})).Return(nil)

	dm := NewService(logger.InitializeLogger(), flashRepo, sessionRepo, llmRepo)
	err := dm.UpdateCardStats("cardStats1", types.IncrementPass, nil, "deckDummy", "meow")
	assert.NoError(t, err)
	flashRepo.CardRepository.AssertExpectations(t)
}
