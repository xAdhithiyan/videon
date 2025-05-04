package api

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/xadhithiyan/videon/service/user"
)

type APIServer struct {
	address string
}

func CreateAPIServer(address string) *APIServer {
	return &APIServer{
		address: address,
	}
}

func (sv *APIServer) Run() error {
	router := chi.NewRouter()

	userStore := user.CreateStore()
	userHandler := user.CreateHandler(userStore)

	router.Route("/api/v1", func(r chi.Router) {
		userHandler.RegsterRoutes(r)
	})

	log.Printf("server started on %s", sv.address)
	return http.ListenAndServe(sv.address, router)
}
