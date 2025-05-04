package user

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/xadhithiyan/videon/types"
)

type Handler struct {
	store types.UserStore
}

func CreateHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegsterRoutes(r chi.Router) {
	r.Post("/register", h.RegisterUser)
}

func (h *Handler) LoginUser(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	log.Print(r.Body)

	w.Write([]byte("working"))
}
