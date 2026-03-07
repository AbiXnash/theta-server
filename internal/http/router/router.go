package router

import (
	"github.com/AbiXnash/theta-server/internal/http/handlers"
	"github.com/labstack/echo/v5"
)

var jsonEndpoints = map[string]string{
	"team":     "team.json",
	"config":   "config.json",
	"events":   "events.json",
	"content":  "content.json",
	"sponsors": "sponsors.json",
	"manifest": "manifest.json",
}

// Register wires all JSON-backed API endpoints on the provided group.
func Register(api *echo.Group, h handlers.JSONHandler) {
	for route, file := range jsonEndpoints {
		api.GET("/"+route, h.ServeFile(file))
	}
}
