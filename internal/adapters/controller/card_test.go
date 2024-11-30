// internal/adapters/controller/card_test.go

package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/robstave/meowmorize/internal/domain/mocks"
	"github.com/robstave/meowmorize/internal/domain/types"
	"github.com/robstave/meowmorize/internal/logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Helper function to get a sample card
func getCard() types.Card {
	return types.Card{
		ID:     "e0c32c1c-b36f-4e10-9f47-b8e88c8ff383",
		DeckID: "123e4567-e89b-12d3-a456-426614174000",
		Front: types.CardFront{
			Text: "Capital of France",
		},
		Back: types.CardBack{
			Text: "Paris",
		},
		Link: "https://example.com/paris",
	}
}

// TestGetCardByID_Success tests the successful retrieval of a card by ID
func TestGetCardByID_Success(t *testing.T) {
	e := echo.New()
	mockService := new(mocks.MeowDomain)

	card := getCard()
	mockService.On("GetCardByID", card.ID).Return(&card, nil)

	controller := NewMeowController(mockService, logger.GetLogger())

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/cards/%s", card.ID), nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(card.ID)

	// Invoke the controller method
	if assert.NoError(t, controller.GetCardByID(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var response types.Card
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, card, response)
	}

	mockService.AssertExpectations(t)
}

// TestGetCardByID_NotFound tests the retrieval of a non-existent card
func TestGetCardByID_NotFound(t *testing.T) {
	e := echo.New()
	mockService := new(mocks.MeowDomain)

	cardID := "non-existent-id"
	mockService.On("GetCardByID", cardID).Return((*types.Card)(nil), errors.New("card not found"))

	controller := NewMeowController(mockService, logger.GetLogger())

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/cards/%s", cardID), nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(cardID)

	// Invoke the controller method
	err := controller.GetCardByID(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, rec.Code)

	var response map[string]string
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Card not found", response["message"])

	mockService.AssertExpectations(t)
}

// TestGetCardByID_ServiceError tests handling of service layer errors when retrieving a card
func TestGetCardByID_ServiceError(t *testing.T) {
	e := echo.New()
	mockService := new(mocks.MeowDomain)

	cardID := "123e4567-e89b-12d3-a456-426614174000"
	mockService.On("GetCardByID", cardID).Return((*types.Card)(nil), errors.New("database error"))

	controller := NewMeowController(mockService, logger.GetLogger())

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/cards/%s", cardID), nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(cardID)

	// Invoke the controller method
	err := controller.GetCardByID(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)

	var response map[string]string
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Failed to retrieve card", response["message"])

	mockService.AssertExpectations(t)
}

// TestCreateCard_Success tests the successful creation of a card
func TestCreateCard_Success(t *testing.T) {
	e := echo.New()
	mockService := new(mocks.MeowDomain)

	reqPayload := CreateCardRequest{
		DeckID: "123e4567-e89b-12d3-a456-426614174000",
		Front: CardContentReq{
			Text: "Capital of Spain",
		},
		Back: CardContentReq{
			Text: "Madrid",
		},
		Link: "https://example.com/madrid",
	}

	createdCard := types.Card{
		ID:     "f1c32c1c-b36f-4e10-9f47-b8e88c8ff384",
		DeckID: reqPayload.DeckID,
		Front: types.CardFront{
			Text: reqPayload.Front.Text,
		},
		Back: types.CardBack{
			Text: reqPayload.Back.Text,
		},
		Link: reqPayload.Link,
	}

	mockService.On("CreateCard", mock.MatchedBy(func(card types.Card) bool {
		return card.DeckID == reqPayload.DeckID &&
			card.Front.Text == reqPayload.Front.Text &&
			card.Back.Text == reqPayload.Back.Text &&
			card.Link == reqPayload.Link
	})).Return(&createdCard, nil)

	controller := NewMeowController(mockService, logger.GetLogger())

	reqBody, err := json.Marshal(reqPayload)
	assert.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/api/cards", bytes.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Invoke the controller method
	if assert.NoError(t, controller.CreateCard(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)

		var response types.Card
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, createdCard, response)
	}

	mockService.AssertExpectations(t)
}

