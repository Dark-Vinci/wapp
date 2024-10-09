package kafka

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

type Read struct {
	reader *kafka.Reader
}

//go:generate mockgen -source reader.go -destination ./mock/reader_mock.go -package mock Reader
type Reader interface {
	Fetch(ctx context.Context, ch chan<- kafka.Message)
	Close() error
}

func NewReader(topics []string, topic, URL, groupID string) *Reader {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{URL},
		GroupID:     groupID,
		Topic:       topic,
		MinBytes:    10e3, // 10KB
		MaxBytes:    10e6, // 10MB
		GroupTopics: topics,
	})

	read := &Read{reader: r}

	reader := Reader(read)

	return &reader
}

func (r *Read) Fetch(ctx context.Context, ch chan<- kafka.Message) {
	for {
		if msg, err := r.reader.ReadMessage(ctx); err != nil {
			fmt.Println("ERR", err)
			close(ch)
			return
		} else {
			ch <- msg
		}
	}
}

func (r *Read) Close() error {
	return r.reader.Close()
}
