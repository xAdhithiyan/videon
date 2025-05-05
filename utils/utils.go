package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/xadhithiyan/videon/types"
)

func ParseJson(r *http.Request, payload any) error {
	if r.Body == nil {
		return fmt.Errorf("request body empty")
	}
	return json.NewDecoder(r.Body).Decode(payload)
}

func WriteJson(w http.ResponseWriter, status int, payload any, cookieDetails *types.Cookie) error {
	if cookieDetails != nil {
		cookie := http.Cookie{Name: cookieDetails.Name, Value: cookieDetails.Value}
		http.SetCookie(w, &cookie)
		log.Print(cookie)
		log.Print("Cookie Set")
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(payload)
}

func WriteError(w http.ResponseWriter, status int, err error) error {
	return WriteJson(w, status, map[string]string{"error": err.Error()}, nil)
}

var Validate = validator.New(validator.WithRequiredStructEnabled())
