package handler

import (
	"github.com/xtt28/freakbot/internal/classifier"
	"github.com/xtt28/freakbot/internal/repository"
)

type Handler struct {
	dbConn repository.Connection
	classifierService classifier.ClassifierService
}

func NewHandler(dbConn repository.Connection, classifierService classifier.ClassifierService) *Handler {
	return &Handler{dbConn, classifierService}
}