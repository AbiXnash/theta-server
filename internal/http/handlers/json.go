package handlers

import (
	"log/slog"
	"net/http"

	"github.com/AbiXnash/theta-server/internal/data"
	"github.com/labstack/echo/v5"
)

// JSONHandler serves static JSON files through HTTP handlers.
type JSONHandler struct {
	loader data.Loader
}

// NewJSONHandler creates a new JSON file handler.
func NewJSONHandler(loader data.Loader) JSONHandler {
	return JSONHandler{loader: loader}
}

// ServeFile returns a handler that serves a specific JSON file as response body.
func (h JSONHandler) ServeFile(fileName string) echo.HandlerFunc {
	return func(c *echo.Context) error {
		payload, err := h.loader.ReadJSON(fileName)
		if err != nil {
			slog.Error("failed to load json", "file", fileName, "error", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "server failed to send data",
			})
		}

		return c.JSONBlob(http.StatusOK, payload)
	}
}
