package main

import (
	"fmt"
	"os"

	"github.com/ngrink/url-shortener/internal/app"
	"github.com/ngrink/url-shortener/internal/config"
	"github.com/ngrink/url-shortener/internal/database"
)

func main() {
	config.Load()
	database.Connect()

	addr := fmt.Sprintf(":%s", os.Getenv("APP_PORT"))

	server := app.NewServer(addr)
	server.Run()
}
