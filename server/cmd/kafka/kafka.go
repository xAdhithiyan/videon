package kafkaSetup

import (
	"context"
	"log"
	"sync"

	"github.com/segmentio/kafka-go"
)

var instance *KafkaConn
var once sync.Once

type KafkaConn struct {
	Compress   *kafka.Conn
	Thumbnails *kafka.Conn
	Transcode  *kafka.Conn
	Watermark  *kafka.Conn
	Summary    *kafka.Conn
}

// singleton implementation bruh
func CreateKafkaConn() *KafkaConn {
	once.Do(func() {
		instance = &KafkaConn{}
		instance.InitializeConns()
	})

	return instance
}

func (conns *KafkaConn) InitializeConns() {
	log.Print("setting kafka connections")
	var err error

	conns.Compress, err = kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "compressVideo", 0)

	if err != nil {
		log.Print("failed to dail leadaer for CompressVideo", err)
	}

	conns.Thumbnails, err = kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "generateThumbnails", 0)

	if err != nil {
		log.Print("fialed to dail leadaer for GenerateThumbnails", err)
	}

	conns.Transcode, err = kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "transcode", 0)

	if err != nil {
		log.Print("fialed to dail leadaer for transcode", err)
	}

	conns.Watermark, err = kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "watermark", 0)

	if err != nil {
		log.Print("fialed to dail leadaer for watermark", err)
	}

	conns.Summary, err = kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "summary", 0)

	if err != nil {
		log.Print("fialed to dail leadaer for summary", err)
	}

}
