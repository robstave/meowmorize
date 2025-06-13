package controller

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"

	"github.com/robstave/meowmorize/internal/domain/types"
)

// AdminGetAllUsers returns all users
// @Summary Get all users
// @Description Retrieve all registered users (admin only)
// @Tags Users
// @Produce json
// @Security BearerAuth
// @Success 200 {array} types.User
// @Failure 403 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /admin/users [get]
func (hc *MeowController) AdminGetAllUsers(c echo.Context) error {
	users, err := hc.service.GetAllUsers()
	if err != nil {
		hc.logger.Error("failed to get users", "error", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "failed to retrieve users"})
	}
	return c.JSON(http.StatusOK, users)
}

// AdminCreateUserRequest represents the payload for creating a user
type AdminCreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

// AdminCreateUser creates a new user
// @Summary Create user
// @Description Create a new user (admin only)
// @Tags Users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param user body AdminCreateUserRequest true "User to create"
// @Success 201 {object} types.User
// @Failure 400 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /admin/users [post]
func (hc *MeowController) AdminCreateUser(c echo.Context) error {
	var req AdminCreateUserRequest
	if err := c.Bind(&req); err != nil {
		hc.logger.Error("failed to bind user", "error", err)
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "invalid user payload"})
	}

	if req.Role == "" {
		req.Role = "user"
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		hc.logger.Error("failed to hash password", "error", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "failed to create user"})
	}

	user := types.User{
		ID:       uuid.New().String(),
		Username: req.Username,
		Password: string(hashedPassword),
		Role:     req.Role,
	}

	if err := hc.service.CreateUser(user); err != nil {
		hc.logger.Error("failed to create user", "error", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "failed to create user"})
	}
	return c.JSON(http.StatusCreated, user)
}

// AdminDeleteUser removes a user
// @Summary Delete user
// @Description Delete a user by ID (admin only)
// @Tags Users
// @Produce json
// @Param id path string true "User ID"
// @Security BearerAuth
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /admin/users/{id} [delete]
func (hc *MeowController) AdminDeleteUser(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "id required"})
	}

	if err := hc.service.DeleteUser(id); err != nil {
		hc.logger.Error("failed to delete user", "error", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "failed to delete user"})
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "user deleted"})
}

// ChangePasswordRequest represents the payload for updating a user's password
type ChangePasswordRequest struct {
	Password string `json:"password"`
}

// ChangePassword allows a logged in user to change their own password
// @Summary Change user password
// @Description Update the authenticated user's password
// @Tags Users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param password body ChangePasswordRequest true "New password"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /user/password [put]
func (hc *MeowController) ChangePassword(c echo.Context) error {
	var req ChangePasswordRequest
	if err := c.Bind(&req); err != nil {
		hc.logger.Error("failed to bind password request", "error", err)
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "invalid request payload"})
	}

	userID, err := getUserIDFromContext(c)
	if err != nil {
		hc.logger.Error("failed to get user from token", "error", err)
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": "unauthorized"})
	}

	if err := hc.service.UpdateUserPassword(userID, req.Password); err != nil {
		hc.logger.Error("failed to update password", "error", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "failed to update password"})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "password updated"})
}
