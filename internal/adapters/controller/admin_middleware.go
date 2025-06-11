package controller

import (
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

// AdminMiddleware ensures the authenticated user has role 'admin'.
func AdminMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user")
		if user == nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "missing token")
		}

		claims, ok := user.(*jwt.Token).Claims.(jwt.MapClaims)
		if !ok {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
		}

		role, ok := claims["role"].(string)
		if !ok || role != "admin" {
			return echo.NewHTTPError(http.StatusForbidden, "admin access required")
		}

		return next(c)
	}
}
