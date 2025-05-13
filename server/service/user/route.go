package user

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/xadhithiyan/videon/config"
	"github.com/xadhithiyan/videon/service/auth"
	"github.com/xadhithiyan/videon/types"
	"github.com/xadhithiyan/videon/utils"
)

type Handler struct {
	store types.UserStore
}

func CreateHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegsterRoutes(r chi.Router) {
	r.Post("/register", h.RegisterUser)
	r.Post("/login", h.LoginUser)
}

func (h *Handler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var payload types.LoginUser
	if err := utils.ParseJson(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		validationError := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, validationError)
		return
	}

	user, err := h.store.GetUserByEmail(payload.Email)
	if err != nil {
		utils.WriteError(w, http.StatusUnauthorized, err)
		return
	}

	if user.Password != payload.Password {
		utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("invalid password"))
		return
	}

	token, err := auth.CreateJWT(config.Env.JWTSecrect, user.Id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	cookie := types.Cookie{Name: "token", Value: token}
	utils.WriteJson(w, http.StatusCreated, map[string]string{"message": "User Authenticated"}, &cookie)
}

func (h *Handler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var payload types.User
	if err := utils.ParseJson(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		validationError := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, validationError)
		return
	}

	if _, err := h.store.GetUserByEmail(payload.Email); err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user already exists"))
		return
	}

	err := h.store.CreateUser(types.User{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: payload.Password,
	})

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusCreated, map[string]string{"message": "User Registered"}, nil)
}
