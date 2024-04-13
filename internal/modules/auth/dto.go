package auth

import "github.com/ngrink/url-shortener/internal/modules/users"

type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string     `json:"token"`
	User  users.User `json:"user"`
}
