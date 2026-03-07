package middleware

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

func APIKey(requiredKey string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			provided := c.Request().Header.Get("X-ABX-THETA-KEY")
			if provided != requiredKey {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"error": "authorization failed",
				})
			}
			return next(c)
		}
	}
}
