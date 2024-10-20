package domain

import (
	"github.com/robstave/meowmorize/internal/domain/types"

	"github.com/google/uuid"
)

// CreateDefaultDeck creates a default deck with a sample card
func (s *Service) CreateDefaultDeck() error {
	// Generate a UUID for the default deck
	deckID := uuid.New().String()

	defaultDeck := types.Deck{
		ID:          deckID,
		Name:        "Default Deck",
		Description: "This is the default deck containing basic flashcards.",

		Cards: []types.Card{
			{
				ID:     uuid.New().String(), // Generate a UUID for the card
				DeckID: deckID,

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
		},
	}

	// Save the default deck to the database
	err := s.deckRepo.CreateDeck(defaultDeck)
	if err != nil {
		s.logger.Error("Failed to create default deck", "error", err)
		return err
	}

	s.logger.Info("Default deck created successfully")
	return nil
}
