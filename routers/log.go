package routers

import (
	"github.com/csivitu/csi-logger/controllers"
	middlewares "github.com/csivitu/csi-logger/middleware"
	"github.com/gofiber/fiber/v2"
)

func LogRouter(app *fiber.App) {

	logRouter := app.Group("/logger", middlewares.ResourceProtect)

	logRouter.Post("/", controllers.AddLog)
	logRouter.Get("/", controllers.GetLogs)
	// logRouter.Get("/filter", controllers.Login)
	// logRouter.Delete("/:id", middlewares.AdminProtect, controllers.Login)

}
