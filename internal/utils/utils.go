package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var Validate = validator.New(validator.WithRequiredStructEnabled())

func ParseJSON(w http.ResponseWriter, r *http.Request, v any) {
	if r.Body == nil {
		WriteError(w, http.StatusBadRequest, fmt.Errorf("missing request body"))
	}

	err := json.NewDecoder(r.Body).Decode(v)
	if err != nil {
		WriteError(w, http.StatusBadRequest, err)
	}
}

func WriteJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, err)
	}
}

func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJSON(w, status, map[string]string{
		"error": err.Error(),
	})
}
