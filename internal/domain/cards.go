// domain/cards.go
package domain

import (
	"errors"

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
