package routers

import (
	"github.com/csivitu/csi-logger/controllers"
	"github.com/gofiber/fiber/v2"
)

func ViewsRouter(app *fiber.App) {

	app.Get("/", controllers.LoginView)
	app.Get("/dashboard", controllers.DashboardView)
}
