// test_helpers.go
package repositories_test

import (
	"testing"

	"github.com/robstave/meowmorize/internal/domain/types"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// SetupTestDB initializes an in-memory SQLite database and performs migrations.
func SetupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to in-memory SQLite database: %v", err)
	}

	// Perform migrations
	err = db.AutoMigrate(&types.Card{}, &types.Deck{})
	if err != nil {
		t.Fatalf("Failed to migrate database: %v", err)
	}

	return db
}

// SeedTestData populates the database with initial data for testing.
func SeedTestData(t *testing.T, db *gorm.DB) (types.Deck, types.Card) {
	// Create a deck
	deck := types.Deck{
		ID:   "deck1",
		Name: "Sample Deck",
	}
	if err := db.Create(&deck).Error; err != nil {
		t.Fatalf("Failed to create deck: %v", err)
	}

	// Create a card
	card := types.Card{
		ID: "card1",
		Front: types.CardFront{
			Text: "Front Text",
		},
		Back: types.CardBack{
			Text: "Back Text",
		},
	}
	if err := db.Create(&card).Error; err != nil {
		t.Fatalf("Failed to create card: %v", err)
	}

	return deck, card
}
