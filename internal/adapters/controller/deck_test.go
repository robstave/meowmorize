// internal/adapters/controller/deck_test.go

package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/robstave/meowmorize/internal/domain/mocks"
	"github.com/robstave/meowmorize/internal/domain/types"
	"github.com/robstave/meowmorize/internal/logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// TestGetAllDecks tests the GetAllDecks controller method
func TestGetAllDecks(t *testing.T) {

	// Your JSON string
	jsonStr := `{
    	"id": "123e4567-e89b-12d3-a456-426614174000",
    	"name": "Capitals",
    		"description": "A deck of world capitals.",
   		 "cards": [
        	{
            "id": "e0c32c1c-b36f-4e10-9f47-b8e88c8ff383",
            "deck_id": "123e4567-e89b-12d3-a456-426614174000",
            "front": {
                "text": "Capital of France"
           	 },
            "back": {
                "text": "Paris"
           	 }
        	}
    	]
    }`

	var deck types.Deck
	err := json.Unmarshal([]byte(jsonStr), &deck)
	if err != nil {
		// Handle error
		panic(err)
	}
	// Create a slice with the parsed deck
	decks := []types.Deck{deck}

	// Initialize Echo
	e := echo.New()

	// Create a mock BLL service
	mockService := new(mocks.MeowDomain)

	// Setup expectations
	mockService.On("GetAllDecks").Return(decks, nil)

	// Initialize HomeController with mock service
	homeController := NewMeowController(mockService, logger.GetLogger())

	// Create a new HTTP GET request
	req := httptest.NewRequest(http.MethodGet, "/api/decks", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Invoke the controller method
	if assert.NoError(t, homeController.GetAllDecks(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		// Decode the response
		var response []types.Deck
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, decks, response)
		assert.Equal(t, 1, len(response))
	}

	// Assert that the expectations were met
	mockService.AssertExpectations(t)
}

// Helper function to create a multipart form file
func createMultipartFormFile(fieldName, fileName string, data []byte) (*bytes.Buffer, string, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(fieldName, fileName)
	if err != nil {
		return nil, "", err
	}
	_, err = part.Write(data)
	if err != nil {
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return body, writer.FormDataContentType(), nil
}

func getDeck() types.Deck {

	// Your JSON string
	jsonStr := `{
    	"id": "123e4567-e89b-12d3-a456-426614174000",
		"deckId": "123e4567-e89b-12d3-a456-426614174000",
    	"name": "Capitals",
    		"description": "A deck of world capitals.",
   		 "cards": [
        	{
            "id": "e0c32c1c-b36f-4e10-9f47-b8e88c8ff383",
            "deck_id": "123e4567-e89b-12d3-a456-426614174000",
            "front": {
                "text": "Capital of France"
           	 },
            "back": {
                "text": "Paris"
           	 }
        	}
    	]
    }`

	var deck types.Deck
	err := json.Unmarshal([]byte(jsonStr), &deck)
	if err != nil {
		// Handle error
		panic(err)
	}

	return deck
}

func TestGetDeck(t *testing.T) {
	// Initialize Echo
	getDeck()

}

// TestImportDeck tests the ImportDeck controller method
func TestImportDeck(t *testing.T) {
	// Initialize Echo
	e := echo.New()

	// Create a mock BLL service
	mockService := new(mocks.MeowDomain)

	// Create sample deck data
	deck := getDeck()

	// Setup expectations
	mockService.On("CreateDeck", deck).Return(nil)

	// Initialize HomeController with mock service
	homeController := NewMeowController(mockService, logger.GetLogger())

	// Prepare JSON data
	jsonData := map[string]types.Deck{
		"deck": deck,
	}
	jsonBytes, err := json.Marshal(jsonData)
	assert.NoError(t, err)

	// Create a multipart form file
	body, contentType, err := createMultipartFormFile("deck_file", "deck.json", jsonBytes)
	assert.NoError(t, err)

	// Create a new HTTP POST request
	req := httptest.NewRequest(http.MethodPost, "/api/decks/import", body)
	req.Header.Set("Content-Type", contentType)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Invoke the controller method
	if assert.NoError(t, homeController.ImportDeck(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)

		// Decode the response
		var response types.Deck
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, deck, response)
	}

	// Assert that the expectations were met
	mockService.AssertExpectations(t)
}

// TestImportDeck_InvalidJSON tests the ImportDeck controller with invalid JSON
func TestImportDeck_InvalidJSON(t *testing.T) {
	// Initialize Echo
	e := echo.New()

	// Create a mock BLL service
	mockService := new(mocks.MeowDomain)

	// Initialize HomeController with mock service
	homeController := NewMeowController(mockService, logger.GetLogger())

	// Prepare invalid JSON data
	invalidJSON := `{"deck": {"id": "123e4567-e89b-12d3-a456-426614174000", "name": "Capitals", "description": "A deck of world capitals.", "cards": [`

	// Create a multipart form file
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("deck_file", "deck.json")
	assert.NoError(t, err)
	_, err = part.Write([]byte(invalidJSON))
	assert.NoError(t, err)
	err = writer.Close()
	assert.NoError(t, err)

	// Create a new HTTP POST request
	req := httptest.NewRequest(http.MethodPost, "/api/decks/import", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Invoke the controller method
	if assert.NoError(t, homeController.ImportDeck(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)

		// Decode the response
		var response map[string]string
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "Invalid JSON format", response["message"])
	}

	// Assert that the service was not called
	mockService.AssertNotCalled(t, "CreateDeck", mock.Anything)
}

// TestImportDeck_ServiceError tests the ImportDeck controller when service.CreateDeck fails
func TestImportDeck_ServiceError(t *testing.T) {
	// Initialize Echo
	e := echo.New()

	// Create a mock BLL service
	mockService := new(mocks.MeowDomain)

	// Create sample deck data
	deck := getDeck()

	// Setup expectations
	mockService.On("CreateDeck", deck).Return(errors.New("database error"))

	// Initialize HomeController with mock service
	homeController := NewMeowController(mockService, logger.GetLogger())

	// Prepare JSON data
	jsonData := map[string]types.Deck{
		"deck": deck,
	}
	jsonBytes, err := json.Marshal(jsonData)
	assert.NoError(t, err)

	// Create a multipart form file
	body, contentType, err := createMultipartFormFile("deck_file", "deck.json", jsonBytes)
	assert.NoError(t, err)

	// Create a new HTTP POST request
	req := httptest.NewRequest(http.MethodPost, "/api/decks/import", body)
	req.Header.Set("Content-Type", contentType)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Invoke the controller method
	if assert.NoError(t, homeController.ImportDeck(c)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)

		// Decode the response
		var response map[string]string
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "Failed to save deck", response["message"])
	}

	// Assert that the expectations were met
	mockService.AssertExpectations(t)
}

