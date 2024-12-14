// services/card_test.go
package domain

import (
	"testing"

	"github.com/robstave/meowmorize/internal/adapters/repositories/mocks"
	"github.com/robstave/meowmorize/internal/domain/types"
	"github.com/robstave/meowmorize/internal/logger"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCardService_GetCardDetails_Success(t *testing.T) {
	cardRepo := new(mocks.CardRepository)
	userRepo := new(mocks.UserRepository)
	dr := new(mocks.DeckRepository)

	// Define the expected card, including the new Links field
	expectedCard := &types.Card{
		ID:     "card1",
		DeckID: "deck1",
		Front: types.CardFront{
			Text: "Front Text",
		},
		Back: types.CardBack{
			Text: "Back Text",
		},
		Link: "https://example.com/resource1",
	}

	//user := types.User{}

	// Setup expectations
	cardRepo.On("GetCardByID", "card1").Return(expectedCard, nil)
	userRepo.On("GetUserByUsername", "meow").Return(nil, nil)
	userRepo.On("CreateUser", mock.MatchedBy(func(u types.User) bool {
		return u.Username != "" // or any other basic validation
	})).Return(nil, nil)

	// Initialize the service with the mock repositories
	dm := NewService(logger.InitializeLogger(), dr, cardRepo, userRepo)

	// Call the method
	card, err := dm.GetCardByID("card1")

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, card)
	assert.Equal(t, expectedCard.ID, card.ID)
	assert.Equal(t, expectedCard.DeckID, card.DeckID)
	assert.Equal(t, expectedCard.Front, card.Front)
	assert.Equal(t, expectedCard.Back, card.Back)
	assert.Equal(t, expectedCard.Link, card.Link) // Assert Links field

	// Assert that the expectations were met
	cardRepo.AssertExpectations(t)
}
