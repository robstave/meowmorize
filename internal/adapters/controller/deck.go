package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (hc *HomeController) RenderDeckPage(c echo.Context) error {
	deckID := c.Param("id")
	deck, err := hc.service.GetDeckByID(deckID)
	if err != nil {
		hc.logger.Error("Failed to retrieve deck", "deck_id", deckID, "error", err)
		return c.String(http.StatusInternalServerError, "Failed to retrieve deck")
	}

	if deck.ID == "" {
		return c.String(http.StatusNotFound, "Deck not found")
	}

	data := map[string]interface{}{
		"Deck": deck,
	}

	return c.Render(http.StatusOK, "deck.html", data)
}

// DeleteDeck handles the deletion of a deck
func (hc *HomeController) DeleteDeck(c echo.Context) error {
	deckID := c.Param("id")

	// Perform deletion
	err := hc.service.DeleteDeck(deckID)
	if err != nil {
		hc.logger.Error("Failed to delete deck", "deck_id", deckID, "error", err)
		return c.Redirect(http.StatusSeeOther, "/deck/"+deckID+"?delete=error")
	}

	// Redirect to dashboard with success message
	return c.Redirect(http.StatusSeeOther, "/?delete=success")
}
