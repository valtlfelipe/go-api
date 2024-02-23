package db

import (
	"context"

	"github.com/redis/go-redis/v9"
)

// DBInterface defines the interface for the DB type
type DBInterface interface {
	Get(key string) string
	Set(key string, value string)
}

type DB struct {
	ctx         context.Context
	redisClient *redis.Client
}

func NewDB(ctx context.Context, redisClient *redis.Client) DBInterface {
	return &DB{
		ctx:         ctx,
		redisClient: redisClient,
	}
}

func (service *DB) Get(key string) string {
	return service.redisClient.Get(service.ctx, key).Val()
}

func (service *DB) Set(key string, value string) {
	service.redisClient.Set(service.ctx, key, value, 0)
}
