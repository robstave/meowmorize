package domain

import (
	"github.com/robstave/meowmorize/internal/domain/types"

	"github.com/google/uuid"
)

// CreateDefaultDeck creates a default deck with a sample card
func (s *Service) CreateDefaultDeck(defaultData bool, userID string) (types.Deck, error) {
	// Generate a UUID for the default deck
	deckID := uuid.New().String()

	s.logger.Info("CreateDefaultDeck")

	defaultDeck := types.Deck{
		ID:          deckID,
		UserID:      userID,
		Name:        "Default Deck",
		Description: "This is the default deck containing basic cards.",
	}

	if defaultData {

		defaultDeck.Cards = []types.Card{
			{
				ID: uuid.New().String(), // Generate a UUID for the card

				Front: struct {
					Text string `gorm:"type:text;not null" json:"text"`
				}{
					Text: "Capital of France",
				},
				Back: struct {
					Text string `gorm:"type:text;not null" json:"text"`
				}{
					Text: "Paris",
				},
			},
			{
				ID: uuid.New().String(), // Generate a UUID for the card

				Front: struct {
					Text string `gorm:"type:text;not null" json:"text"`
				}{
					Text: "Capital of Norway",
				},
				Back: struct {
					Text string `gorm:"type:text;not null" json:"text"`
				}{
					Text: "Oslo",
				},
			},
		}
	} else {
		defaultDeck.Cards = []types.Card{}
	}

	// Save the default deck to the database
	err := s.deckRepo.CreateDeck(defaultDeck)
	if err != nil {
		s.logger.Error("Failed to create default deck", "error", err)
		return defaultDeck, err
	}

	s.logger.Info("Default deck created successfully")
	return defaultDeck, nil
}
