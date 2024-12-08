package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// CreateDeckRequest represents the request payload for creating a deck
type CreateDeckRequest struct {
	DefaultData bool `json:"defaultData"`
}

// CreateDeck handles the creation of a new deck, optionally with default data
// @Summary Create a new deck
// @Description Create a new deck with or without default data
// @Tags Decks
// @Accept  json
// @Produce  json
// @Param deck body CreateDeckRequest true "Deck creation parameters"
// @Success 200 {object} types.Deck
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /decks/default [post]
func (hc *MeowController) CreateDefaultDeck(c echo.Context) error {
	hc.logger.Info("Creating a new deck")

	var req CreateDeckRequest
	if err := c.Bind(&req); err != nil {
		hc.logger.Error("Invalid request payload", "error", err)
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid request payload",
		})
	}

	deck, err := hc.service.CreateDefaultDeck(req.DefaultData)
	if err != nil {
		hc.logger.Error("Failed to create deck", "error", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to create deck",
		})
	}

	return c.JSON(http.StatusOK, deck)

}
