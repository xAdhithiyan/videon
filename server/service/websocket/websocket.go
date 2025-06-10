package websocket

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/websocket"
	"github.com/xadhithiyan/videon/service/auth"
	"github.com/xadhithiyan/videon/types"
)

type ws struct {
	videoFuns types.VideoFuns
}

func CreateWS(videoFuns types.VideoFuns) *ws {
	return &ws{videoFuns: videoFuns}
}

func (ws *ws) RegisterRouters(r chi.Router) {
	r.Get("/ws", ws.wsHandler)
}

var upgrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (ws *ws) wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrade.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade error: ", err)
		return
	}

	var connMutex sync.Mutex
	defer conn.Close()

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		token, _ := r.Cookie("token")
		userID, _ := auth.AuthenticateJwt(token.Value)

		if messageType == 1 || messageType == 2 {
			go func() {
				chuckId, ok := ws.videoFuns.ParseData(p, userID)

				connMutex.Lock()
				defer connMutex.Unlock()

				chunkReponse := types.ChunkResponse{ChunkId: chuckId, Uploaded: false}
				if ok {
					chunkReponse.Uploaded = true
					chuckRepnseBytes, _ := json.Marshal(chunkReponse)
					if err := conn.WriteMessage(2, chuckRepnseBytes); err != nil {
						log.Println("Write error: ", err)

					}
				} else {
					chuckRepnseBytes, _ := json.Marshal(chunkReponse)
					if err := conn.WriteMessage(2, chuckRepnseBytes); err != nil {
						log.Println("Write error: ", err)

					}
				}
			}()
		}

	}
}
