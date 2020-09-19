package data

import (
	"context"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
)

var onceCache sync.Once

var (
	instanceCache *RedisClient
)

type CacheClient interface {
	Set(key string, value string, expiration time.Duration) error
	Get(key string) (string, error)
}

type RedisClient struct {
	client *redis.Client
}

func (r *RedisClient) Set(key string, value string, expiration time.Duration) error {
	return r.client.Set(context.Background(), key, value, expiration).Err()
}

func (r *RedisClient) Get(key string) (string, error) {
	return r.client.Get(context.Background(), key).Result()
}

func NewCacheClient() CacheClient {
	return initConnectionCache()
}

func initConnectionCache() CacheClient {
	onceCache.Do(func() { // <-- atomic, does not allow repeating

		conn := redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		})

		instanceCache = &RedisClient{client: conn}
	})

	return instanceCache
}
