// internal/adapters/controller/llm.go

package controller

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// LLMRequest represents the request for LLM explanations
type LLMRequest struct {
	CardID string `json:"card_id" validate:"required"`
	Prompt string `json:"prompt" validate:"required"`
}

// @Summary Get LLM explanation for a card
// @Description Get an AI-generated explanation for a flashcard based on the provided prompt
// @Tags Cards
// @Accept json
// @Produce json
// @Param request body LLMRequest true "LLM Request"
// @Security BearerAuth
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /cards/explain [post]
func (c *MeowController) ExplainCard(ctx echo.Context) error {
	c.logger.Info("fff")
	cardID := ctx.Param("id")
	if cardID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Card ID is required")
	}

	// Get the card first to validate it exists and get its content
	card, err := c.service.GetCardByID(cardID)
	if err != nil {
		if err.Error() == "card not found" {
			return ctx.JSON(http.StatusNotFound, echo.Map{"message": "Card not found"})
		}
		c.logger.Error("Failed to retrieve card", "error", err)
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"message": "Failed to retrieve card"})
	}
	c.logger.Info("bbb")

	var req LLMRequest
	if err := ctx.Bind(&req); err != nil {
		c.logger.Error("Failed to bind request", "error", err)
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid request format"})
	}

	// Format the prompt to include card context
	formattedPrompt := fmt.Sprintf(`This is regarding the following flashcard:

### Front
%s

### Back
%s

Question: %s`, card.Front.Text, card.Back.Text, req.Prompt)

	// Call the LLM service
	response, err := c.service.GetExplanation(formattedPrompt)
	if err != nil {
		c.logger.Error("Failed to get LLM explanation", "error", err)
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"message": "Failed to generate explanation"})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"explanation": response,
	})
}
