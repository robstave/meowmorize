// internal/adapters/repositories/deck.go
package repositories

import (
	"github.com/robstave/meowmorize/internal/domain/types"

	"gorm.io/gorm"
)

type DeckRepository interface {
	GetAllDecks() ([]types.Deck, error)
	GetAllDecksByUser(userID string) ([]types.Deck, error)
	CreateDeck(deck types.Deck) error
	DeleteDeck(deckID string) error
	GetDeckByID(deckID string) (types.Deck, error)
	UpdateDeck(deck types.Deck) error // New method
	AddCardToDeck(deckID string, card types.Card) error
	WithTransaction(fn func(txDeckRepo DeckRepository, txCardRepo CardRepository) error) error
	RemoveCardAssociation(deckID string, cardID string) error
	AddCardAssociation(deckID string, cardID string) error
}

type DeckRepositorySQLite struct {
	db *gorm.DB
}

func NewDeckRepositorySQLite(db *gorm.DB) DeckRepository {
	return &DeckRepositorySQLite{db: db}
}

func (r *DeckRepositorySQLite) GetAllDecksByUser(userID string) ([]types.Deck, error) {
	var decks []types.Deck
	if err := r.db.Preload("Cards").Where("user_id = ?", userID).Find(&decks).Error; err != nil {
		return nil, err
	}
	return decks, nil
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

func (r *DeckRepositorySQLite) AddCardToDeck(deckID string, card types.Card) error {
	var deck types.Deck
	if err := r.db.Where("id = ?", deckID).First(&deck).Error; err != nil {
		return err
	}
	// Use GORM's Association mode to append the card.
	return r.db.Model(&deck).Association("Cards").Append(&card)
}

func (r *DeckRepositorySQLite) DeleteDeck(deckID string) error {

	// Create a deck instance with the primary key set
	deck := types.Deck{ID: deckID}
	// Clear the join table entries (cards associated with the deck)
	if err := r.db.Model(&deck).Association("Cards").Clear(); err != nil {
		return err
	}

	// Delete the deck itself
	if err := r.db.Delete(&deck).Error; err != nil {
		return err
	}
	return nil
}

func (r *DeckRepositorySQLite) UpdateDeck(deck types.Deck) error {
	return r.db.Save(&deck).Error
}

// WithTransaction runs the provided function within a database transaction,
// providing transactional versions of the deck and card repositories.
func (r *DeckRepositorySQLite) WithTransaction(fn func(txDeckRepo DeckRepository, txCardRepo CardRepository) error) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		txDeckRepo := NewDeckRepositorySQLite(tx)
		txCardRepo := NewCardRepositorySQLite(tx)
		return fn(txDeckRepo, txCardRepo)
	})
}

// RemoveCardAssociation removes the association between a deck and a card.
func (r *DeckRepositorySQLite) RemoveCardAssociation(deckID string, cardID string) error {
	var deck types.Deck
	if err := r.db.Where("id = ?", deckID).First(&deck).Error; err != nil {
		return err
	}
	// Use GORM association to remove the card.
	return r.db.Model(&deck).Association("Cards").Delete(&types.Card{ID: cardID})
}

// AddCardAssociation adds an association between the   deck and the given card.
func (r *DeckRepositorySQLite) AddCardAssociation(deckID string, cardID string) error {
	var deck types.Deck
	if err := r.db.Where("id = ?", deckID).First(&deck).Error; err != nil {
		return err
	}
	return r.db.Model(&deck).Association("Cards").Append(&types.Card{ID: cardID})
}
