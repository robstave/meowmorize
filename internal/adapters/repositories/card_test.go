// repositories/card_test.go
package repositories

import (
	"testing"

	th "github.com/robstave/meowmorize/internal/adapters/repositories/repositories_test"
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
	assert.Equal(t, existingCard.Front.Text, card.Front.Text)
	assert.Equal(t, existingCard.Back.Text, card.Back.Text)
}

func TestCardRepositorySQLite_GetCardByID_NotFound(t *testing.T) {
	cardRepo, _ := initializeCardRepository(t)

	card, err := cardRepo.GetCardByID("non-existent-id")
	assert.NoError(t, err)
	assert.Nil(t, card)
}
