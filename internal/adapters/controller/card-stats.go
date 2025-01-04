// internal/adapters/controller/card.go

package controller

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/robstave/meowmorize/internal/domain/types"
)

// CardStatsRequest represents the expected payload for updating card stats
type CardStatsRequest struct {
	CardID string           `json:"card_id" validate:"required"`
	DeckID string           `json:"deck_id,omitempty"`
	Action types.CardAction `json:"action" validate:"required,oneof=IncrementFail IncrementPass IncrementSkip SetStars Retire Unretire ResetStats"`
	Value  *int             `json:"value,omitempty"` // Used only for SetStars
}

// @Summary Update card statistics
// @Description Update the statistics of a card based on the specified action
// @Tags Cards
// @Accept json
// @Produce json
// @Param stats body CardStatsRequest true "Card Stats Update"
// @Security BearerAuth
// @Success 200 {object} types.Card
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /cards/stats [post]
func (c *MeowController) UpdateCardStats(ctx echo.Context) error {
	var req CardStatsRequest
	if err := ctx.Bind(&req); err != nil {
		c.logger.Error("Failed to bind stats request", "error", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	// Optional: Add validation here if using a validation library
	// e.g., if err := ctx.Validate(req); err != nil { ... }

	// Update the card stats
	// WE are passing the deckID in case we want to update the session too
	if err := c.service.UpdateCardStats(req.CardID, req.Action, req.Value, req.DeckID); err != nil {
		if err.Error() == "card not found" {
			c.logger.Warn("Card not found", "card_id", req.CardID)
			return ctx.JSON(http.StatusNotFound, echo.Map{
				"message": "Card not found",
			})
		}
		c.logger.Error("Failed to update card stats", "error", err)
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to update card statistics",
		})
	}

	// Retrieve the updated card to return
	updatedCard, err := c.service.GetCardByID(req.CardID)
	if err != nil {
		c.logger.Error("Failed to retrieve updated card", "card_id", req.CardID, "error", err)
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to retrieve updated card",
		})
	}

	return ctx.JSON(http.StatusOK, updatedCard)
}

// ClearDeckStatsRequest represents the expected payload for clearing deck statistics
type ClearDeckStatsRequest struct {
	ClearSession bool `json:"clearSession" validate:"required"`
	ClearStats   bool `json:"clearStats" validate:"required"`
}

// @Summary Clear deck statistics
// @Description Clears the statistics for a specific deck. Can optionally clear session data and/or card statistics.
// @Tags Decks
// @Accept json
// @Produce json
// @Param id path string true "Deck ID"
// @Param stats body ClearDeckStatsRequest true "Clear Deck Statistics"
// @Security BearerAuth
// @Success 200 {object} map[string]string
// @Failure 400 {object} echo.HTTPError
// @Failure 404 {object} echo.HTTPError
// @Failure 500 {object} echo.HTTPError
// @Router /decks/stats/{id} [post]
func (c *MeowController) ClearDeckStats(ctx echo.Context) error {
	deckID := ctx.Param("id")
	if deckID == "" {
		c.logger.Warn("ClearDeckStats called without deck ID")
		return echo.NewHTTPError(http.StatusBadRequest, "Deck ID is required")
	}

	var req ClearDeckStatsRequest
	if err := ctx.Bind(&req); err != nil {
		c.logger.Error("Failed to bind ClearDeckStatsRequest", "error", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	// Optional: Add validation here if using a validation library
	// e.g., if err := ctx.Validate(req); err != nil { ... }

	// Call the service to clear deck statistics
	if err := c.service.ClearDeckStats(deckID, req.ClearSession, req.ClearStats); err != nil {
		// Determine the type of error to return appropriate HTTP status codes
		if err.Error() == fmt.Sprintf("deck with ID %s not found", deckID) {
			c.logger.Warn("Deck not found", "deckID", deckID)
			return ctx.JSON(http.StatusNotFound, echo.Map{
				"message": "Deck not found",
			})
		}

		c.logger.Error("Failed to clear deck statistics", "deckID", deckID, "error", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to clear deck statistics")
	}

	c.logger.Info("Deck statistics cleared successfully", "deckID", deckID)
	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "Deck statistics cleared successfully",
	})
}
