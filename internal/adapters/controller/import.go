package controller

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/robstave/meowmorize/internal/domain/types"
)

func (hc *HomeController) ImportDeck(c echo.Context) error {
	// Read JSON from file upload
	file, err := c.FormFile("deck_file")
	if err != nil {
		return c.String(http.StatusBadRequest, "Failed to read deck file")
	}

	src, err := file.Open()
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to open file")
	}
	defer src.Close()

	var deckData struct {
		Deck types.Deck `json:"deck"`
	}

	if err := json.NewDecoder(src).Decode(&deckData); err != nil {
		return c.String(http.StatusBadRequest, "Failed to parse JSON")
	}

	hc.logger.Info("Imported deck", "id", deckData.Deck.ID, "name", deckData.Deck.Name, "len", len(deckData.Deck.Cards))
	for _, card := range deckData.Deck.Cards {
		hc.logger.Info("Imported Card",
			"uuid", card.ID,
			"did", card.DeckID,
			"front", card.Front.Text,
			"back", card.Back.Text)
	}

	deckID, err := uuid.Parse(deckData.Deck.ID)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid deck ID format")
	}
	deckData.Deck.ID = deckID.String()

	// Save the deck to the database
	if err := hc.service.CreateDeck(deckData.Deck); err != nil {
		return c.String(http.StatusInternalServerError, "Failed to save deck")
	}

	return c.String(http.StatusOK, "Deck imported successfully")
}
