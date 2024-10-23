package controller

import (
	"log/slog"

	"github.com/robstave/meowmorize/internal/domain"
)

type HomeController struct {
	service domain.BLL
	logger  *slog.Logger
}

func NewHomeController(service domain.BLL, logger *slog.Logger) *HomeController {
	return &HomeController{service: service, logger: logger}
}
