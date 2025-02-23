package controller

import (
	"encoding/json"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/robstave/meowmorize/internal/domain/types"
)

// setValidUserToken sets a valid JWT token (using a literal pointer) with the provided username in the context.
func setValidUserToken2(c echo.Context, username string) {
	token := &jwt.Token{
		Method: jwt.SigningMethodHS256,
		Claims: jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour).Unix(),
		},
	}
	c.Set("user", token)
}

// getDeck returns a sample deck decoded from JSON.
// Updated to include the "user_id" field and remove any deprecated fields.
func getDeck() types.Deck {
	jsonStr := `{
		"id": "123e4567-e89b-12d3-a456-426614174000",
		"name": "Capitals",
		"description": "A deck of world capitals.",
		"user_id": "testuser",
		"cards": [
			{
				"id": "e0c32c1c-b36f-4e10-9f47-b8e88c8ff383",
				"front": { "text": "Capital of France" },
				"back": { "text": "Paris" },
				"link": "https://example.com/paris"
			}
		]
	}`
	var deck types.Deck
	if err := json.Unmarshal([]byte(jsonStr), &deck); err != nil {
		panic(err)
	}
	return deck
}
