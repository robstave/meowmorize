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

func TestUpdateCard(t *testing.T) {
	// Setup: Create a test card and save it to the database
	originalCard := types.Card{
		ID:     "test-id",
		DeckID: "deck-1",
		Front:  types.CardFront{Text: "Original Front"},
		Back:   types.CardBack{Text: "Original Back"},
	}

	cardRepo, _ := initializeCardRepository(t)

	err := cardRepo.CreateCard(originalCard)
	if err != nil {
		t.Fatalf("Failed to create card: %v", err)
	}

	// Update the card
	updatedCard := originalCard
	updatedCard.Front.Text = "Updated Front"
	err = cardRepo.UpdateCard(updatedCard)
	if err != nil {
		t.Fatalf("Failed to update card: %v", err)
	}

	// Retrieve the card and verify the update
	retrievedCard, err := cardRepo.GetCardByID("test-id")
	if err != nil {
		t.Fatalf("Failed to retrieve card: %v", err)
	}
	if retrievedCard.Front.Text != "Updated Front" {
		t.Errorf("Expected Front text to be 'Updated Front', got '%s'", retrievedCard.Front.Text)
	}
}

func TestDeleteCardByID(t *testing.T) {
	// Setup: Create a test card
	card := types.Card{
		ID:     "delete-id",
		DeckID: "deck-1",
		Front:  types.CardFront{Text: "Front"},
		Back:   types.CardBack{Text: "Back"},
	}
	repo, _ := initializeCardRepository(t)

	err := repo.CreateCard(card)
	if err != nil {
		t.Fatalf("Failed to create card: %v", err)
	}

	// Delete the card
	err = repo.DeleteCardByID("delete-id")
	if err != nil {
		t.Fatalf("Failed to delete card: %v", err)
	}

	// Verify deletion
	deletedCard, err := repo.GetCardByID("delete-id")
	if err != nil {
		t.Fatalf("Error retrieving card: %v", err)
	}
	if deletedCard != nil {
		t.Errorf("Expected card to be deleted, but found: %+v", deletedCard)
	}
}

/*
func TestCloneCardToDeck(t *testing.T) {
	// Setup: Create a source card
	sourceCard := types.Card{
		ID:     "source-id",
		DeckID: "deck-1",
		Front:  types.CardFront{Text: "Source Front"},
		Back:   types.CardBack{Text: "Source Back"},
	}
	repo, _ := initializeCardRepository(t)

	err := repo.CreateCard(sourceCard)
	if err != nil {
		t.Fatalf("Failed to create source card: %v", err)
	}

	// Clone the card to a new deck
	clonedCard, err := repo.CloneCardToDeck("source-id", "deck-2")
	if err != nil {
		t.Fatalf("Failed to clone card: %v", err)
	}

	// Verify the cloned card
	if clonedCard.ID == "" {
		t.Errorf("Expected cloned card to have a new ID")
	}
	if clonedCard.DeckID != "deck-2" {
		t.Errorf("Expected cloned card to have DeckID 'deck-2', got '%s'", clonedCard.DeckID)
	}
	if clonedCard.Front.Text != "Source Front" || clonedCard.Back.Text != "Source Back" {
		t.Errorf("Cloned card's content does not match the source card")
	}
}
*/