// TestImportDeck_DuplicateID tests the ImportDeck controller when a deck with the same ID already exists
func TestImportDeck_DuplicateID(t *testing.T) {
	// Initialize Echo
	e := echo.New()

	// Create a mock BLL service
	mockService := new(mocks.MeowDomain)

	// Create sample deck data
	deck := getDeck()

	// Setup expectations
	mockService.On("CreateDeck", deck).Return(errors.New("deck ID already exists"))

	// Initialize HomeController with mock service
	homeController := NewMeowController(mockService, logger.GetLogger())

	// Prepare JSON data
	jsonData := map[string]types.Deck{
		"deck": deck,
	}
	jsonBytes, err := json.Marshal(jsonData)
	assert.NoError(t, err)

	// Create a multipart form file
	body, contentType, err := createMultipartFormFile("deck_file", "deck.json", jsonBytes)
	assert.NoError(t, err)

	// Create a new HTTP POST request
	req := httptest.NewRequest(http.MethodPost, "/api/decks/import", body)
	req.Header.Set("Content-Type", contentType)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Invoke the controller method
	if assert.NoError(t, homeController.ImportDeck(c)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)

		// Decode the response
		var response map[string]string
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "Failed to save deck", response["message"])
	}

	// Assert that the expectations were met
	mockService.AssertExpectations(t)
}

