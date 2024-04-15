package web

import (
	"github.com/gorilla/mux"
	"github.com/ngrink/url-shortener/internal/database"
	"github.com/ngrink/url-shortener/internal/modules/auth"
	"github.com/ngrink/url-shortener/internal/modules/urls"
)

var Controller *WebController

func init() {
	urlsRepository := urls.NewUrlsSqlRepository(database.DB)
	urlsService := urls.NewUrlsService(urlsRepository)

	Controller = NewWebController(urlsService)
}

func SetupRoutes(r *mux.Router) {
	pr := r.PathPrefix("/").Subrouter()
	pr.Use(auth.Authorized)

	pr.HandleFunc("/", Controller.Index).Methods("GET")
	pr.HandleFunc("/urls/{id}", Controller.Url).Methods("GET")
	r.HandleFunc("/register", Controller.Register).Methods("GET")
	r.HandleFunc("/login", Controller.Login).Methods("GET")
}
