package ratelimiter

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"

	"github.com/mathcale/goexpert-rate-limiter-challenge/config"
)

type RateLimiterDatastoreInterface interface {
	Name() string
	Get(key string) (int64, error)
}

func NewRedisRateLimiterDatastore(
	cfg config.Conf,
	logger zerolog.Logger,
) (RateLimiterDatastoreInterface, error) {
	addr := fmt.Sprintf("%s:%d", cfg.RedisHost, cfg.RedisPort)

	logger.Debug().Msgf("Connecting to Redis on [%s]", addr)

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: cfg.RedisPassword,
		DB:       cfg.RedisDB,
	})

	if _, err := client.Ping(context.Background()).Result(); err != nil {
		return nil, err
	}

	logger.Debug().Msgf("Redis successfully connected")

	return &RedisRateLimiterDatastore{
		Client: client,
	}, nil
}