func TestCreateDeck_Success(t *testing.T) {
	e := echo.New()
	mockService := new(mocks.MeowDomain)

	deck := getDeck()
	mockService.On("CreateDeck", deck).Return(nil)

	controller := NewMeowController(mockService, logger.GetLogger())

	deckJSON, err := json.Marshal(deck)
	assert.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/api/decks", bytes.NewReader(deckJSON))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, controller.CreateDeck(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)

		var response types.Deck
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, deck, response)
	}

	mockService.AssertExpectations(t)
}

func TestCreateDeck_InvalidData(t *testing.T) {
	e := echo.New()
	mockService := new(mocks.MeowDomain)

	controller := NewMeowController(mockService, logger.GetLogger())

	// Missing required fields (e.g., Name)
	invalidDeck := types.Deck{
		ID:    "123e4567-e89b-12d3-a456-426614174000",
		Cards: []types.Card{},
	}
	deckJSON, err := json.Marshal(invalidDeck)
	assert.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/api/decks", bytes.NewReader(deckJSON))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Depending on your validation logic, you might need to adjust the expected response
	if assert.NoError(t, controller.CreateDeck(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)

		var response map[string]string
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "Invalid deck data", response["message"])
	}

	mockService.AssertNotCalled(t, "CreateDeck", mock.Anything)
}

func TestCreateDeck_ServiceError(t *testing.T) {
	e := echo.New()
	mockService := new(mocks.MeowDomain)

	deck := getDeck()
	mockService.On("CreateDeck", deck).Return(errors.New("database error"))

	controller := NewMeowController(mockService, logger.GetLogger())

	deckJSON, err := json.Marshal(deck)
	assert.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/api/decks", bytes.NewReader(deckJSON))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, controller.CreateDeck(c)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)

		var response map[string]string
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "Failed to create deck", response["message"])
	}

	mockService.AssertExpectations(t)
}

func TestCreateDefaultDeck_Success(t *testing.T) {
	e := echo.New()
	mockService := new(mocks.MeowDomain)

	mockService.On("CreateDefaultDeck", true).Return(nil)

	controller := NewMeowController(mockService, logger.GetLogger())

	requestBody := `{
		"defaultData": true
	}`
	req := httptest.NewRequest(http.MethodPost, "/api/decks", strings.NewReader(requestBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, controller.CreateDefaultDeck(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)

		var response map[string]string
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "Default deck created successfully", response["message"])
	}

	mockService.AssertExpectations(t)
}

func TestGetDeckByID_Success(t *testing.T) {
	e := echo.New()
	mockService := new(mocks.MeowDomain)

	deck := getDeck()
	mockService.On("GetDeckByID", deck.ID).Return(deck, nil)

	controller := NewMeowController(mockService, logger.GetLogger())

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/decks/%s", deck.ID), nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(deck.ID)

	if assert.NoError(t, controller.GetDeckByID(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var response types.Deck
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, deck, response)
	}

	mockService.AssertExpectations(t)
}

func TestGetDeckByID_NotFound(t *testing.T) {
	e := echo.New()
	mockService := new(mocks.MeowDomain)

	deckID := "non-existent-id"
	mockService.On("GetDeckByID", deckID).Return(types.Deck{}, errors.New("deck not found"))

	controller := NewMeowController(mockService, logger.GetLogger())

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/decks/%s", deckID), nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(deckID)

	if assert.NoError(t, controller.GetDeckByID(c)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)

		var response map[string]string
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "Failed to retrieve deck", response["message"])
	}

	mockService.AssertExpectations(t)
}

