// controllers/meow_controller.go
package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetCardByID retrieves a card by its ID
// @Summary Get a card by ID
// @Description Retrieve a single card by its ID
// @Tags Cards
// @Produce  json
// @Param id path string true "Card ID"
// @Success 200 {object} types.Card
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /cards/{id} [get]
func (hc *MeowController) GetCardByID(c echo.Context) error {
	cardID := c.Param("id")

	card, err := hc.service.GetCardByID(cardID)
	if err != nil {
		if err.Error() == "card not found" {
			hc.logger.Warn("Card not found", "cardID", cardID)
			return c.JSON(http.StatusNotFound, echo.Map{
				"message": "Card not found",
			})
		}

		hc.logger.Error("Failed to get card by ID", "error", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to retrieve card",
		})
	}

	return c.JSON(http.StatusOK, card)
}
