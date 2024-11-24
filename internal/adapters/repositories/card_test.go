// repositories/card_test.go
package repositories

import (
	"testing"

	th "github.com/robstave/meowmorize/internal/adapters/repositories/repositories_test"
	"github.com/robstave/meowmorize/internal/domain/types"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func initializeCardRepository(t *testing.T) (CardRepository, *gorm.DB) {
	db := th.SetupTestDB(t)
	cardRepo := NewCardRepositorySQLite(db)
	return cardRepo, db
}

func TestCardRepositorySQLite_GetCardByID_Success(t *testing.T) {
	cardRepo, db := initializeCardRepository(t)
	_, existingCard := th.SeedTestData(t, db)

	card, err := cardRepo.GetCardByID(existingCard.ID)
	assert.NoError(t, err)
	assert.NotNil(t, card)
	assert.Equal(t, existingCard.ID, card.ID)
	assert.Equal(t, existingCard.DeckID, card.DeckID)
	assert.Equal(t, existingCard.Front.Text, card.Front.Text)
	assert.Equal(t, existingCard.Back.Text, card.Back.Text)
}

func TestCardRepositorySQLite_GetCardByID_NotFound(t *testing.T) {
	cardRepo, _ := initializeCardRepository(t)

	card, err := cardRepo.GetCardByID("non-existent-id")
	assert.NoError(t, err)
	assert.Nil(t, card)
}

func TestCardRepositorySQLite_GetCardsByDeckID_Success(t *testing.T) {
	cardRepo, db := initializeCardRepository(t)
	deck, card := th.SeedTestData(t, db)

	cards, err := cardRepo.GetCardsByDeckID(deck.ID)
	assert.NoError(t, err)
	assert.Len(t, cards, 1)
	assert.Equal(t, card.ID, cards[0].ID)
}

func TestCardRepositorySQLite_GetCardsByDeckID_NoCards(t *testing.T) {
	cardRepo, db := initializeCardRepository(t)
	// Create an empty deck
	deck := types.Deck{
		ID:   "empty-deck",
		Name: "Empty Deck",
	}
	err := db.Create(&deck).Error
	assert.NoError(t, err)

	cards, err := cardRepo.GetCardsByDeckID(deck.ID)
	assert.NoError(t, err)
	assert.Len(t, cards, 0)
}

func TestCardRepositorySQLite_CreateCard_Success(t *testing.T) {
	cardRepo, db := initializeCardRepository(t)
	deck, _ := th.SeedTestData(t, db)

	newCard := types.Card{
		ID:     "card2",
		DeckID: deck.ID,
		Front: types.CardFront{
			Text: "New Front Text",
		},
		Back: types.CardBack{
			Text: "New Back Text",
		},
	}

	err := cardRepo.CreateCard(newCard)
	assert.NoError(t, err)

	// Verify the card was created
	var card types.Card
	err = db.First(&card, "id = ?", newCard.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, newCard.ID, card.ID)
	assert.Equal(t, newCard.Front.Text, card.Front.Text)
	assert.Equal(t, newCard.Back.Text, card.Back.Text)
}

func TestCardRepositorySQLite_CreateCard_DuplicateID(t *testing.T) {
	cardRepo, db := initializeCardRepository(t)
	_, existingCard := th.SeedTestData(t, db)

	duplicateCard := types.Card{
		ID:     existingCard.ID, // Duplicate ID
		DeckID: existingCard.DeckID,
		Front: types.CardFront{
			Text: "Duplicate Front",
		},
		Back: types.CardBack{
			Text: "Duplicate Back",
		},
	}

	err := cardRepo.CreateCard(duplicateCard)
	assert.Error(t, err)
}
