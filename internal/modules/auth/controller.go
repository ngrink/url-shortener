package auth

import (
	"encoding/json"
	"net/http"
)

type AuthController struct {
	service *AuthService
}

func NewAuthController(service *AuthService) *AuthController {
	return &AuthController{service: service}
}

func (c *AuthController) LoginByCredentials(w http.ResponseWriter, r *http.Request) {
	var input LoginRequest

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := c.service.LoginByCredentials(input.Login, input.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
