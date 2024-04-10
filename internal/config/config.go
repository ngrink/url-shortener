package config

import (
	"log"

	"github.com/joho/godotenv"
)

type Config struct {
	APP_PORT          int
	POSTGRES_HOST     string
	POSTGRES_PORT     int
	POSTGRES_USER     string
	POSTGRES_PASSWORD string
	POSTGRES_DB       string
	PGADMIN_EMAIL     string
	PGADMIN_PASSWORD  string
}

func init() {
	Load()
}

func Load() {
	err := godotenv.Load("./config/.env.local")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
