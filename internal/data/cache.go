package data

import (
	"context"
	"encoding/json"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
)

var onceCache sync.Once

var (
	instanceCache *RedisClient
)

type CacheClient interface {
	GetInt(key string) (int, error)
	SetInt(key string, value int, expiration time.Duration) error
	Set(key string, value interface{}, expiration time.Duration) error
	Get(key string, src interface{}) error
}

type RedisClient struct {
	client *redis.Client
}

func (r *RedisClient) Set(key string, value interface{}, expiration time.Duration) error {
	cacheEntry, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return r.client.Set(context.Background(), key, cacheEntry, expiration).Err()
}

func (r *RedisClient) SetInt(key string, value int, expiration time.Duration) error {
	return r.client.Set(context.Background(), key, value, expiration).Err()
}

func (r *RedisClient) Get(key string, src interface{}) error {
	val, err := r.client.Get(context.Background(), key).Result()
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(val), &src)
}

func (r *RedisClient) GetInt(key string) (int, error) {
	val, err := r.client.Get(context.Background(), key).Result()
	if err != nil {
		if err == redis.Nil {
			return 0, nil
		}
		return 0, err
	}
	return strconv.Atoi(val)
}

func NewCacheClient() CacheClient {
	return initConnectionCache()
}

func initConnectionCache() CacheClient {
	onceCache.Do(func() { // <-- atomic, does not allow repeating

		conn := redis.NewClient(&redis.Options{
			Addr:     "redis-master.default.svc.cluster.local:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		})

		if err := conn.Ping(context.Background()).Err(); err != nil {
			log.Fatalf("Unable to connect to redis : %s\n", err)
		}

		instanceCache = &RedisClient{client: conn}
	})

	return instanceCache
}
