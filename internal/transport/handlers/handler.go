package handlers

import (
	"net/http"
	"urlshortener/internal/service"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	srv *service.Service
}

func NewHandler(srv *service.Service) *Handler {
	return &Handler{
		srv: srv,
	}
}

func (h *Handler) HealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
