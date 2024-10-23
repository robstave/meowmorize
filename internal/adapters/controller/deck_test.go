// internal/adapters/controller/deck_test.go

package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/robstave/meowmorize/internal/domain/mocks"
	"github.com/robstave/meowmorize/internal/domain/types"
	"github.com/stretchr/testify/assert"
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
            "deckId": "123e4567-e89b-12d3-a456-426614174000",
            "front": {
                "text": "Capital of France"
           	 },
            "back": {
                "text": "Paris"
           	 }
        	}
    	]
    }`

	// Initialize Echo
	e := echo.New()

	// Create a mock BLL service
	mockService := new(mocks.MeowDomain)

	var deck types.Deck
	err := json.Unmarshal([]byte(jsonStr), &deck)
	if err != nil {
		// Handle error
		panic(err)
	}
	// Create a slice with the parsed deck
	decks := []types.Deck{deck}
	// Setup expectations
	mockService.On("GetAllDecks").Return(decks, nil)

	// Initialize HomeController with mock service
	homeController := NewHomeController(mockService, nil)

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
