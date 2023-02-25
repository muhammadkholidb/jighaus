package handler

import "github.com/labstack/echo/v4"

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func RegisterRoutes(e *echo.Echo, h *Handler) {
	// API
	e.POST("/v1/upload", h.Upload)

	// HTML
	e.File("/", "public/index.html")
}
