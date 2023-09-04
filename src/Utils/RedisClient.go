package Utils

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var Client *redis.Client

func InitRedisConnection() error {
	ctx := context.Background()
	Client := redis.NewClient(&redis.Options{
		Addr:     "194.233.95.186:6381",
		Password: "eP5gtsmn8SSSUfZQBkJIcaj0pcy8t+c4XPuiPik8gxMOan6XoTCLQuDXV8g+nLRIYpuAYdgywu9gJB+X",
		// DB:       0,
	})

	_, err := Client.Ping(ctx).Result()
	if err != nil {
		return err
	}

	return nil
}

func CloseRedisConnection() error {
	if Client != nil {
		return Client.Close()
	}
	return nil
}
