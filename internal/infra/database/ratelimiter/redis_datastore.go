package ratelimiter

import "github.com/redis/go-redis/v9"

type RedisRateLimiterDatastore struct {
	Client *redis.Client
}

func (r *RedisRateLimiterDatastore) Name() string {
	return "RedisDatastore"
}

func (r *RedisRateLimiterDatastore) Get(key string) (int64, error) {
	// TODO: Implement this
	return 0, nil
}
