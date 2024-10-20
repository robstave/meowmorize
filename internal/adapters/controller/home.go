package controller

import (
	"log/slog"
	"net/http"

	"github.com/robstave/meowmorize/internal/domain"

	"github.com/labstack/echo/v4"
)

type HomeController struct {
	service domain.BLL
	logger  *slog.Logger
}

func NewHomeController(service domain.BLL, logger *slog.Logger) *HomeController {
	return &HomeController{service: service, logger: logger}
}

// Home handles the home page and lists all decks
func (hc *HomeController) Home(c echo.Context) error {
	decks, err := hc.service.GetAllDecks()
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to retrieve decks")
	}

	// If no decks exist, create a default deck
	if len(decks) == 0 {
		err := hc.service.CreateDefaultDeck()
		if err != nil {
			return c.String(http.StatusInternalServerError, "Failed to create default deck")
		}
		decks, _ = hc.service.GetAllDecks()
	}

	data := map[string]interface{}{
		"Decks": decks,
	}

	return c.Render(http.StatusOK, "home.html", data)
}
