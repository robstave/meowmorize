// internal/adapters/controller/auth.go
package controller

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

// JWTSecret is the secret key used to sign JWT tokens.
// In a production environment, ensure this is secure and not hard-coded.
var JWTSecret = []byte("your_secret_key")

// LoginRequest represents the expected payload for login
type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// LoginResponse represents the response containing the JWT token
type LoginResponse struct {
	Token string `json:"token"`
}

// Login handles user authentication
// @Summary User Login
// @Description Authenticate user and return a JWT token
// @Tags Authentication
// @Accept  json
// @Produce  json
// @Param credentials body LoginRequest true "User Credentials"
// @Success 200 {object} LoginResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /login [post]
func (hc *MeowController) Login(c echo.Context) error {
	var req LoginRequest
	if err := c.Bind(&req); err != nil {
		hc.logger.Error("Failed to bind login request", "error", err)
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid request payload",
		})
	}

	user, err := hc.service.GetUserByUsername(req.Username)
	if err != nil {
		hc.logger.Error("Error fetching user", "username", req.Username, "error", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Internal server error",
		})
	}

	if user == nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": "Invalid username or password",
		})
	}

	// Compare passwords
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": "Invalid username or password",
		})
	}

	// Create JWT token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["username"] = user.Username
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix() // Token expires after 72 hours

	t, err := token.SignedString(JWTSecret)
	if err != nil {
		hc.logger.Error("Failed to sign JWT token", "error", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to generate token",
		})
	}

	return c.JSON(http.StatusOK, LoginResponse{
		Token: t,
	})
}
