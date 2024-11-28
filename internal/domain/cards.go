// domain/cards.go
package domain

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/robstave/meowmorize/internal/domain/types"
)

func (s *Service) GetCardByID(cardID string) (*types.Card, error) {
	s.logger.Info("Retrieving card by ID", "cardID", cardID)

	card, err := s.cardRepo.GetCardByID(cardID)
	if err != nil {
		s.logger.Error("Error retrieving card", "error", err)
		return nil, err
	}

	if card == nil {
		s.logger.Warn("Card not found", "cardID", cardID)
		return nil, errors.New("card not found")
	}

	return card, nil
}

// CreateCard adds a new card to the repository
func (s *Service) CreateCard(card types.Card) (*types.Card, error) {
	// Generate a UUID for the card if not already set
	if card.ID == "" {
		card.ID = uuid.New().String()
	}

	// Validate DeckID exists
	_, err := s.deckRepo.GetDeckByID(card.DeckID)
	if err != nil {
		s.logger.Error("Deck not found", "deck_id", card.DeckID, "error", err)
		return nil, err
	}

	// Create the card
	if err := s.cardRepo.CreateCard(card); err != nil {
		s.logger.Error("Failed to create card", "error", err)
		return nil, err
	}

	s.logger.Info("Card created successfully", "card_id", card.ID)
	return &card, nil
}

// UpdateCard updates an existing card in the repository
func (s *Service) UpdateCard(card types.Card) error {
	// Ensure the card exists
	existingCard, err := s.cardRepo.GetCardByID(card.ID)
	if err != nil {
		s.logger.Error("Card not found", "card_id", card.ID, "error", err)
		return err
	}

	// Update fields
	existingCard.Front = card.Front
	existingCard.Back = card.Back
	existingCard.Link = card.Link

	// Save the updated card
	if err := s.cardRepo.UpdateCard(*existingCard); err != nil {
		s.logger.Error("Failed to update card", "card_id", card.ID, "error", err)
		return err
	}

	s.logger.Info("Card updated successfully", "card_id", card.ID)
	return nil
}

// New method: DeleteCardByID
func (s *Service) DeleteCardByID(cardID string) error {
	if cardID == "" {
		return fmt.Errorf("card ID must be provided for deletion")
	}
	return s.cardRepo.DeleteCardByID(cardID)
}

// New method: CloneCardToDeck
func (s *Service) CloneCardToDeck(cardID string, targetDeckID string) (*types.Card, error) {
	if cardID == "" {
		return nil, fmt.Errorf("source card ID must be provided for cloning")
	}
	if targetDeckID == "" {
		return nil, fmt.Errorf("target deck ID must be provided for cloning")
	}

	// Optionally, you can add business logic here, such as verifying that the target deck exists.

	// Delegate the cloning operation to the repository
	return s.cardRepo.CloneCardToDeck(cardID, targetDeckID)
}
