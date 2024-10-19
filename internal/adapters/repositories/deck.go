package repositories

import (
	"github.com/robstave/meowmorize/internal/domain/types"

	"gorm.io/gorm"
)

type DeckRepository interface {
	GetAllDecks() ([]types.Deck, error)
	CreateDeck(deck types.Deck) error
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
