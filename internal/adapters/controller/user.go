package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
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

// AdminCreateUser creates a new user
// @Summary Create user
// @Description Create a new user (admin only)
// @Tags Users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param user body types.User true "User to create"
// @Success 201 {object} types.User
// @Failure 400 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /admin/users [post]
func (hc *MeowController) AdminCreateUser(c echo.Context) error {
	var user types.User
	if err := c.Bind(&user); err != nil {
		hc.logger.Error("failed to bind user", "error", err)
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "invalid user payload"})
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
