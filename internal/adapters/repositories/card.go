// internal/adapters/repositories/card.go
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
	GetCardByID(cardID string) (*types.Card, error)
	UpdateCard(card types.Card) error
	DeleteCardByID(cardID string) error
	CloneCardToDeck(cardID string, targetDeckID string) (*types.Card, error)
	CountDeckAssociations(cardID string) (int, error)
}

type CardRepositorySQLite struct {
	db *gorm.DB
}

func NewCardRepositorySQLite(db *gorm.DB) CardRepository {
	return &CardRepositorySQLite{db: db}
}

func (r *CardRepositorySQLite) GetCardsByDeckID(deckID string) ([]types.Card, error) {
	var deck types.Deck
	if err := r.db.Preload("Cards").First(&deck, "id = ?", deckID).Error; err != nil {
		return nil, err
	}
	return deck.Cards, nil
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
		originalCard, err := r.GetCardByID(cardID)
		if err != nil {
			return fmt.Errorf("error retrieving original card: %w", err)
		}
		if originalCard == nil {
			return fmt.Errorf("no card found with ID %s", cardID)
		}
		cloned := *originalCard
		cloned.ID = uuid.New().String() // New UUID for cloned card
		// Create the cloned card:
		if err := tx.Create(&cloned).Error; err != nil {
			return fmt.Errorf("error creating cloned card: %w", err)
		}
		newCard = &cloned
		// Associate the cloned card with the target deck:
		var deck types.Deck
		if err := tx.Where("id = ?", targetDeckID).First(&deck).Error; err != nil {
			return fmt.Errorf("target deck not found: %w", err)
		}
		if err := tx.Model(&deck).Association("Cards").Append(newCard); err != nil {
			return fmt.Errorf("error associating cloned card with deck: %w", err)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return newCard, nil
}

// CountDeckAssociations returns the number of decks a card is associated with.
func (r *CardRepositorySQLite) CountDeckAssociations(cardID string) (int, error) {
	var card types.Card
	if err := r.db.Where("id = ?", cardID).First(&card).Error; err != nil {
		return 0, err
	}
	count := r.db.Model(&card).Association("Decks").Count()
	return int(count), nil
}
