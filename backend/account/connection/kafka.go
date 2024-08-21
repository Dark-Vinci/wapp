package connection

import (
	"github.com/rs/zerolog"
	"github.com/segmentio/kafka-go"

	"github.com/dark-vinci/wapp/backend/account/env"
	"github.com/dark-vinci/wapp/backend/sdk/constants"
)

type Kafka struct {
	Writer *kafka.Writer
	Reader *kafka.Reader
}

type KafkaInterface interface{}

func (k *Kafka) Close() {
	_ = k.Reader.Close()
	_ = k.Writer.Close()
}

func NewKafka(z zerolog.Logger, e *env.Environment) *Kafka {
	log := z.With().Str(constants.PackageStrHelper, packageName).Logger()

	w := kafka.Writer{
		Addr:     kafka.TCP(e.KafkaURL),
		Topic:    "topic",
		Balancer: &kafka.LeastBytes{},
	}

	log.Info().Msg("connected to kafka writer")

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{e.KafkaURL},
		GroupID:  "groupID",
		Topic:    "topic",
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})

	log.Info().Msg("connected to kafka reader")

	return &Kafka{
		Writer: &w,
		Reader: r,
	}
}
