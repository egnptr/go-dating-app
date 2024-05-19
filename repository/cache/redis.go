package cache

import (
	"time"
)

type RedisCache struct {
	host    string
	db      int
	expires time.Duration
}

func NewRedisCache(host string, db int, expires time.Duration) Repo {
	return &RedisCache{
		host:    host,
		db:      db,
		expires: expires,
	}
}
