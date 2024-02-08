// Package redisClient | https://github.com/redis/go-redis
package redisClient

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

var (
	ctx, connection = context.Background(), redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "89831143406",
		DB:       0,
	})
)

// SetValue Функция сохранения значения с заданным временем хранения
func SetValue(key string, value any, duration time.Duration) {
	err := connection.Set(ctx, key, value, duration).Err()
	if err != nil {
		panic(err)
	}
}

// GetValue функция получения значения по ключу
func GetValue(key string) (result any) {
	val, err := connection.Get(ctx, key).Result()
	if err != nil {
		panic(err)
	}

	return val
}
