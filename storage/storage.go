package storage

import (
	"database/sql"
	"task/storage/postgres"

	"go.mongodb.org/mongo-driver/mongo"
)

type Storage interface {
	UserService() postgres.UserRepo
}

type storageImpl struct {
	Postgres *sql.DB
	Mongo    *mongo.Database
}

func NewStorage(db *sql.DB, mdb *mongo.Database) Storage {
	return &storageImpl{
		Postgres: db,
		Mongo:    mdb,
	}
}

func (S *storageImpl) UserService() postgres.UserRepo {
	return postgres.NewUserRepo(S.Postgres)
}
