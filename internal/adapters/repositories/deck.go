package repositories

import (
	"github.com/robstave/meowmorize/internal/domain/types"

	"gorm.io/gorm"
)

type DeckRepository interface {
	GetAllDecks() ([]types.Deck, error)
	CreateDeck(deck types.Deck) error
	DeleteDeck(deckID string) error // New method
	GetDeckByID(deckID string) (types.Deck, error)
}

type DeckRepositorySQLite struct {
	db *gorm.DB
}

func NewDeckRepositorySQLite(db *gorm.DB) DeckRepository {
	return &DeckRepositorySQLite{db: db}
}

func (r *DeckRepositorySQLite) GetAllDecks() ([]types.Deck, error) {
	var decks []types.Deck
	if err := r.db.Preload("Cards").Find(&decks).Error; err != nil {
		return nil, err
	}
	return decks, nil
}

func (r *DeckRepositorySQLite) CreateDeck(deck types.Deck) error {
	return r.db.Create(&deck).Error
}

func (r *DeckRepositorySQLite) GetDeckByID(deckID string) (types.Deck, error) {
	var deck types.Deck
	if err := r.db.Preload("Cards").Where("id = ?", deckID).First(&deck).Error; err != nil {
		return types.Deck{}, err
	}
	return deck, nil
}

func (r *DeckRepositorySQLite) DeleteDeck(deckID string) error {
	// Begin a transaction to ensure both deck and its cards are deleted atomically
	tx := r.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	// Delete associated cards first to maintain referential integrity
	if err := tx.Where("deck_id = ?", deckID).Delete(&types.Card{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Delete the deck
	if err := tx.Where("id = ?", deckID).Delete(&types.Deck{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Commit the transaction
	return tx.Commit().Error
}
