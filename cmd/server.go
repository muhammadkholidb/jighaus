package cmd

import (
	"embed"
	"jighaus/internal/config"
	"jighaus/internal/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func Server(public embed.FS) {
	err := config.Load()
	if err != nil {
		log.Fatal("Failed to load config: ", err)
	}
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	h := handler.NewHandler(public)
	handler.RegisterRoutes(e, h)
	err = e.Start(config.Data.Port)
	if err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
