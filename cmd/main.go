package main

import (
	"fmt"
	"task/api"
	"task/config"
	"task/logs"
	"task/storage"
	"task/storage/postgres"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func main(){
	cfg := config.LoadConfig()
	logger := logs.InitLogger()

	db, err := postgres.ConnectDB(cfg)
	if err != nil{
		logger.Error(fmt.Sprintf("Postgresga ulanishda xatolik: %v", err))
		panic(err)
	}

	storage := storage.NewStorage(db, &mongo.Database{})

	connector := api.NewConnector(gin.Default())
	connector.SetUpRoutes(storage, logger)
	err = connector.StartRouter(cfg)
	if err != nil{
		logger.Error(fmt.Sprintf("Router run bo'lmadi: %v", err))
		panic(err)
	}
}