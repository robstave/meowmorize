// controllers/meow_controller.go
package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/robstave/meowmorize/internal/domain/types"
)

// GetCardByID retrieves a card by its ID
// @Summary Get a card by ID
// @Description Retrieve a single card by its ID
// @Tags Cards
// @Produce  json
// @Param id path string true "Card ID"
// @Security BearerAuth
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

// CreateCardRequest represents the expected payload for creating a card
type CreateCardRequest struct {
	DeckID string         `json:"deck_id" validate:"required,uuid"`
	Front  CardContentReq `json:"front" validate:"required"`
	Back   CardContentReq `json:"back" validate:"required"`
	Link   string         `json:"link"`
}

// CardContentReq represents the content structure for front and back of a card
type CardContentReq struct {
	Text string `json:"text" validate:"required"`
}

// UpdateCardRequest represents the expected payload for updating a card
type UpdateCardRequest struct {
	Front *CardContentReq `json:"front"`
	Back  *CardContentReq `json:"back"`
	Link  *string         `json:"link"`
}

// @Summary Create a new card
// @Description Create a new card and associate it with a deck
// @Tags cards
// @Accept json
// @Produce json
// @Param card body CreateCardRequest true "Create Card"
// @Param id path string true "deck id
// @Security BearerAuth
// @Success 201 {object} types.Card
// @Failure 400 {object} echo.HTTPError
// @Failure 500 {object} echo.HTTPError
// @Router /cards [post]
func (c *MeowController) CreateCard(ctx echo.Context) error {

	userID, err := getUserIDFromContext(ctx)
	if err != nil {
		c.logger.Error("Failed to extract user ID from token", "error", err)
		return ctx.JSON(http.StatusUnauthorized, echo.Map{"message": "unauthorized"})
	}

	deckID := ctx.Param("id")
	if deckID == "" {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": "Deck ID is required"})
	}

	// Verify that the deck belongs to the logged-in user.
	deck, err := c.service.GetDeckByID(deckID)
	if err != nil {
		c.logger.Error("Failed to retrieve deck 1", "deckID", deckID, "error", err)
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"message": "Failed to retrieve deck"})
	}
	if deck.UserID != userID {
		return ctx.JSON(http.StatusUnauthorized, echo.Map{"message": "Not authorized to add card to this deck"})
	}

	var req CreateCardRequest
	if err := ctx.Bind(&req); err != nil {
		c.logger.Error("Failed to bind request", "error", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	// Optional: Add validation here if using a validation library
	// e.g., if err := ctx.Validate(req); err != nil { ... }

	// Create a Card object
	newCard := types.Card{
		// ID will be generated by the service or repository

		Front: types.CardFront{
			Text: req.Front.Text,
		},
		Back: types.CardBack{
			Text: req.Back.Text,
		},
		Link: req.Link,
	}

	// Call the service to create the card
	ccard, err := c.service.CreateCard(newCard, deckID)
	if err != nil {
		c.logger.Error("Failed to create card", "error", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create card")
	}

	newCard.ID = ccard.ID

	c.logger.Info("Card created successfully", "deck_id", newCard.ID)
	return ctx.JSON(http.StatusCreated, newCard)
}

// @Summary Update an existing card
// @Description Update the details of an existing card by its ID
// @Tags cards
// @Accept json
// @Produce json
// @Param id path string true "Card ID"
// @Param card body UpdateCardRequest true "Update Card"
// @Security BearerAuth
// @Success 200 {object} types.Card
// @Failure 400 {object} echo.HTTPError
// @Failure 404 {object} echo.HTTPError
// @Failure 500 {object} echo.HTTPError
// @Router /cards/{id} [put]
func (c *MeowController) UpdateCard(ctx echo.Context) error {
	cardID := ctx.Param("id")
	if cardID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Card ID is required")
	}

	var req UpdateCardRequest
	if err := ctx.Bind(&req); err != nil {
		c.logger.Error("Failed to bind request", "error", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	// Retrieve the existing card
	existingCard, err := c.service.GetCardByID(cardID)
	if err != nil {
		c.logger.Error("Failed to retrieve card", "card_id", cardID, "error", err)
		return echo.NewHTTPError(http.StatusNotFound, "Card not found")
	}

	// Update the fields if they are provided
	if req.Front != nil {
		existingCard.Front.Text = req.Front.Text
	}
	if req.Back != nil {
		existingCard.Back.Text = req.Back.Text
	}
	if req.Link != nil {
		existingCard.Link = *req.Link
	}

	// Call the service to update the card
	if err := c.service.UpdateCard(*existingCard); err != nil {
		c.logger.Error("Failed to update card", "card_id", cardID, "error", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update card")
	}

	c.logger.Info("Card updated successfully", "card_id", cardID)
	return ctx.JSON(http.StatusOK, existingCard)
}

// DeleteCard deletes a card by its ID
// @Summary Delete a card
// @Description Delete a card by its ID
// @Tags Cards
// @Produce json
// @Param id path string true "Card ID"
// @Security BearerAuth
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /cards/{id} [delete]
func (c *MeowController) DeleteCard(ctx echo.Context) error {
	cardID := ctx.Param("id")

	if cardID == "" {
		c.logger.Warn("DeleteCard called without card ID")
		return echo.NewHTTPError(http.StatusBadRequest, "Card ID is required")
	}

	err := c.service.DeleteCardByID(cardID)
	if err != nil {
		// Check if the error is due to the card not being found
		if err.Error() == "card not found" {
			c.logger.Warn("Attempted to delete a non-existent card", "cardID", cardID)
			return ctx.JSON(http.StatusNotFound, echo.Map{
				"message": "Card not found",
			})
		}
		// For other errors, return a generic internal server error
		c.logger.Error("Failed to delete card", "cardID", cardID, "error", err)
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to delete card",
		})
	}

	c.logger.Info("Card deleted successfully", "cardID", cardID)
	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "Card deleted successfully",
	})
}
