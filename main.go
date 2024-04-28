package main

import (
	"github.com/csivitu/csi-logger/config"
	"github.com/csivitu/csi-logger/helpers"
	"github.com/csivitu/csi-logger/initializers"
	"github.com/csivitu/csi-logger/routers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
	initializers.AddLogger()
	initializers.ConnectToCache()
	initializers.AutoMigrate()
}

func main() {
	defer initializers.LoggerCleanUp()

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		ErrorHandler: helpers.ErrorHandler,
		Views: engine,
	})

	app.Use(helmet.New())
	app.Use(config.CORS())

	app.Use(logger.New())

	app.Use(config.CORSEmbeddeerPolicy)
	
	routers.Config(app)

	app.Listen(":" + initializers.CONFIG.PORT)
}
