package cache

import (
	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	Client *redis.Client
}

func NewRedisCache(host string, db int) Repo {
	client := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: "",
		DB:       db,
	})

	return &RedisCache{
		Client: client,
	}
}
