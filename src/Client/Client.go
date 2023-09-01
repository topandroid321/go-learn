package Client

import (
	"fmt"

	ApiRoutes "go-learning/src/Api"
	Middleware "go-learning/src/Api/middleware"
	"go-learning/src/Interfaces"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func RunServer(params Interfaces.SystemInterface) {
	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "Test App v1.0.1",
		JSONEncoder:   json.Marshal,
		JSONDecoder:   json.Unmarshal,
	})
	app.Use(cors.New(cors.Config{
		AllowOrigins: params.CorsAllowOrigin,
		AllowHeaders: params.CorsAllowHeader,
	}))

	app.Use(func(c *fiber.Ctx) error { return Middleware.CheckAuth(c) }) // auth middleware check
	app.Use(logger.New())                                                // logging

	ApiRoutes.InitRoutes(app)
	app.Get("/metrics", monitor.New())

	app.Listen("localhost:" + fmt.Sprint(params.Port))
}
