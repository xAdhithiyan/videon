package video

import (
	"encoding/binary"
	"encoding/json"
	"log"
)

type Handler struct {
}

func CreateHandler() *Handler {
	return &Handler{}
}

type MetaData struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	VideoType   string `json:"videoType"`
	TotalChunks int    `json:"totalChunks"`
}

func (h *Handler) ParseData(msg []byte) (int, bool) {
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
	var metaData MetaData
	if err := json.Unmarshal(metaDataByte, &metaData); err != nil {
		log.Println("Error unmarshalling metadata", err)
		return -1, false
	}

	log.Print(metaData)
	// log.Print(msg[4+metaDataLength:])
	return 1, true
}
