package main

import (
	"fmt"
	"log/slog"

	"github.com/AbiXnash/theta-server/internal/config"
	"github.com/AbiXnash/theta-server/internal/data"
	httpHandlers "github.com/AbiXnash/theta-server/internal/http/handlers"
	httpMiddleware "github.com/AbiXnash/theta-server/internal/http/middleware"
	httpRouter "github.com/AbiXnash/theta-server/internal/http/router"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func main() {
	cfg := config.Load()
	loader := data.NewLoader(cfg.DataDir)
	handler := httpHandlers.NewJSONHandler(loader)

	e := echo.New()
	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	e.GET("/health", func(c *echo.Context) error {
		return c.JSON(200, map[string]string{"status": "ok"})
	})

	api := e.Group("/api")
	api.Use(httpMiddleware.APIKey(cfg.APIKey))
	httpRouter.Register(api, handler)

	addr := fmt.Sprintf(":%s", cfg.Port)
	if err := e.Start(addr); err != nil {
		slog.Error("failed to start server", "address", addr, "error", err)
	}
}
