package users

import (
	"github.com/gorilla/mux"
	"github.com/ngrink/url-shortener/internal/database"
)

var c *UsersController

func init() {
	r := NewUsersRepository(database.DB)
	s := NewUsersService(r)
	c = NewUsersController(s)
}

func SetupRoutes(r *mux.Router) {
	r.HandleFunc("/api/v1/users", c.CreateUser).Methods("POST")
	r.HandleFunc("/api/v1/users", c.GetAllUsers).Methods("GET")
	r.HandleFunc("/api/v1/users/{id}", c.GetUser).Methods("GET")
	r.HandleFunc("/api/v1/users/{id}", c.UpdateUser).Methods("PATCH")
	r.HandleFunc("/api/v1/users/{id}", c.DeleteUser).Methods("DELETE")
}
