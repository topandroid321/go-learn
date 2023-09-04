package Redis

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

type GetRedisKey struct {
	Key string `json:"key"`
}

func GetRedisByKey(c *fiber.Ctx) error {
	key := GetRedisKey{}
	ctx := context.Background()

	// CONNECT TO REDIS
	RedisClient := redis.NewClient(&redis.Options{
		Addr:     "194.233.95.186:6381",
		Password: "eP5gtsmn8SSSUfZQBkJIcaj0pcy8t+c4XPuiPik8gxMOan6XoTCLQuDXV8g+nLRIYpuAYdgywu9gJB+X",
		DB:       0,
	})

	if errBody := c.BodyParser(&key); errBody != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"statusCode": fiber.StatusBadRequest,
			"error":      "Cannot parse JSON",
		})
	}

	if key.Key == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"statusCode": fiber.StatusBadRequest,
			"error":      "key, is Required",
		})
	}

	data, errRedis := RedisClient.Get(ctx, key.Key).Result()
	if errRedis != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"statusCode": fiber.StatusBadRequest,
			"error":      "Failed get data from redis",
		})
	}

	return c.JSON(fiber.Map{
		"statusCode": fiber.StatusOK,
		"data": fiber.Map{
			"valid":    true,
			"messages": "success-get-data-redis",
			"key":      key.Key,
			"data":     data,
		},
	})
}
