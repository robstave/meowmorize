package controller

import (
	"log/slog"

	"github.com/robstave/meowmorize/internal/domain"
)

type HomeController struct {
	service domain.MeowDomain
	logger  *slog.Logger
}

func NewHomeController(service domain.MeowDomain, logger *slog.Logger) *HomeController {
	return &HomeController{service: service, logger: logger}
}
