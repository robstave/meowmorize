// domain/cards.go
package domain

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/robstave/meowmorize/internal/domain/types"
)

func (s *Service) GetCardByID(cardID string) (*types.Card, error) {
	s.logger.Info("Retrieving card by ID", "cardID", cardID)

	card, err := s.flashcardRepo.GetCardByID(cardID)
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

func (s *Service) CreateCard(card types.Card, deckID string, userID string) (*types.Card, error) {
	// Create the card in the cards table
	if card.ID == "" {
		card.ID = uuid.New().String()
	}
	card.UserID = userID
	if err := s.flashcardRepo.CreateCard(card); err != nil {
		return nil, err
	}
	// Associate card with the given deck (use a helper method, or directly use GORM's association API)
	if err := s.flashcardRepo.AddCardToDeck(deckID, card); err != nil {
		return nil, err
	}
	return &card, nil
}

// UpdateCard updates an existing card in the repository
func (s *Service) UpdateCard(card types.Card) error {
	// Ensure the card exists
	existingCard, err := s.flashcardRepo.GetCardByID(card.ID)
	if err != nil {
		s.logger.Error("Card not found", "card_id", card.ID, "error", err)
		return err
	}

	// Update fields
	existingCard.Front = card.Front
	existingCard.Back = card.Back
	existingCard.Link = card.Link

	// Save the updated card
	if err := s.flashcardRepo.UpdateCard(*existingCard); err != nil {
		s.logger.Error("Failed to update card", "card_id", card.ID, "error", err)
		return err
	}

	s.logger.Info("Card updated successfully", "card_id", card.ID)
	return nil
}

func (s *Service) DeleteCardByID(cardID string) error {
	if cardID == "" {
		return fmt.Errorf("card ID must be provided for deletion")
	}
	return s.flashcardRepo.DeleteCardByID(cardID)
}

func (s *Service) CloneCardToDeck(cardID string, targetDeckID string) (*types.Card, error) {
	if cardID == "" {
		return nil, fmt.Errorf("source card ID must be provided for cloning")
	}
	if targetDeckID == "" {
		return nil, fmt.Errorf("target deck ID must be provided for cloning")
	}

	// Optionally, you can add business logic here, such as verifying that the target deck exists.

	// Delegate the cloning operation to the repository
	return s.flashcardRepo.CloneCardToDeck(cardID, targetDeckID)
}

// UpdateCardStats updates the card based on the provided action
func (s *Service) UpdateCardStats(cardID string, action types.CardAction, value *int, deckID string, userID string) error {
	card, err := s.flashcardRepo.GetCardByID(cardID)
	if err != nil {
		s.logger.Error("Failed to retrieve card", "card_id", cardID, "error", err)
		return err
	}
	if card == nil {
		return errors.New("card not found")
	}

	switch action {
	case types.IncrementFail:
		card.FailCount++
	case types.IncrementPass:
		card.PassCount++
	case types.IncrementSkip:
		card.SkipCount++
	case types.SetStars:
		if value != nil {
			card.StarRating = *value
		}

	case types.Retire:
		card.Retired = true
	case types.Unretire:
		card.Retired = false
	case types.ResetStats:
		card.FailCount = 0
		card.PassCount = 0
		card.SkipCount = 0
	default:
		return fmt.Errorf("unknown action: %s", action)
	}

	// Update the ReviewedAt timestamp
	card.ReviewedAt = time.Now()

	// Update the UpdatedAt timestamp is handled by GORM automatically

	err = s.flashcardRepo.UpdateCard(*card)
	if err != nil {
		s.logger.Error("Failed to update card stats", "card_id", cardID, "error", err)
		return err
	}

	err = s.AdjustSession(deckID, cardID, action, card.StarRating, userID)
	if err != nil {
		s.logger.Error("Failed to update session", "card_id", cardID, "deck_id", deckID, "error", err)
		return err
	}

	s.logger.Info("Card stats updated successfully", "card_id", cardID, "action", action)
	return nil
}
