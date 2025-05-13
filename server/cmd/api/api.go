package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/xadhithiyan/videon/middlware"
	"github.com/xadhithiyan/videon/service/user"
	"github.com/xadhithiyan/videon/service/video"
	"github.com/xadhithiyan/videon/service/websocket"
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

	videoStore := video.CreateStore(sv.db)
	videoHandler := video.CreateHandler(videoStore)
	ws := websocket.CreateWS(videoHandler)

	router.Route("/api/v1", func(r chi.Router) {
		r.Use(middlware.AuthVerification)
		userHandler.RegsterRoutes(r)

		ws.RegisterRouters(r)
	})

	log.Printf("server started on %s", sv.address)
	return http.ListenAndServe(sv.address, router)
}
