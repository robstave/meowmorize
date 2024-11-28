// domain/cards.go
package domain

import (
	"errors"
	"fmt"

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

// New method: UpdateCard
func (s *Service) UpdateCard(card types.Card) error {
	if card.ID == "" {
		return fmt.Errorf("card ID must be provided for update")
	}
	// Additional validation or business rules can be applied here
	return s.cardRepo.UpdateCard(card)
}

func (s *Service) CreateCard(card types.Card) error {
	if card.ID == "" {
		return fmt.Errorf("card ID must be provided for update")
	}
	// Additional validation or business rules can be applied here
	return s.cardRepo.CreateCard(card)
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
