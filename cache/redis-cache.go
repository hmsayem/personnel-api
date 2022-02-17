package cache

import (
	"context"
	"encoding/json"
	redis "github.com/go-redis/redis/v8"
	"github.com/hmsayem/clean-architecture-implementation/entity"
	"time"
)

type redisCache struct {
	host   string
	db     int
	expire time.Duration
}

func NewRedisCache(host string, db int, expire time.Duration) *redisCache {
	return &redisCache{
		host:   host,
		db:     db,
		expire: expire,
	}
}

func (cache *redisCache) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cache.host,
		Password: "", // no password set
		DB:       cache.db,
	})
}

func (cache *redisCache) Set(key string, value *entity.Employee) error {
	client := cache.getClient()
	json, err := json.Marshal(value)
	if err != nil {
		return err
	}
	if err := client.Set(context.Background(), key, json, cache.expire*time.Second).Err(); err != nil {
		return err
	}
	return err
}

func (cache *redisCache) Get(key string) (*entity.Employee, error) {
	client := cache.getClient()
	val, err := client.Get(context.Background(), key).Result()
	if err != nil {
		return nil, err
	}
	employee := entity.Employee{}
	if err := json.Unmarshal([]byte(val), &employee); err != nil {
		return nil, err
	}
	return &employee, nil
}
