// services/card_test.go
package domain

import (
	"testing"

	"github.com/robstave/meowmorize/internal/adapters/repositories/mocks"
	"github.com/robstave/meowmorize/internal/domain/types"
	"github.com/robstave/meowmorize/internal/logger"

	"github.com/stretchr/testify/assert"
)

func TestCardService_GetCardDetails_Success(t *testing.T) {
	cardRepo := new(mocks.CardRepository)
	dr := new(mocks.DeckRepository)
	// Initialize the service with the mock repository

	// Define the expected card
	expectedCard := &types.Card{
		ID:     "card1",
		DeckID: "deck1",
		Front: types.CardFront{
			Text: "Front Text",
		},
		Back: types.CardBack{
			Text: "Back Text",
		},
	}

	// Setup expectations
	cardRepo.On("GetCardByID", "card1").Return(expectedCard, nil)

	dm := NewService(logger.InitializeLogger(), dr, cardRepo)

	// Call the method
	card, err := dm.GetCardByID("card1")

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, card)
	assert.Equal(t, expectedCard, card)

	// Assert that the expectations were met
	cardRepo.AssertExpectations(t)
}

/*
func TestCardService_GetCardDetails_NotFound(t *testing.T) {
	// Create a mock CardRepository
	mockRepo := new(mocks.CardRepository)

	// Setup expectations
	mockRepo.On("GetCardByID", "non-existent").Return(nil, nil)

	// Initialize the service with the mock repository
	cardService := services.NewCardService(mockRepo)

	// Call the method
	card, err := cardService.GetCardDetails("non-existent")

	// Assertions
	assert.NoError(t, err)
	assert.Nil(t, card)

	// Assert that the expectations were met
	mockRepo.AssertExpectations(t)
}



*/
