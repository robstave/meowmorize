package controller

import (
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/robstave/meowmorize/internal/domain/types"
)

// Helper: extract userID from JWT claims.
func getUserIDFromContext(c echo.Context) (string, error) {
	user := c.Get("user")
	if user == nil {
		return "", echo.NewHTTPError(http.StatusUnauthorized, "user not found in context")
	}
	// JWT claims are stored in user (of type *jwt.Token)
	claims, ok := user.(*jwt.Token).Claims.(jwt.MapClaims)
	if !ok {
		return "", echo.NewHTTPError(http.StatusUnauthorized, "invalid JWT claims")
	}
	userID, ok := claims["username"].(string)
	if !ok || userID == "" {
		return "", echo.NewHTTPError(http.StatusUnauthorized, "user_id not found in token")
	}
	return userID, nil
}

// CreateDeck handles deck creation ensuring the deck is owned by the logged-in user.
// @Summary Create a new deck
// @Description Create a new deck owned by the authenticated user
// @Tags Decks
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param deck body types.Deck true "Deck to create"
// @Success 201 {object} types.Deck
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /decks [post]
func (hc *MeowController) CreateDeck(c echo.Context) error {
	userID, err := getUserIDFromContext(c)
	if err != nil {
		hc.logger.Error("Failed to extract user id from token", "error", err)
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": "unauthorized"})
	}

	var deck types.Deck
	if err := c.Bind(&deck); err != nil {
		hc.logger.Error("Failed to bind deck data", "error", err)
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid deck data"})
	}

	// Set the owner of the deck
	deck.UserID = userID

	// Call the service to create the deck
	if err := hc.service.CreateDeck(deck); err != nil {
		hc.logger.Error("Failed to create deck", "error", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Failed to create deck"})
	}
	return c.JSON(http.StatusCreated, deck)
}

// UpdateDeck ensures that the deck being updated belongs to the logged-in user.
// @Summary Update a deck
// @Description Update an existing deck, verifying ownership
// @Tags Decks
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Deck ID"
// @Param deck body types.Deck true "Updated Deck"
// @Success 200 {object} types.Deck
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /decks/{id} [put]
func (hc *MeowController) UpdateDeck(c echo.Context) error {
	userID, err := getUserIDFromContext(c)
	if err != nil {
		hc.logger.Error("Failed to extract user id from token", "error", err)
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": "unauthorized"})
	}

	deckID := c.Param("id")
	if deckID == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Deck ID is required"})
	}

	var req types.Deck
	if err := c.Bind(&req); err != nil {
		hc.logger.Error("Failed to bind deck update data", "error", err)
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid deck data"})
	}

	// Retrieve existing deck to verify ownership
	existingDeck, err := hc.service.GetDeckByID(deckID)
	if err != nil {
		hc.logger.Error("Failed to retrieve deck ctr", "deckID", deckID, "error", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Failed to retrieve deck"})
	}
	if existingDeck.UserID != userID {
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": "Not authorized to update this deck"})
	}

	// Update deck fields; ensure ownership is not modified
	existingDeck.Name = req.Name
	existingDeck.Description = req.Description
	// Note: Cards association may be handled via a separate endpoint

	if err := hc.service.UpdateDeck(existingDeck); err != nil {
		hc.logger.Error("Failed to update deck", "deckID", deckID, "error", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Failed to update deck"})
	}

	return c.JSON(http.StatusOK, existingDeck)
}

// DeleteDeck verifies deck ownership before deletion.
// @Summary Delete a deck
// @Description Delete a deck owned by the authenticated user
// @Tags Decks
// @Security BearerAuth
// @Param id path string true "Deck ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /decks/{id} [delete]
func (hc *MeowController) DeleteDeck(c echo.Context) error {
	userID, err := getUserIDFromContext(c)
	if err != nil {
		hc.logger.Error("Failed to extract user id from token", "error", err)
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": "unauthorized"})
	}

	deckID := c.Param("id")
	if deckID == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Deck ID is required"})
	}

	// Retrieve the deck to verify ownership
	existingDeck, err := hc.service.GetDeckByID(deckID)
	if err != nil {
		hc.logger.Error("Failed to retrieve deck 2", "deckID", deckID, "error", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Failed to retrieve deck"})
	}
	if existingDeck.UserID != userID {
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": "Not authorized to delete this deck"})
	}

	if err := hc.service.DeleteDeck(deckID); err != nil {
		hc.logger.Error("Failed to delete deck controller", "deckID", deckID, "error", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Failed to delete deck"})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "Deck deleted successfully"})
}

// GetAllDecks retrieves all decks for the logged-in user.
// @Summary Get all decks for the logged-in user
// @Description Retrieve a list of all decks owned by the authenticated user
// @Tags Decks
// @Produce json
// @Security BearerAuth
// @Success 200 {array} types.Deck
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /decks [get]
func (hc *MeowController) GetAllDecks(c echo.Context) error {
	userID, err := getUserIDFromContext(c)
	if err != nil {
		hc.logger.Error("Failed to extract user id from token", "error", err)
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": "unauthorized"})
	}

	decks, err := hc.service.GetAllDecks(userID)
	if err != nil {
		hc.logger.Error("Failed to retrieve decks 3", "error", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Failed to retrieve decks"})
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
