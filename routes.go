package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v5"
)

func loadJSON(path string) ([]byte, error) {
	var err error
	json, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("cannot get the proper data for %s: %w", path, err)
	}

	return json, nil
}

func validateKey(providedKey string) bool {
	key := "somethings"

	if key == providedKey {
		return true
	}

	return false
}

func GetTeams(c *echo.Context) error {
	proviedKey := c.Request().Header.Get("X-ABX-THETA-KEY")
	if !validateKey(proviedKey) {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Authotization failed",
		})
	}

	json, err := loadJSON("data/teams.json")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Server failed to send data",
		})
	}
	return c.JSONBlob(http.StatusOK, json)
}
