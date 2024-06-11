package connection

import (
	"github.com/rs/zerolog"
	kafka "github.com/segmentio/kafka-go"

	"github.com/dark-vinci/linkedout/backend/account/env"
)

type Kafka struct {
	Writer *kafka.Writer
	Reader *kafka.Reader
}

func (k *Kafka) Close() {
	k.Reader.Close()
	k.Writer.Close()
}

func New(z zerolog.Logger, e *env.Environment) *Kafka {
	log := z.With().Str("KEY", "value").Logger()

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
