package websocket

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/websocket"
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

	defer conn.Close()

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		if messageType == 1 || messageType == 2 {
			_, ok := ws.videoFuns.ParseData(p)
			if ok {
				if err := conn.WriteMessage(messageType, []byte("Chunk recieved")); err != nil {
					log.Println("Write error: ", err)
					break
				}
			} else {
				if err := conn.WriteMessage(messageType, []byte("Chunk not recieved")); err != nil {
					log.Println("Write error: ", err)
					break
				}

			}
		}

	}
}
