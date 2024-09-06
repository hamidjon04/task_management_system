package handler

import (
	"log/slog"
	"task/storage"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
}

type hanlerimpl struct {
	Storage storage.Storage
	Logger  *slog.Logger
}

func NewHandler(storage storage.Storage, logger *slog.Logger) Handler {
	return &hanlerimpl{
		Storage: storage,
		Logger:  logger,
	}
}
