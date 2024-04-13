package users

import (
	"github.com/gorilla/mux"
	"github.com/ngrink/url-shortener/internal/database"
)

var Controller *UsersController
var Service *UsersService
var Repository *IUsersRepository

func init() {
	Repository := NewUsersSQLRepository(database.DB)
	Service = NewUsersService(Repository)
	Controller = NewUsersController(Service)
}

func SetupAPIRoutes(public *mux.Router, protected *mux.Router) {
	public.HandleFunc("/users", Controller.CreateUser).Methods("POST")

	protected.HandleFunc("/users", Controller.GetAllUsers).Methods("GET")
	protected.HandleFunc("/users/{id}", Controller.GetUser).Methods("GET")
	protected.HandleFunc("/users/{id}", Controller.UpdateUser).Methods("PATCH")
	protected.HandleFunc("/users/{id}", Controller.DeleteUser).Methods("DELETE")
}
