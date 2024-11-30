package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/robstave/meowmorize/internal/domain/types"
)

// CreateDeck handles the creation of a new deck
// @Summary Create a new deck
// @Description Create a new deck with provided details
// @Tags Decks
// @Accept  json
// @Produce  json
// @Param deck body types.Deck true "Deck to create"
// @Success 201 {object} types.Deck
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /decks [post]
func (hc *MeowController) CreateDeck(c echo.Context) error {
	var deck types.Deck
	if err := c.Bind(&deck); err != nil {
		hc.logger.Error("Failed to bind deck data", "error", err)
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid deck data",
		})
	}

	if deck.Name == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid deck data",
		})
	}

	if err := hc.service.CreateDeck(deck); err != nil {
		hc.logger.Error("Failed to create deck", "error", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to create deck",
		})
	}

	return c.JSON(http.StatusCreated, deck)
}

// CreateDeck handles the creation of a new deck
// @Summary Create a default deck
// @Description Create a new deck with provided details
// @Tags Decks
// @Accept  json
// @Produce  json
// @Param deck body types.Deck true "Deck to create"
// @Success 201 {object} types.Deck
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /decks/default [post]
func (hc *MeowController) CreateDefaultDeck(c echo.Context) error {
	hc.logger.Info("Creating default deck")

	if err := hc.service.CreateDefaultDeck(); err != nil {
		hc.logger.Error("Failed to create default deck", "error", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to create default deck",
		})
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "created default deck",
	})
}

// DeleteDeck handles the deletion of a deck
// @Summary Delete a deck
// @Description Delete a deck by its ID
// @Tags Decks
// @Param id path string true "Deck ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /decks/{id} [delete]
func (hc *MeowController) DeleteDeck(c echo.Context) error {
	deckID := c.Param("id")

	if deckID == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Deck ID is required",
		})
	}

	// Perform deletion
	err := hc.service.DeleteDeck(deckID)
	if err != nil {
		hc.logger.Error("Failed to delete deck", "error", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to delete deck",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Deck deleted successfully",
	})
}

// GetAllDecks retrieves all decks
// @Summary Get all decks
// @Description Retrieve a list of all decks
// @Tags Decks
// @Produce  json
// @Success 200 {array} types.Deck
// @Failure 500 {object} map[string]string
// @Router /decks [get]
func (hc *MeowController) GetAllDecks(c echo.Context) error {
	decks, err := hc.service.GetAllDecks()
	if err != nil {
		hc.logger.Error("Failed to retrieve decks", "error", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to retrieve decks",
		})
	}

	return c.JSON(http.StatusOK, decks)
}

// GetDeckByID retrieves a deck by its ID
// @Summary Get a deck by ID
// @Description Retrieve a single deck by its ID
// @Tags Decks
// @Produce  json
// @Param id path string true "Deck ID"
// @Success 200 {object} types.Deck
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /decks/{id} [get]
func (hc *MeowController) GetDeckByID(c echo.Context) error {
	deckID := c.Param("id")

	deck, err := hc.service.GetDeckByID(deckID)
	if err != nil {
		hc.logger.Error("Failed to get deck by ID", "error", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to retrieve deck",
		})
	}

	//if (deck == types.Deck{}) {
	//	return c.JSON(http.StatusNotFound, echo.Map{
	//		"message": "Deck not found",
	//	})
	//}

	return c.JSON(http.StatusOK, deck)
}

// UpdateDeck handles updating an existing deck
// @Summary Update a deck
// @Description Update an existing deck by its ID
// @Tags Decks
// @Accept  json
// @Produce  json
// @Param id path string true "Deck ID"
// @Param deck body types.Deck true "Updated Deck"
// @Success 200 {object} types.Deck
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /decks/{id} [put]
func (hc *MeowController) UpdateDeck(c echo.Context) error {
	deckID := c.Param("id")
	if deckID == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Deck ID is required",
		})
	}

	var updatedDeck types.Deck

	if err := c.Bind(&updatedDeck); err != nil {
		hc.logger.Error("Failed to bind deck data", "error", err)
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid deck data",
		})
	}

	if updatedDeck.Name == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid deck data",
		})
	}

	// Ensure the ID in the path matches the ID in the payload
	if updatedDeck.ID != deckID {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Deck ID mismatch",
		})
	}

	if err := hc.service.UpdateDeck(updatedDeck); err != nil {
		hc.logger.Error("Failed to update deck", "error", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to update deck",
		})
	}

	return c.JSON(http.StatusOK, updatedDeck)
}

// ExportDeck handles the export of a deck as a JSON file
// @Summary Export a deck
// @Description Export a deck as a JSON file
// @Tags decks
// @Produce application/json
// @Param id path string true "Deck ID"
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
