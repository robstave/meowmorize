// internal/adapters/controller/import.go

package controller

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/robstave/meowmorize/internal/domain/types"
)

// ImportDeck handles the import deck POST request
// @Summary Import a deck from a JSON file
// @Description Import a new deck by uploading a JSON file
// @Tags Decks
// @Accept multipart/form-data
// @Produce json
// @Param deck_file formData file true "Deck JSON File"
// @Success 201 {object} types.Deck
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /decks/import [post]
func (hc *MeowController) ImportDeck(c echo.Context) error {

	// Read JSON from file upload
	file, err := c.FormFile("deck_file")
	if err != nil {
		hc.logger.Error("Failed to read deck file", "error", err)
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Deck file is required",
		})
	}

	src, err := file.Open()
	if err != nil {
		hc.logger.Error("Failed to open file", "error", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to open deck file",
		})
	}
	defer src.Close()

	var deckData struct {
		Deck types.Deck `json:"deck"`
	}

	hc.logger.Info("Decoding deck JSON")

	if err := json.NewDecoder(src).Decode(&deckData); err != nil {
		hc.logger.Error("JSON decoding failed", "error", err)
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid JSON format",
		})
	}

	// Logging for debugging
	hc.logger.Info("Imported deck", "id", deckData.Deck.ID, "name", deckData.Deck.Name,
		"number_of_cards", len(deckData.Deck.Cards))

	for _, card := range deckData.Deck.Cards {
		if card.ID == "" || card.DeckID == "" || card.Front.Text == "" || card.Back.Text == "" {
			hc.logger.Warn("Incomplete card data", "card", card)
		} else {
			hc.logger.Info("Imported Card",
				"uuid", card.ID,
				"did", card.DeckID,
				"front", card.Front.Text,
				"back", card.Back.Text)
		}
	}

	// Assign DeckID to each card if not already set
	for i := range deckData.Deck.Cards {
		if deckData.Deck.Cards[i].DeckID == "" {
			deckData.Deck.Cards[i].DeckID = deckData.Deck.ID
		}
	}

	// Save the deck to the database
	if err := hc.service.CreateDeck(deckData.Deck); err != nil {
		hc.logger.Error("Failed to save deck", "error", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to save deck",
		})
	}

	hc.logger.Info("Deck imported successfully", "id", deckData.Deck.ID)

	return c.JSON(http.StatusCreated, deckData.Deck)
}
