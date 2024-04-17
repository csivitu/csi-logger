package routers

import (
	"github.com/csivitu/csi-logger/controllers"
	middlewares "github.com/csivitu/csi-logger/middleware"
	"github.com/gofiber/fiber/v2"
)

func UserRouter(app *fiber.App) {

	app.Post("/user/register", middlewares.RootUserProtect, controllers.Register)
	app.Post("/user/admin", middlewares.RootUserProtect, controllers.MakeAdmin)

	userRouter := app.Group("/user")
	userRouter.Post("/login", controllers.Login)

}
