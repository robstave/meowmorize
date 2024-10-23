package controller

import (
	"log/slog"

	"github.com/robstave/meowmorize/internal/domain"
)

type MeowController struct {
	service domain.MeowDomain
	logger  *slog.Logger
}

func NewMeowController(service domain.MeowDomain, logger *slog.Logger) *MeowController {
	return &MeowController{service: service, logger: logger}
}
