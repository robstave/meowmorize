// repositories/deck_test.go

package repositories

import (
	"testing"

	th "github.com/robstave/meowmorize/internal/adapters/repositories/repositories_test"
	"github.com/robstave/meowmorize/internal/domain/types"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func initializeDeckRepository(t *testing.T) (DeckRepository, *gorm.DB) {
	db := th.SetupTestDB(t)
	deckRepo := NewDeckRepositorySQLite(db)
	return deckRepo, db
}

func TestDeckRepositorySQLite_GetAllDecks_WithDecks(t *testing.T) {
	deckRepo, db := initializeDeckRepository(t)
	deck1 := types.Deck{ID: "deck1", Name: "Deck One"}
	deck2 := types.Deck{ID: "deck2", Name: "Deck Two"}

	err := db.Create(&deck1).Error
	assert.NoError(t, err)
	err = db.Create(&deck2).Error
	assert.NoError(t, err)

	decks, err := deckRepo.GetAllDecks()
	assert.NoError(t, err)
	assert.Len(t, decks, 2)
}

func TestDeckRepositorySQLite_GetAllDecks_NoDecks(t *testing.T) {
	deckRepo, _ := initializeDeckRepository(t)

	decks, err := deckRepo.GetAllDecks()
	assert.NoError(t, err)
	assert.Len(t, decks, 0)
}

func TestDeckRepositorySQLite_GetDeckByID_NotFound(t *testing.T) {
	deckRepo, _ := initializeDeckRepository(t)

	_, err := deckRepo.GetDeckByID("non-existent-deck")
	assert.Error(t, err)
	assert.True(t, gorm.ErrRecordNotFound == err || err.Error() == "record not found")
}
func TestDeckRepositorySQLite_CreateDeck_Success(t *testing.T) {
	deckRepo, db := initializeDeckRepository(t)

	newDeck := types.Deck{
		ID:   "deck2",
		Name: "New Deck",
	}

	err := deckRepo.CreateDeck(newDeck)
	assert.NoError(t, err)

	// Verify the deck was created
	var deck types.Deck
	err = db.First(&deck, "id = ?", newDeck.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, newDeck.Name, deck.Name)
}
func TestDeckRepositorySQLite_CreateDeck_DuplicateID(t *testing.T) {
	deckRepo, db := initializeDeckRepository(t)
	existingDeck := types.Deck{
		ID:   "deck1",
		Name: "Existing Deck",
	}
	err := db.Create(&existingDeck).Error
	assert.NoError(t, err)

	duplicateDeck := types.Deck{
		ID:   existingDeck.ID, // Duplicate ID
		Name: "Duplicate Deck",
	}

	err = deckRepo.CreateDeck(duplicateDeck)
	assert.Error(t, err)
}

func TestDeckRepositorySQLite_DeleteDeck_NotFound(t *testing.T) {
	deckRepo, _ := initializeDeckRepository(t)

	err := deckRepo.DeleteDeck("non-existent-deck")
	assert.Nil(t, err)
}

func TestDeckRepositorySQLite_UpdateDeck_Success(t *testing.T) {
	deckRepo, db := initializeDeckRepository(t)
	deck, _ := th.SeedTestData(t, db)

	// Update the deck's name
	deck.Name = "Updated Deck Name"
	err := deckRepo.UpdateDeck(deck)
	assert.NoError(t, err)

	// Verify the update
	var updatedDeck types.Deck
	err = db.First(&updatedDeck, "id = ?", deck.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, "Updated Deck Name", updatedDeck.Name)
}

func TestDeckRepositorySQLite_UpdateDeck_NotFound(t *testing.T) {
	deckRepo, _ := initializeDeckRepository(t)

	nonExistentDeck := types.Deck{
		ID:   "non-existent-deck",
		Name: "Should Not Exist",
	}

	d := deckRepo.UpdateDeck(nonExistentDeck)
	assert.Nil(t, d)
}
