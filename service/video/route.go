package video

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/xadhithiyan/videon/types"
)

type Handler struct {
	store types.VideoStore
}

func CreateHandler(store types.VideoStore) *Handler {
	return &Handler{
		store: store,
	}
}

func (h *Handler) RegsterRoutes(r chi.Router) {
	r.Post("/video", h.UploadVideo)
}

func (h *Handler) UploadVideo(w http.ResponseWriter, r *http.Request) {
	str := h.store.UploadtoS3() + " working"

	w.Write([]byte(str))
}
