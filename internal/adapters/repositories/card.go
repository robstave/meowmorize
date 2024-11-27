package repositories

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/robstave/meowmorize/internal/domain/types"

	"gorm.io/gorm"
)

type CardRepository interface {
	GetCardsByDeckID(deckID string) ([]types.Card, error)
	CreateCard(card types.Card) error
	GetCardByID(cardID string) (*types.Card, error) // New method added
	UpdateCard(card types.Card) error
	DeleteCardByID(cardID string) error
	CloneCardToDeck(cardID string, targetDeckID string) (*types.Card, error)
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

func (r *CardRepositorySQLite) GetCardByID(cardID string) (*types.Card, error) {
	var card types.Card
	if err := r.db.First(&card, "id = ?", cardID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil // Or handle as per your application's error handling strategy
		}
		return nil, err
	}
	return &card, nil
}

func (r *CardRepositorySQLite) UpdateCard(card types.Card) error {
	if card.ID == "" {
		return fmt.Errorf("card ID is required for update")
	}
	return r.db.Save(&card).Error
}

func (r *CardRepositorySQLite) DeleteCardByID(cardID string) error {
	if cardID == "" {
		return fmt.Errorf("card ID is required for deletion")
	}
	result := r.db.Delete(&types.Card{}, "id = ?", cardID)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("no card found with ID %s", cardID)
	}
	return nil
}

func (r *CardRepositorySQLite) CloneCardToDeck(cardID string, targetDeckID string) (*types.Card, error) {
	if cardID == "" {
		return nil, fmt.Errorf("source card ID is required for cloning")
	}
	if targetDeckID == "" {
		return nil, fmt.Errorf("target deck ID is required for cloning")
	}

	var newCard *types.Card

	err := r.db.Transaction(func(tx *gorm.DB) error {
		// Retrieve the original card
		originalCard, err := r.GetCardByID(cardID)
		if err != nil {
			return fmt.Errorf("error retrieving original card: %w", err)
		}
		if originalCard == nil {
			return fmt.Errorf("no card found with ID %s", cardID)
		}

		// Create a copy of the original card
		cloned := *originalCard
		cloned.ID = uuid.New().String() // Generate a new UUID
		cloned.DeckID = targetDeckID    // Assign to the target deck

		// Create the new card in the database
		if err := tx.Create(&cloned).Error; err != nil {
			return fmt.Errorf("error creating cloned card: %w", err)
		}

		newCard = &cloned
		return nil
	})

	if err != nil {
		return nil, err
	}

	return newCard, nil
}
