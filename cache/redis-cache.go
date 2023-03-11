package cache

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/hmsayem/clean-architecture-implementation/entity"
	"os"
	"time"
)

type redisCache struct {
	*redis.Client
}

func NewRedisCache() EmployeeCache {
	return &redisCache{
		Client: redis.NewClient(&redis.Options{
			Addr:     os.Getenv("REDIS_SERVER_HOST"),
			Password: "", // no password set
			DB:       0,
		}),
	}
}

func (cache *redisCache) Set(key string, value *entity.Employee) error {
	d, err := json.Marshal(value)
	if err != nil {
		return err
	}
	if err := cache.Client.Set(context.Background(), key, d, 0*time.Second).Err(); err != nil {
		return err
	}
	return err
}

func (cache *redisCache) Get(key string) (*entity.Employee, error) {
	val, err := cache.Client.Get(context.Background(), key).Result()
	if err != nil {
		return nil, err
	}
	employee := entity.Employee{}
	if err := json.Unmarshal([]byte(val), &employee); err != nil {
		return nil, err
	}
	return &employee, nil
}

func (cache *redisCache) Del(key string) error {
	return cache.Client.Del(context.Background(), key).Err()
}
