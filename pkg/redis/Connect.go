package redis

import (
	"sync"
	"uptime/pkg/config"

	"github.com/redis/go-redis/v9"
)

var doOnce sync.Once
var singletonConnection *redis.Client

func Connect() *redis.Client {
	doOnce.Do(func() {
		redisAddr := config.Get("REDIS_HOST") + ":" + config.Get("REDIS_PORT")

		singletonConnection = redis.NewClient(&redis.Options{
			Addr:     redisAddr,
			Password: config.Get("REDIS_PASSWORD"),
			DB:       0,
		})
	})

	return singletonConnection
}
