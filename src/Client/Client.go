package Client

import (
	"fmt"

	ApiRoutes "go-learning/src/Api"
	"go-learning/src/Interfaces"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

var App = fiber.New(fiber.Config{
	Prefork:       true,
	CaseSensitive: true,
	StrictRouting: true,
	ServerHeader:  "Fiber",
	AppName:       "Test App v1.0.1",
	JSONEncoder:   json.Marshal,
	JSONDecoder:   json.Unmarshal,
})

func RunServer(params Interfaces.SystemInterface) {
	App.Use(cors.New(cors.Config{
		AllowOrigins: params.CorsAllowOrigin,
		AllowHeaders: params.CorsAllowHeader,
	}))

	ApiRoutes.InitRoutes(App)
	App.Get("/metrics", monitor.New())

	App.Listen("localhost:" + fmt.Sprint(params.Port))
}
