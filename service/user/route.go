package user

import (
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
	// r.Post("/register")
}

func (h *Handler) LoginUser(r *chi.Route) {

}

func (h *Handler) RegisterUser(r *chi.Route) {

}
