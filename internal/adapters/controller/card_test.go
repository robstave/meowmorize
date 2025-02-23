// internal/adapters/controller/card_test.go
package controller

import (
	"github.com/robstave/meowmorize/internal/domain/types"
)

// Helper function to get a sample card
func getCard() types.Card {
	return types.Card{
		ID: "e0c32c1c-b36f-4e10-9f47-b8e88c8ff383",
		Front: types.CardFront{
			Text: "Capital of France",
		},
		Back: types.CardBack{
			Text: "Paris",
		},
		Link: "https://example.com/paris",
	}
}
