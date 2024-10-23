// internal/adapters/controller/deck_test.go

package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"mime/multipart"
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
