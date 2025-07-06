package repositories

import "gorm.io/gorm"

// FlashcardsRepository combines deck, card and user repositories.
type FlashcardsRepository interface {
	DeckRepository
	CardRepository
	UserRepository
}

// FlashcardsRepositorySQLite implements FlashcardsRepository using SQLite.
type FlashcardsRepositorySQLite struct {
	DeckRepository
	CardRepository
	UserRepository
}

// NewFlashcardsRepositorySQLite initializes a FlashcardsRepository backed by SQLite.
func NewFlashcardsRepositorySQLite(db *gorm.DB) FlashcardsRepository {
	return &FlashcardsRepositorySQLite{
		DeckRepository: NewDeckRepositorySQLite(db),
		CardRepository: NewCardRepositorySQLite(db),
		UserRepository: NewUserRepositorySQLite(db),
	}
}
