package api

import (
	"log/slog"
	"task/api/handler"
	"task/config"
	"task/storage"

	"github.com/gin-gonic/gin"
)

type Connector interface{
	SetUpRoutes(storage storage.Storage, logger *slog.Logger)	
	StartRouter(cfg config.Config)error
}

type connectorImpl struct{
	Router *gin.Engine
}

func NewConnector(r *gin.Engine)Connector{
	return &connectorImpl{
		Router: r,
	}
}

func(R *connectorImpl) SetUpRoutes(storage storage.Storage, logger *slog.Logger){
	h := handler.NewHandler(storage, logger)
	auth := R.Router.Group("/auth")
	auth.POST("/register", h.Register)
	auth.POST("/login", h.Login)
	
}

func(R *connectorImpl) StartRouter(cfg config.Config)error{
	return R.Router.Run(cfg.APP_PORT)
}