func TestUpdateDeck_Success(t *testing.T) {
	e := echo.New()
	mockService := new(mocks.MeowDomain)

	deck := getDeck()
	updatedDeck := deck
	updatedDeck.Name = "Updated Capitals"
	updatedDeck.Description = "Updated Description"

	mockService.On("UpdateDeck", updatedDeck).Return(nil)
	mockService.On("GetDeckByID", deck.ID).Return(deck, nil)

	controller := NewMeowController(mockService, logger.GetLogger())

	deckJSON, err := json.Marshal(updatedDeck)
	assert.NoError(t, err)

	req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/api/decks/%s", deck.ID), bytes.NewReader(deckJSON))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(deck.ID)

	if assert.NoError(t, controller.UpdateDeck(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var response types.Deck
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, updatedDeck, response)
	}

	mockService.AssertExpectations(t)
}

func TestDeleteDeck_Success(t *testing.T) {
	e := echo.New()
	mockService := new(mocks.MeowDomain)

	deckID := "123e4567-e89b-12d3-a456-426614174000"
	mockService.On("DeleteDeck", deckID).Return(nil)

	controller := NewMeowController(mockService, logger.GetLogger())

	req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/api/decks/%s", deckID), nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(deckID)

	if assert.NoError(t, controller.DeleteDeck(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var response map[string]string
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "Deck deleted successfully", response["message"])
	}

	mockService.AssertExpectations(t)
}
func TestDeleteDeck_NotFound(t *testing.T) {
	e := echo.New()
	mockService := new(mocks.MeowDomain)

	deckID := "non-existent-id"
	mockService.On("DeleteDeck", deckID).Return(errors.New("deck not found"))

	controller := NewMeowController(mockService, logger.GetLogger())

	req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/api/decks/%s", deckID), nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(deckID)

	if assert.NoError(t, controller.DeleteDeck(c)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)

		var response map[string]string
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "Failed to delete deck", response["message"])
	}

	mockService.AssertExpectations(t)
}

func TestDeleteDeck_InvalidID(t *testing.T) {
	e := echo.New()
	mockService := new(mocks.MeowDomain)

	controller := NewMeowController(mockService, logger.GetLogger())

	req := httptest.NewRequest(http.MethodDelete, "/api/decks/", nil) // Missing ID
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("") // Empty ID

	if assert.NoError(t, controller.DeleteDeck(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)

		var response map[string]string
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "Deck ID is required", response["message"])
	}

	mockService.AssertExpectations(t)
}

func TestExportDeck_Success(t *testing.T) {
	e := echo.New()
	mockService := new(mocks.MeowDomain)

	deck := getDeck()
	mockService.On("ExportDeck", deck.ID).Return(deck, nil)

	controller := NewMeowController(mockService, logger.GetLogger())

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/decks/export/%s", deck.ID), nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(deck.ID)

	if assert.NoError(t, controller.ExportDeck(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "application/json", rec.Header().Get("Content-Type"))
		expectedFilename := fmt.Sprintf("deck-%s.json", deck.ID)
		assert.Equal(t, fmt.Sprintf("attachment; filename=\"%s\"", expectedFilename), rec.Header().Get("Content-Disposition"))

		var response types.Deck
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, deck, response)
	}

	mockService.AssertExpectations(t)
}
func TestExportDeck_NotFound(t *testing.T) {
	e := echo.New()
	mockService := new(mocks.MeowDomain)

	deckID := "non-existent-id"
	mockService.On("ExportDeck", deckID).Return(types.Deck{}, errors.New("deck not found"))

	controller := NewMeowController(mockService, logger.GetLogger())

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/decks/export/%s", deckID), nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(deckID)

	if assert.NoError(t, controller.ExportDeck(c)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)

		var response map[string]string
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "Failed to export deck", response["message"])
	}

	mockService.AssertExpectations(t)
}
func TestExportDeck_InvalidID(t *testing.T) {
	e := echo.New()
	mockService := new(mocks.MeowDomain)

	controller := NewMeowController(mockService, logger.GetLogger())

	req := httptest.NewRequest(http.MethodGet, "/api/decks/export/", nil) // Missing ID
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("") // Empty ID

	if assert.NoError(t, controller.ExportDeck(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)

	}

	mockService.AssertNotCalled(t, "ExportDeck", mock.Anything)
}
