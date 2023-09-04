package Redis

import (
	"context"

	"go-learning/src/Utils/RedisClient"

	"github.com/gofiber/fiber/v2"
)

type GetRedisKey struct {
	Key string `json:"key"`
}

func GetRedisByKey(c *fiber.Ctx) error {
	key := GetRedisKey{}
	ctx := context.Background()

	// CONNECT TO REDIS
	errRedis := RedisClient.InitRedisConnection()
	if errRedis != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"statusCode": fiber.StatusBadRequest,
			"error":      "Failed connect to redis",
		})
	}
	// defer Redis.CloseRedisConnection()

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

	data, errRedis := RedisClient.Client.Get(ctx, key.Key).Result()
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
