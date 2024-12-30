// internal/adapters/controller/import.go

package controller

import (
	"encoding/json"
	"fmt"
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
// @Security BearerAuth
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

	var deck types.Deck

	hc.logger.Info("Decoding deck JSON")

	if err := json.NewDecoder(src).Decode(&deck); err != nil {
		hc.logger.Error("JSON decoding failed", "error", err)
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid JSON format",
		})
	}

	// Logging for debugging
	hc.logger.Info("Imported deck", "id", deck.ID, "name", deck.Name,
		"number_of_cards", len(deck.Cards))

	for _, card := range deck.Cards {
		if card.ID == "" || card.DeckID == "" || card.Front.Text == "" || card.Back.Text == "" {
			hc.logger.Warn("Incomplete card data", "card", card)
		} else {

			hc.logger.Info("Imported Card",
				"uuid", card.ID,
				"did", card.DeckID,
				"link", card.Link,
				"front", card.Front.Text,
				"back", card.Back.Text)
		}
	}

	// Assign DeckID to each card if not already set
	for i := range deck.Cards {
		if deck.Cards[i].DeckID == "" {
			deck.Cards[i].DeckID = deck.ID
		}

	}

	// Save the deck to the database
	if err := hc.service.CreateDeck(deck); err != nil {
		hc.logger.Error("Failed to save deck", "error", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to save deck",
		})
	}

	hc.logger.Info("Deck imported successfully", "id", deck.ID)

	return c.JSON(http.StatusCreated, deck)
}

// ExportDeck handles the export of a deck as a JSON file
// @Summary Export a deck
// @Description Export a deck as a JSON file
// @Tags decks
// @Produce application/json
// @Param id path string true "Deck ID"
// @Security BearerAuth
// @Success 200 {object} types.Deck
// @Failure 400 {object} echo.HTTPError
// @Failure 500 {object} echo.HTTPError
// @Router /decks/export/{id} [get]
func (c *MeowController) ExportDeck(ctx echo.Context) error {
	deckID := ctx.Param("id")
	if deckID == "" {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": "Deck ID is required",
		})
	}

	deck, err := c.service.ExportDeck(deckID)
	if err != nil {
		c.logger.Error("Failed to export deck", "deck_id", deckID, "error", err)
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to export deck",
		})
	}

	// Marshal the deck into pretty JSON
	deckJSON, err := json.MarshalIndent(deck, "", "  ") // Indent with two spaces
	if err != nil {
		c.logger.Error("Failed to marshal deck", "deck_id", deckID, "error", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to export deck")
	}

	// Set headers to prompt file download
	ctx.Response().Header().Set("Content-Type", "application/json")
	ctx.Response().Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"deck-%s.json\"", deckID))

	return ctx.JSONBlob(http.StatusOK, deckJSON)
}
