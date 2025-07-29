package handler

import (
	"github.com/xtt28/freakbot/internal/classifier"
	"github.com/xtt28/freakbot/internal/commands"
	"github.com/xtt28/freakbot/internal/repository"
)

type Handler struct {
	dbConn            repository.Connection
	classifierService classifier.ClassifierService
	commandRegistry   *commands.CommandRegistry
}

func NewHandler(
	dbConn repository.Connection,
	classifierService classifier.ClassifierService,
	commandRegistry *commands.CommandRegistry,
) *Handler {
	return &Handler{dbConn, classifierService, commandRegistry}
}
