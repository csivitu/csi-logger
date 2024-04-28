package routers

import (
	"github.com/csivitu/csi-logger/controllers"
	middlewares "github.com/csivitu/csi-logger/middleware"
	"github.com/gofiber/fiber/v2"
)

func ViewsRouter(app *fiber.App) {

	app.Get("/", controllers.LoginView)
	app.Get("/dashboard", middlewares.Protect, controllers.DashboardView)
}
