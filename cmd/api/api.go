package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/xadhithiyan/videon/service/user"
)

type APIServer struct {
	db      *sql.DB
	address string
}

func CreateAPIServer(address string, db *sql.DB) *APIServer {
	return &APIServer{
		address: address,
		db:      db,
	}
}

func (sv *APIServer) Run() error {
	router := chi.NewRouter()

	userStore := user.CreateStore(sv.db)
	userHandler := user.CreateHandler(userStore)

	router.Route("/api/v1", func(r chi.Router) {
		userHandler.RegsterRoutes(r)
	})

	log.Printf("server started on %s", sv.address)
	return http.ListenAndServe(sv.address, router)
}
