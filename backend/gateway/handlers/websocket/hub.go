package websocket

import (
	"context"
	"encoding/json"
	"github.com/dark-vinci/wapp/backend/gateway/env"

	"github.com/google/uuid"
	"github.com/rs/zerolog"

	"github.com/dark-vinci/wapp/backend/sdk/constants"
	"github.com/dark-vinci/wapp/backend/sdk/utils/redis"
	//"github.com/dark-vinci/wapp/backend/sdk/utils/redis"
)

type Hub struct {
	// use redis instead of a map
	redis      redis.Operations
	Clients    map[*Client]struct{}
	Broadcast  chan []byte
	Register   chan *Client
	Unregister chan *Client
	ServerName string
	logger     zerolog.Logger
}

func NewHub(logger zerolog.Logger, e *env.Environment) *Hub {
	red := redis.NewRedis(&logger, e.RedisURL, e.RedisPassword, e.RedisUsername)

	return &Hub{
		logger:     logger,
		ServerName: uuid.New().String(),
		redis:      *red,
	}
}

func (h *Hub) Start() {
	// subscribe to events posted on redis
	go func() {
		b := make(chan []byte)

		h.redis.Subscribe(context.Background(), constants.WebsocketChannel, b)

		for {
			select {
			case msg := <-b:
				var c Message

				if err := json.Unmarshal(msg, &c); err != nil {
				}

				// ignore message sent by the same server
				if c.Server != h.ServerName {
					h.Broadcast <- msg
				}
			}
		}
	}()

	for {
		select {
		// register a client
		case client := <-h.Register:
			h.Clients[client] = struct{}{}

			//delete a client
		case client := <-h.Unregister:
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				close(client.Send)
			}

			// write to client
		case message := <-h.Broadcast:
			// broadcast to other servers
			go func() {
				// todo: retry
				_ = h.redis.Broadcast(context.Background(), "redis-key", message)
			}()

			// todo: add go routine
			for client := range h.Clients {
				select {
				case client.Send <- message:
					//todo: log the info
					h.logger.Info().Msg("message received from client")
				default:
					// if we cant send, close the send channel and delete client
					close(client.Send)
					delete(h.Clients, client)
				}
			}
		}
	}
}
