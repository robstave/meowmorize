package controller

import (
	"github.com/labstack/echo/v4"
)

// DeleteDeck handles the deletion of a deck
func (hc *HomeController) DeleteDeck(c echo.Context) error {
	deckID := c.Param("id")

	// Perform deletion
	err := hc.service.DeleteDeck(deckID)

	//TODO
}
