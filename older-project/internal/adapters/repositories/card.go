package repositories

import (
	"github.com/robstave/meowmorize/internal/domain/types"

	"gorm.io/gorm"
)

type CardRepository interface {
	GetCardsByDeckID(deckID string) ([]types.Card, error)
	CreateCard(card types.Card) error
}

type CardRepositorySQLite struct {
	db *gorm.DB
}

func NewCardRepositorySQLite(db *gorm.DB) CardRepository {
	return &CardRepositorySQLite{db: db}
}

func (r *CardRepositorySQLite) GetCardsByDeckID(deckID string) ([]types.Card, error) {
	var cards []types.Card
	if err := r.db.Where("deck_id = ?", deckID).Find(&cards).Error; err != nil {
		return nil, err
	}
	return cards, nil
}

func (r *CardRepositorySQLite) CreateCard(card types.Card) error {
	return r.db.Create(&card).Error
}
