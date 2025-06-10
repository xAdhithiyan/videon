package video

import (
	"encoding/binary"
	"encoding/json"
	"log"

	"github.com/xadhithiyan/videon/types"
)

type Handler struct {
	store types.VideoStore
}

func CreateHandler(store types.VideoStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) ParseData(msg []byte, userID int) (int, bool) {
	if len(msg) < 4 {
		log.Println("No Metadata Length")
		return -1, false
	}
	metaDataLength := binary.LittleEndian.Uint32(msg[:4])
	if len(msg) < int(metaDataLength+4) {
		log.Println("No MetaData")
		return -1, false
	}

	metaDataByte := msg[4 : 4+metaDataLength]
	var metaData types.MetaData
	if err := json.Unmarshal(metaDataByte, &metaData); err != nil {
		log.Println("Error unmarshalling metadata", err)
		return -1, false
	}

	if metaData.Id == 0 {
		err := h.store.AddVideoDB(userID, metaData)
		if err != nil {
			log.Print("Error while uploading video details to postgres: ", err)
		}
	}

	if err := h.store.UploadS3(metaData, msg); err != nil {
		log.Print("Error uploading into S3: ", err)
		return -1, false
	}
	return metaData.Id, true
}
