package kafka

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/segmentio/kafka-go"
	"github.com/xadhithiyan/videon/types"
	"github.com/xadhithiyan/videon/utils"
)

type Handler struct {
	producer types.Producer
}

func CreateHandler(producer types.Producer) *Handler {
	return &Handler{producer: producer}
}

func (k *Handler) RegisterRoutes(r chi.Router) {
	r.Get("/kafka/get-topics", k.GetKafkaTopics)
	r.Post("/kafka/producer", k.producer.SetupProducer)
}

func (k *Handler) GetKafkaTopics(w http.ResponseWriter, r *http.Request) {
	conn, err := kafka.Dial("tcp", "localhost:9092")
	if err != nil {
		log.Print(err)
	}

	partitions, err := conn.ReadPartitions()
	if err != nil {
		log.Print(err)
	}

	m := map[string]struct{}{}

	for _, p := range partitions {
		// struct{}{} is basically a 0 byte struct. map is used to only store unique values(multiple partitions with the same topic)
		m[p.Topic] = struct{}{}
	}

	utils.WriteJson(w, http.StatusOK, m, nil)
}
