package connection

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/rs/zerolog"

	"github.com/dark-vinci/wapp/backend/account/env"
	"github.com/dark-vinci/wapp/backend/sdk/constants"
)

type RedisClient struct {
	Val *redis.Client
}

type RedisOps interface {
	Get(ctx context.Context, key string) ([]byte, error)
	Set(ctx context.Context, key string, value []byte) error
}

func (r *RedisClient) Close() {
	_ = r.Val.Close()
}

func NewRedis(z *zerolog.Logger, e *env.Environment) *RedisOps {
	log := z.With().Str(constants.PackageStrHelper, packageName).Logger()

	r := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis server address
		Password: "",               // No password set
		DB:       0,                // Use default DB
	})

	log.Info().Msg("connected to redis db")

	red := &RedisClient{
		Val: r,
	}

	redOps := RedisOps(red)

	return &redOps
}

func (r *RedisClient) Get(ctx context.Context, key string) ([]byte, error) {
	return r.Val.WithContext(ctx).Get(ctx, key).Bytes()
}

func (r *RedisClient) Set(ctx context.Context, key string, value []byte) error {
	return r.Val.WithContext(ctx).Set(ctx, key, value, 0).Err()
}
