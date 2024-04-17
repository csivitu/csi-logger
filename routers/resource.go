package routers

import (
	"github.com/csivitu/csi-logger/controllers"
	middlewares "github.com/csivitu/csi-logger/middleware"
	"github.com/gofiber/fiber/v2"
)

func ResourceRouter(app *fiber.App) {

	resourceRoutes := app.Group("/resource", middlewares.AdminProtect )

	resourceRoutes.Post("/", controllers.CreateResource)
	resourceRoutes.Get("/", controllers.GetAllResources)
	resourceRoutes.Delete("/:id", controllers.DeleteResource)
	resourceRoutes.Patch("/:id", controllers.UpdateResource)

}
