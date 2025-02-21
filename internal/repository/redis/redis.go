package redis

import (
	"context"
	"fmt"
	"go-clean-arch/config"

	"github.com/go-redis/redis/v8"
)

type RedisRepository struct {
	client *redis.Client
}

func NewRedisRepository(cfg config.Config) (*RedisRepository, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr: cfg.RedisConfig.Host + ":" + cfg.RedisConfig.Port, // e.g., "localhost:6379"
	})

	// ping connection
	ctx := context.Background()
	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}

	return &RedisRepository{client: rdb}, nil
}

func (r *RedisRepository) Close() error {
	return r.client.Close()
}

func (r *RedisRepository) Set(ctx context.Context, key, value string) error {
	return r.client.Set(ctx, key, value, 0).Err()
}

func (r *RedisRepository) Get(ctx context.Context, key string) (string, error) {
	val, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", fmt.Errorf("key not found")
	} else if err != nil {
		return "", err
	}
	return val, nil
}
