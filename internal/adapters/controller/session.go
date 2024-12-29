// internal/adapters/controller/session.go
package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/robstave/meowmorize/internal/domain/types"
)

// StartSessionRequest represents the expected payload for starting a session
type StartSessionRequest struct {
	DeckID string              `json:"deck_id" validate:"required,uuid"`
	Count  int                 `json:"count" validate:"min=1"`
	Method types.SessionMethod `json:"method" validate:"required,oneof=Random Fails Skips Worst"`
}

// StartSession handles the initiation of a new review session for a deck
// @Summary Start a new review session
// @Description Initiate a new review session for a specific deck with the given parameters
// @Tags Sessions
// @Accept  json
// @Produce  json
// @Param session body StartSessionRequest true "Session Parameters"
// @Security BearerAuth
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /sessions/start [post]
func (hc *MeowController) StartSession(c echo.Context) error {
	var req StartSessionRequest
	if err := c.Bind(&req); err != nil {
		hc.logger.Error("Failed to bind start session request", "error", err)
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid request payload",
		})
	}

	// Optional: Add validation here if using a validation library
	// e.g., if err := c.Validate(req); err != nil { ... }

	// Start the session
	if err := hc.service.StartSession(req.DeckID, req.Count, req.Method); err != nil {
		// You can handle specific errors if your service returns them
		hc.logger.Error("Failed to start session", "error", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to start session",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Session started successfully",
	})
}

// GetNextCardResponse represents the response containing the next card ID
type GetNextCardResponse struct {
	CardID string `json:"card_id"`
}

// GetNextCard retrieves the next card ID in the current session
// @Summary Get the next card in the session
// @Description Retrieve the ID of the next card to review in the current session
// @Tags Sessions
// @Produce  json
// @Param deck_id query string true "Deck ID"
// @Success 200 {object} GetNextCardResponse
// @Security BearerAuth
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /sessions/next [get]
func (hc *MeowController) GetNextCard(c echo.Context) error {
	deckID := c.QueryParam("deck_id")
	if deckID == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Deck ID is required",
		})
	}

	cardID, err := hc.service.GetNextCard(deckID)
	if err != nil {
		if err.Error() == "session does not exist for the given deck" {
			hc.logger.Warn("Session not found for deck", "deck_id", deckID)
			return c.JSON(http.StatusNotFound, echo.Map{
				"message": "Session not found for the given deck",
			})
		}
		hc.logger.Error("Failed to get next card", "deck_id", deckID, "error", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to retrieve next card",
		})
	}

	if cardID == "" {
		return c.JSON(http.StatusNotFound, echo.Map{
			"message": "No more cards in the session",
		})
	}

	return c.JSON(http.StatusOK, GetNextCardResponse{
		CardID: cardID,
	})
}

// ClearSession handles the termination of a review session for a deck
// @Summary Clear a review session
// @Description Terminate and clear the current review session for a specific deck
// @Tags Sessions
// @Produce  json
// @Param deck_id query string true "Deck ID"
// @Security BearerAuth
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /sessions/clear [delete]
func (hc *MeowController) ClearSession(c echo.Context) error {
	deckID := c.QueryParam("deck_id")
	if deckID == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Deck ID is required",
		})
	}

	err := hc.service.ClearSession(deckID)
	if err != nil {
		if err.Error() == "session does not exist for the given deck" {
			hc.logger.Warn("Attempted to clear a non-existent session", "deck_id", deckID)
			return c.JSON(http.StatusNotFound, echo.Map{
				"message": "Session not found for the given deck",
			})
		}
		hc.logger.Error("Failed to clear session", "deck_id", deckID, "error", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to clear session",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Session cleared successfully",
	})
}

// GetSessionStatsResponse represents the session statistics
type GetSessionStatsResponse struct {
	TotalCards   int               `json:"total_cards"`
	ViewedCount  int               `json:"viewed_count"`
	Remaining    int               `json:"remaining"`
	CurrentIndex int               `json:"current_index"`
	CardStats    []types.CardStats `json:"card_stats"`
}

// GetSessionStats retrieves the statistics of the current session for a deck
// @Summary Get session statistics
// @Description Retrieve the statistics of the current review session for a specific deck
// @Tags Sessions
// @Produce  json
// @Param deck_id query string true "Deck ID"
// @Security BearerAuth
// @Success 200 {object} GetSessionStatsResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /sessions/stats [get]
func (hc *MeowController) GetSessionStats(c echo.Context) error {
	deckID := c.QueryParam("deck_id")
	if deckID == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Deck ID is required",
		})
	}

	stats, err := hc.service.GetSessionStats(deckID)
	if err != nil {
		if err.Error() == "session does not exist for the given deck" {
			hc.logger.Warn("Session not found for deck", "deck_id", deckID)
			return c.JSON(http.StatusNotFound, echo.Map{
				"message": "Session not found for the given deck",
			})
		}
		hc.logger.Error("Failed to get session stats", "deck_id", deckID, "error", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to retrieve session statistics",
		})
	}

	return c.JSON(http.StatusOK, GetSessionStatsResponse{
		TotalCards:   stats.TotalCards,
		ViewedCount:  stats.ViewedCount,
		Remaining:    stats.Remaining,
		CurrentIndex: stats.CurrentIndex,
		CardStats:    stats.CardStats,
	})
}
