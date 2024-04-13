package urls

import (
	"github.com/gorilla/mux"
	"github.com/ngrink/url-shortener/internal/database"
)

var Controller *UrlsController
var Service *IUrlsService
var Repository *IUrlsRepository

func init() {
	Repository := NewUrlsSqlRepository(database.DB)
	Service := NewUrlsService(Repository)
	Controller = NewUrlsController(Service)
}

func SetupRoutes(router *mux.Router) {
	router.HandleFunc("/{key}", Controller.RedirectByKey).Methods("GET")
}

func SetupAPIRoutes(public *mux.Router, protected *mux.Router) {
	protected.HandleFunc("/urls", Controller.CreateUrl).Methods("POST")
	protected.HandleFunc("/urls", Controller.GetAllUrls).Methods("GET")
	protected.HandleFunc("/users/{userId}/urls", Controller.GetUserUrls).Methods("GET")
	protected.HandleFunc("/urls/{id}", Controller.GetUrl).Methods("GET")
	protected.HandleFunc("/urls/{id}", Controller.DeleteUrl).Methods("DELETE")
	protected.HandleFunc("/urls/{id}/visits", Controller.GetUrlVisits).Methods("GET")
}
