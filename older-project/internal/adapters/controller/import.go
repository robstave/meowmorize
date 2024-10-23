package controller

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/robstave/meowmorize/internal/domain/types"
)

// RenderImportDeck handles rendering the import deck page
func (hc *HomeController) RenderImportDeck(c echo.Context) error {
	return c.Render(http.StatusOK, "import.html", nil)
}

// ImportDeck handles the import deck POST request
func (hc *HomeController) ImportDeck(c echo.Context) error {

	hc.logger.Info("Importing deck")

	// Read JSON from file upload
	file, err := c.FormFile("deck_file")
	if err != nil {
		hc.logger.Error("Failed to read deck file", "error", err)
		return c.Redirect(http.StatusSeeOther, "/import-deck?import=error")
	}

	src, err := file.Open()
	if err != nil {
		hc.logger.Error("Failed to open file", "error", err)
		return c.Redirect(http.StatusSeeOther, "/import-deck?import=error")
	}
	defer src.Close()

	var deckData struct {
		Deck types.Deck `json:"deck"`
	}

	hc.logger.Info("Decoding")

	if err := json.NewDecoder(src).Decode(&deckData); err != nil {
		hc.logger.Error("JSON decoding failed", "error", err)
		return c.Redirect(http.StatusSeeOther, "/import-deck?import=error")
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
		return c.Redirect(http.StatusSeeOther, "/import-deck?import=error")
	}

	// Redirect to import page with success message
	return c.Redirect(http.StatusSeeOther, "/import-deck?import=success")
}
