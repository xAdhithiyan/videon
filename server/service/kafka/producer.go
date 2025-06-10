package kafka

import (
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/segmentio/kafka-go"
	kafkaSetup "github.com/xadhithiyan/videon/cmd/kafka"
	"github.com/xadhithiyan/videon/types"
	"github.com/xadhithiyan/videon/utils"
)

type Producer struct {
}

func CreateProducer() *Producer {
	return &Producer{}
}

func (p *Producer) SetupProducer(w http.ResponseWriter, r *http.Request) {
	allCons := kafkaSetup.CreateKafkaConn()

	var payload types.ProducerPayload
	if err := utils.ParseJson(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		validationErr := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, validationErr)
		return
	}

	operationsMap := map[*kafka.Conn]bool{
		allCons.Compress:   payload.Options.Compress,
		allCons.Thumbnails: payload.Options.Thumbnails,
		allCons.Transcode:  payload.Options.Transcode,
		allCons.Watermark:  payload.Options.Watermark,
		allCons.Summary:    payload.Options.Summary,
	}
	for conn, v := range operationsMap {
		if v {
			conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
			_, err := conn.WriteMessages(
				kafka.Message{Value: []byte(strconv.Itoa(payload.VideoId))},
			)

			if err != nil {
				utils.WriteError(w, http.StatusBadRequest, err)
				return
			}
		}
	}

	utils.WriteJson(w, http.StatusOK, map[string]string{"Message": "ok"}, nil)
}
