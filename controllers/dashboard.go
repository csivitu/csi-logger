package controllers

import (

	"github.com/csivitu/csi-logger/initializers"
	"github.com/csivitu/csi-logger/models"
	"github.com/gofiber/fiber/v2"
)

func DashboardView(c *fiber.Ctx) error {
	user := c.Locals("loggedInUser").(models.User)

	var resources []models.Resource
	if err := initializers.DB.Find(&resources); err.Error != nil {
		return c.Render("error", fiber.Map{
			"Status_Code": 	400,
			"Message":     "No resources found",
			"Title":       "Error",
		})
	}

	if !user.Admin {
        for i := range resources {
            resources[i].APIKey = ""
        }
    }

	return c.Render("dashboard", fiber.Map{
		"Title": "Dashboard",
		"Resources": resources,
	})
}
