package controllers

import (
	"fmt"

	"github.com/csivitu/csi-logger/config"
	"github.com/gofiber/fiber/v2"
)

func DashboardView(c *fiber.Ctx) error {

	sess, err := config.Store.Get(c)
    if err != nil {
        panic(err)
    }

	keys := sess.Keys()

	fmt.Println(keys)
	
	return c.Render("dashboard", fiber.Map{
		"Title": "Dashboard",
	})
}