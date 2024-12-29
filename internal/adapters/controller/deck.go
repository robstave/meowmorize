package controller

import (
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
// @Security BearerAuth
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

// DeleteDeck handles the deletion of a deck
// @Summary Delete a deck
// @Description Delete a deck by its ID
// @Tags Decks
// @Param id path string true "Deck ID"
// @Security BearerAuth
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
// @Security BearerAuth
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
// @Security BearerAuth
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

// UpdateDeckRequest represents the expected payload for updating a deck
type UpdateDeckRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

// UpdateDeck handles updating an existing deck
// @Summary Update a deck
// @Description Update the name (and optionally the ID) of an existing deck by its ID
// @Tags Decks
// @Accept  json
// @Produce  json
// @Param id path string true "Deck ID"
// @Security BearerAuth
// @Param deck body UpdateDeckRequest true "Updated Deck"
// @Success 200 {object} types.Deck
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /decks/{id} [put]
func (hc *MeowController) UpdateDeck(c echo.Context) error {
	deckID := c.Param("id")

	if deckID == "" {
		hc.logger.Warn("UpdateDeck called without deck ID")
		return echo.NewHTTPError(http.StatusBadRequest, "Deck ID is required")
	}

	var req UpdateDeckRequest
	if err := c.Bind(&req); err != nil {
		hc.logger.Error("Failed to bind UpdateDeckRequest", "error", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid deck data")
	}

	// Optional: Add validation here if using a validation library
	// e.g., if err := c.Validate(req); err != nil { ... }

	// Fetch the existing deck
	existingDeck, err := hc.service.GetDeckByID(deckID)
	if err != nil {
		hc.logger.Error("Failed to retrieve deck", "deckID", deckID, "error", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve deck")
	}

	// Check if the deck exists
	if existingDeck.ID == "" {
		hc.logger.Warn("Deck not found", "deckID", deckID)
		return c.JSON(http.StatusNotFound, echo.Map{
			"message": "Deck not found",
		})
	}

	// Update only the Name field
	existingDeck.Name = req.Name
	existingDeck.Description = req.Description

	// Call the service to update the deck
	if err := hc.service.UpdateDeck(existingDeck); err != nil {
		hc.logger.Error("Failed to update deck", "deckID", deckID, "error", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update deck")
	}

	hc.logger.Info("Deck updated successfully", "deckID", deckID)
	return c.JSON(http.StatusOK, existingDeck)
}

// CollapseDecksRequest represents the expected payload for collapsing decks
type CollapseDecksRequest struct {
	TargetDeckID string `json:"target_deck_id" validate:"required,uuid"`
	SourceDeckID string `json:"source_deck_id" validate:"required,uuid"`
}

// @Summary Collapse two decks
// @Description Merge all cards from the source deck into the target deck by removing each card from the source deck, deleting it, and adding it to the target deck.
// @Tags Decks
// @Accept  json
// @Produce  json
// @Param collapse body CollapseDecksRequest true "Deck IDs to collapse"
// @Security BearerAuth
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /decks/collapse [post]
func (hc *MeowController) CollapseDecks(c echo.Context) error {
	var req CollapseDecksRequest
	if err := c.Bind(&req); err != nil {
		hc.logger.Error("Failed to bind collapse decks request", "error", err)
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid request payload",
		})
	}

	// Optional: Add validation here if using a validation library
	// e.g., if err := c.Validate(req); err != nil { ... }

	// Call the service to collapse decks
	if err := hc.service.CollapseDecks(req.TargetDeckID, req.SourceDeckID); err != nil {
		hc.logger.Error("Failed to collapse decks", "targetDeckID", req.TargetDeckID, "sourceDeckID", req.SourceDeckID, "error", err)
		// Determine the type of error to return appropriate HTTP status codes
		if err.Error() == "deck with ID "+req.SourceDeckID+" not found" || err.Error() == "deck with ID "+req.TargetDeckID+" not found" {
			return c.JSON(http.StatusNotFound, echo.Map{
				"message": "One or both decks not found",
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to collapse decks",
		})
	}

	hc.logger.Info("Decks collapsed successfully", "targetDeckID", req.TargetDeckID, "sourceDeckID", req.SourceDeckID)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Decks collapsed successfully",
	})
}
