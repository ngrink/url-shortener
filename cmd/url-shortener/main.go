package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/ngrink/url-shortener/internal/config"
	"github.com/ngrink/url-shortener/internal/database"
)

func main() {
	// load configuration
	config.Load()

	// connect to database
	database.Connect()

	// setup routes
	router := mux.NewRouter()

	// run server
	addr := fmt.Sprintf(":%s", os.Getenv("APP_PORT"))

	log.Printf("Starting server on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}
