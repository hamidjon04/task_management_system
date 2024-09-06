package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	DB_HOST     string
	DB_PORT     string
	DB_USER     string
	DB_NAME     string
	DB_PASSWORD string
	APP_PORT    string
	SIGNING_KEY string
}

func LoadConfig() Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Println(".env filedan ma'lumotlarni o'qib bo'lmadi!!!")
	}

	config := Config{}
	config.DB_HOST = cast.ToString(coalesce("DB_HOST", "postgres"))
	config.DB_PORT = cast.ToString(coalesce("DB_PORT", "5432"))
	config.DB_USER = cast.ToString(coalesce("DB_USER", "postgres"))
	config.DB_NAME = cast.ToString(coalesce("DB_NAME", "task_managment"))
	config.DB_PASSWORD = cast.ToString(coalesce("DB_PASSWORD", "hamidjon4424"))
	config.APP_PORT = cast.ToString(coalesce("APP_PORT", "2024"))
	config.SIGNING_KEY = cast.ToString(coalesce("SIGNING_KEY", "secret"))

	return config
}

func coalesce(key string, defaultValue interface{}) interface{} {
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	}
	return defaultValue
}
