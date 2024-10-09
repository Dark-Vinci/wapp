package kafka

import (
	"context"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
)

type Write struct {
	w kafka.Writer
}

//go:generate mockgen -source writer.go -destination ./mock/writer_mock.go -package mock Writer
type Writer interface {
	Close() error
	Write(ctx context.Context, id uuid.UUID, message []byte) error
}

func NewWriter(URL, topic string) *Writer {
	w := kafka.Writer{
		Addr:                   kafka.TCP(URL),
		Topic:                  topic,
		Balancer:               &kafka.LeastBytes{},
		MaxAttempts:            0,
		WriteBackoffMin:        0,
		WriteBackoffMax:        0,
		BatchSize:              0,
		BatchBytes:             0,
		BatchTimeout:           0,
		ReadTimeout:            0,
		WriteTimeout:           0,
		RequiredAcks:           0,
		Async:                  false,
		Completion:             nil,
		Compression:            0,
		Logger:                 nil,
		ErrorLogger:            nil,
		Transport:              nil,
		AllowAutoTopicCreation: false,
	}

	wr := &Write{w: w}

	writer := Writer(wr)

	return &writer
}

func (w *Write) Write(ctx context.Context, id uuid.UUID, message []byte) error {
	msg := kafka.Message{
		HighWaterMark: 0,
		Key:           []byte(id.String()),
		Value:         message,
		Headers:       nil,
		//WriterData:    nil,
		//Time:          time.Time{},
		//Partition:     0,
		//Offset:        0,
		//Topic:         "",
	}

	if err := w.w.WriteMessages(ctx, msg); err != nil {
		return err
	}

	return nil
}

func (w *Write) Close() error {
	return w.w.Close()
}
