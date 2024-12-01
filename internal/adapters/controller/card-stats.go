// internal/adapters/controller/card.go

package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/robstave/meowmorize/internal/domain/types"
)

// CardStatsRequest represents the expected payload for updating card stats
type CardStatsRequest struct {
	CardID string           `json:"card_id" validate:"required"`
	Action types.CardAction `json:"action" validate:"required,oneof=IncrementFail IncrementPass IncrementSkip SetStars Retire Unretire ResetStats"`
	Value  *int             `json:"value,omitempty"` // Used only for SetStars
}

// @Summary Update card statistics
// @Description Update the statistics of a card based on the specified action
// @Tags Cards
// @Accept json
// @Produce json
// @Param stats body CardStatsRequest true "Card Stats Update"
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
	if err := c.service.UpdateCardStats(req.CardID, req.Action, req.Value); err != nil {
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
