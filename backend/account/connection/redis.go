package connection

import (
	"github.com/go-redis/redis/v8"
	"github.com/rs/zerolog"

	"github.com/dark-vinci/linkedout/backend/account/env"
)

type RedisClient struct {
	Val *redis.Client
}

func (r *RedisClient) Close() {
	r.Val.Close()
}

func NewRedisClient(z *zerolog.Logger, e *env.Environment) *RedisClient {
	log := z.With().Str("KEY", packageName).Logger()

	r := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis server address
		Password: "",               // No password set
		DB:       0,                // Use default DB
	})

	log.Info().Msg("connected to redis db")

	return &RedisClient{
		Val: r,
	}
}
