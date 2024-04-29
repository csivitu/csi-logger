package routers

import (
	"github.com/csivitu/csi-logger/controllers"
	"github.com/gofiber/fiber/v2"
)

func PingRouter(app *fiber.App) {
	app.Get("/ping", controllers.Ping)

}