// TestCreateCard_InvalidPayload tests creating a card with invalid payload
/*
func TestCreateCard_InvalidPayload(t *testing.T) {
	e := echo.New()
	mockService := new(mocks.MeowDomain)

	controller := NewMeowController(mockService, logger.GetLogger())

	// Invalid JSON (missing closing brace)
	invalidJSON := `{"deck_id": "", "front": {"text": ""}, "back": {"text": ""}, "link": ""`

	req := httptest.NewRequest(http.MethodPost, "/api/cards", bytes.NewReader([]byte(invalidJSON)))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Invoke the controller method
	err := controller.CreateCard(c)
	assert.Error(t, err)
	//assert.Equal(t, http.StatusBadRequest, rec.Code)

	var response map[string]string
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Invalid request payload", response["message"])

	mockService.AssertNotCalled(t, "CreateCard", mock.Anything)
}
*/

// TestCreateCard_ServiceError tests handling of service layer errors when creating a card
func TestCreateCard_ServiceError(t *testing.T) {
	e := echo.New()
	mockService := new(mocks.MeowDomain)

	reqPayload := CreateCardRequest{
		DeckID: "123e4567-e89b-12d3-a456-426614174000",
		Front: CardContentReq{
			Text: "Capital of Germany",
		},
		Back: CardContentReq{
			Text: "Berlin",
		},
		Link: "https://example.com/berlin",
	}

	mockService.On("CreateCard", mock.MatchedBy(func(card types.Card) bool {
		return card.DeckID == reqPayload.DeckID &&
			card.Front.Text == reqPayload.Front.Text &&
			card.Back.Text == reqPayload.Back.Text &&
			card.Link == reqPayload.Link
	})).Return((*types.Card)(nil), errors.New("database error"))

	controller := NewMeowController(mockService, logger.GetLogger())

	reqBody, err := json.Marshal(reqPayload)
	assert.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/api/cards", bytes.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Invoke the controller method
	err = controller.CreateCard(c)
	if assert.Error(t, err) {
		// Define a custom error handler to process the error
		e.HTTPErrorHandler = func(err error, c echo.Context) {
			if he, ok := err.(*echo.HTTPError); ok {
				var msg string
				switch m := he.Message.(type) {
				case string:
					msg = m
				case map[string]interface{}:
					c.JSON(he.Code, m)
					return
				default:
					msg = "Internal Server Error"
				}
				c.JSON(he.Code, echo.Map{
					"message": msg,
				})
			} else {
				c.JSON(http.StatusInternalServerError, echo.Map{
					"message": "Internal Server Error",
				})
			}
		}

		// Let Echo's error handler process the error
		e.HTTPErrorHandler(err, c)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)

		var response map[string]string
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "Failed to create card", response["message"])
	}

	mockService.AssertExpectations(t)
}

// TestUpdateCard_Success tests the successful update of a card
func TestUpdateCard_Success(t *testing.T) {
	e := echo.New()
	mockService := new(mocks.MeowDomain)

	card := getCard()
	updatedCard := card
	updatedCard.Front.Text = "Updated Front Text"
	updatedCard.Back.Text = "Updated Back Text"
	updatedCard.Link = "https://example.com/updated-link"

	mockService.On("GetCardByID", card.ID).Return(&card, nil)
	mockService.On("UpdateCard", updatedCard).Return(nil)

	controller := NewMeowController(mockService, logger.GetLogger())

	updatePayload := UpdateCardRequest{
		Front: &CardContentReq{
			Text: "Updated Front Text",
		},
		Back: &CardContentReq{
			Text: "Updated Back Text",
		},
		Link: &updatedCard.Link,
	}

	reqBody, err := json.Marshal(updatePayload)
	assert.NoError(t, err)

	req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/api/cards/%s", card.ID), bytes.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(card.ID)

	// Invoke the controller method
	if assert.NoError(t, controller.UpdateCard(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var response types.Card
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, updatedCard, response)
	}

	mockService.AssertExpectations(t)
}

