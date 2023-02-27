package handler

import (
	"embed"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	public embed.FS
}

func NewHandler(public embed.FS) *Handler {
	return &Handler{
		public: public,
	}
}

func RegisterRoutes(e *echo.Echo, h *Handler) {
	// API
	e.POST("/v1/upload", h.Upload)

	// HTML
	fs := echo.MustSubFS(h.public, "public")
	e.StaticFS("/", fs)
}
