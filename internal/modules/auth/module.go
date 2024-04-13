package auth

import (
	"github.com/gorilla/mux"
	"github.com/ngrink/url-shortener/internal/modules/users"
)

var Controller *AuthController
var Service *AuthService

func init() {
	Service = NewAuthService(users.Service)
	Controller = NewAuthController(Service)
}

func SetupAPIRoutes(public *mux.Router, private *mux.Router) {
	public.HandleFunc("/auth/login", Controller.LoginByCredentials).Methods("POST")
}