// TestUpdateCard_InvalidPayload tests updating a card with invalid payload
func TestUpdateCard_InvalidPayload(t *testing.T) {
	e := echo.New()
	mockService := new(mocks.MeowDomain)

	controller := NewMeowController(mockService, logger.GetLogger())

	// Invalid JSON
	invalidJSON := `{"front": {"text": ""}, "back": {"text": ""}, "link": ""`

	req := httptest.NewRequest(http.MethodPut, "/api/cards/123e4567-e89b-12d3-a456-426614174000", bytes.NewReader([]byte(invalidJSON)))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("123e4567-e89b-12d3-a456-426614174000")

	// Invoke the controller method
	err := controller.UpdateCard(c)
	if assert.Error(t, err) {
		// Define a custom error handler to process the error
		e.HTTPErrorHandler = func(err error, c echo.Context) {
			if he, ok := err.(*echo.HTTPError); ok {
				var msg string
				switch m := he.Message.(type) {
				case string:
					msg = m
				case map[string]interface{}:
					c.JSON(he.Code, m)
					return
				default:
					msg = "Internal Server Error"
				}
				c.JSON(he.Code, echo.Map{
					"message": msg,
				})
			} else {
				c.JSON(http.StatusInternalServerError, echo.Map{
					"message": "Internal Server Error",
				})
			}
		}

		// Let Echo's error handler process the error
		e.HTTPErrorHandler(err, c)

		assert.Equal(t, http.StatusBadRequest, rec.Code)

		var response map[string]string
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "Invalid request payload", response["message"])
	}

	mockService.AssertNotCalled(t, "UpdateCard", mock.Anything)
}

// TestUpdateCard_NotFound tests updating a non-existent card
func TestUpdateCard_NotFound(t *testing.T) {
	e := echo.New()
	mockService := new(mocks.MeowDomain)

	cardID := "non-existent-id"
	mockService.On("GetCardByID", cardID).Return((*types.Card)(nil), errors.New("card not found"))

	controller := NewMeowController(mockService, logger.GetLogger())

	updatePayload := UpdateCardRequest{
		Front: &CardContentReq{
			Text: "Updated Front Text",
		},
		Back: &CardContentReq{
			Text: "Updated Back Text",
		},
		Link: nil,
	}

	reqBody, err := json.Marshal(updatePayload)
	assert.NoError(t, err)

	req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/api/cards/%s", cardID), bytes.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(cardID)

	// Define a custom error handler to process the error
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		if he, ok := err.(*echo.HTTPError); ok {
			var msg string
			switch m := he.Message.(type) {
			case string:
				msg = m
			case map[string]interface{}:
				c.JSON(he.Code, m)
				return
			default:
				msg = "Internal Server Error"
			}
			c.JSON(he.Code, echo.Map{
				"message": msg,
			})
		} else {
			c.JSON(http.StatusInternalServerError, echo.Map{
				"message": "Internal Server Error",
			})
		}
	}

	// Invoke the controller method
	err = controller.UpdateCard(c)
	if assert.Error(t, err) {
		// Let Echo's error handler process the error
		e.HTTPErrorHandler(err, c)

		assert.Equal(t, http.StatusNotFound, rec.Code)

		var response map[string]string
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "Card not found", response["message"])
	}

	mockService.AssertExpectations(t)
}

// TestUpdateCard_ServiceError tests handling of service layer errors when updating a card
func TestUpdateCard_ServiceError(t *testing.T) {
	e := echo.New()
	mockService := new(mocks.MeowDomain)

	card := getCard()
	mockService.On("GetCardByID", card.ID).Return(&card, nil)

	updatedCard := card
	updatedCard.Front.Text = "Updated Front Text"

	mockService.On("UpdateCard", updatedCard).Return(errors.New("database error"))

	controller := NewMeowController(mockService, logger.GetLogger())

	updatePayload := UpdateCardRequest{
		Front: &CardContentReq{
			Text: "Updated Front Text",
		},
		Back: nil,
		Link: nil,
	}

	reqBody, err := json.Marshal(updatePayload)
	assert.NoError(t, err)

	req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/api/cards/%s", card.ID), bytes.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(card.ID)

	// Invoke the controller method
	err = controller.UpdateCard(c)
	if assert.Error(t, err) {
		// Define a custom error handler to process the error
		e.HTTPErrorHandler = func(err error, c echo.Context) {
			if he, ok := err.(*echo.HTTPError); ok {
				var msg string
				switch m := he.Message.(type) {
				case string:
					msg = m
				case map[string]interface{}:
					c.JSON(he.Code, m)
					return
				default:
					msg = "Internal Server Error"
				}
				c.JSON(he.Code, echo.Map{
					"message": msg,
				})
			} else {
				c.JSON(http.StatusInternalServerError, echo.Map{
					"message": "Internal Server Error",
				})
			}
		}

		// Let Echo's error handler process the error
		e.HTTPErrorHandler(err, c)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)

		var response map[string]string
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "Failed to update card", response["message"])
	}

	mockService.AssertExpectations(t)
}

