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

	// Setup expectations
	cardRepo.On("GetCardByID", "card1").Return(expectedCard, nil)

	// Initialize the service with the mock repositories
	dm := NewService(logger.InitializeLogger(), dr, cardRepo)

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