// TestDeleteCard_Success tests the successful deletion of a card
func TestDeleteCard_Success(t *testing.T) {
	e := echo.New()
	mockService := new(mocks.MeowDomain)

	cardID := "e0c32c1c-b36f-4e10-9f47-b8e88c8ff383"
	mockService.On("DeleteCardByID", cardID).Return(nil)

	controller := NewMeowController(mockService, logger.GetLogger())

	req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/api/cards/%s", cardID), nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(cardID)

	// Invoke the controller method
	if assert.NoError(t, controller.DeleteCard(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var response map[string]string
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "Card deleted successfully", response["message"])
	}

	mockService.AssertExpectations(t)
}

// TestDeleteCard_NotFound tests deleting a non-existent card
func TestDeleteCard_NotFound(t *testing.T) {
	e := echo.New()
	mockService := new(mocks.MeowDomain)

	cardID := "non-existent-id"
	mockService.On("DeleteCardByID", cardID).Return(errors.New("card not found"))

	controller := NewMeowController(mockService, logger.GetLogger())

	req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/api/cards/%s", cardID), nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(cardID)

	// Invoke the controller method
	err := controller.DeleteCard(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, rec.Code)

	var response map[string]string
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Card not found", response["message"])

	mockService.AssertExpectations(t)
}

// TestDeleteCard_InvalidID tests deleting a card without providing an ID
func TestDeleteCard_InvalidID(t *testing.T) {
	e := echo.New()
	mockService := new(mocks.MeowDomain)

	controller := NewMeowController(mockService, logger.GetLogger())

	req := httptest.NewRequest(http.MethodDelete, "/api/cards/", nil) // Missing ID
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("") // Empty ID

	// Invoke the controller method
	err := controller.DeleteCard(c)
	if assert.Error(t, err) {
		// Define a custom error handler to process the error
		e.HTTPErrorHandler = func(err error, c echo.Context) {
			if he, ok := err.(*echo.HTTPError); ok {
				var msg string
				switch m := he.Message.(type) {
				case string:
					msg = m
				case map[string]interface{}:
					c.JSON(he.Code, m)
					return
				default:
					msg = "Internal Server Error"
				}
				c.JSON(he.Code, echo.Map{
					"message": msg,
				})
			} else {
				c.JSON(http.StatusInternalServerError, echo.Map{
					"message": "Internal Server Error",
				})
			}
		}

		// Let Echo's error handler process the error
		e.HTTPErrorHandler(err, c)

		assert.Equal(t, http.StatusBadRequest, rec.Code)

		var response map[string]string
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "Card ID is required", response["message"])
	}

	mockService.AssertNotCalled(t, "DeleteCardByID", mock.Anything)
}

// TestUpdateCard_NoFields tests updating a card without providing any fields
func TestUpdateCard_NoFields(t *testing.T) {
	e := echo.New()
	mockService := new(mocks.MeowDomain)

	card := getCard()
	mockService.On("GetCardByID", card.ID).Return(&card, nil)
	mockService.On("UpdateCard", mock.Anything).Return(nil) // No changes

	controller := NewMeowController(mockService, logger.GetLogger())

	// Empty update payload
	updatePayload := UpdateCardRequest{}

	reqBody, err := json.Marshal(updatePayload)
	assert.NoError(t, err)

	req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/api/cards/%s", card.ID), bytes.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(card.ID)

	// Invoke the controller method
	if assert.NoError(t, controller.UpdateCard(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var response types.Card
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, card, response)
	}

	mockService.AssertExpectations(t)
}
