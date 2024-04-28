package middlewares

import (

	"github.com/csivitu/csi-logger/config"
	"github.com/csivitu/csi-logger/initializers"
	"github.com/csivitu/csi-logger/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func verifyToken(c *fiber.Ctx, user *models.User) error {

	sess, err := config.Store.Get(c)
	if err != nil {
		return c.Render("error", fiber.Map{
			"Status_Code": 	fiber.StatusForbidden,
			"Message":     "Session not found!",
			"Title":       "Error",
		})
	}

	userID := sess.Get("userID")
	if userID == nil || userID == "" {
		return c.Render("error", fiber.Map{
			"Status_Code": 	fiber.StatusUnauthorized,
			"Message":     "Session not found!",
			"Title":       "Error",
		})
	}

	if err := initializers.DB.First(user, "id = ?", userID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Render("error", fiber.Map{
				"Status_Code": 	fiber.StatusNotFound,
				"Message":     "User not found!",
				"Title":       "Error",
			})
		}
		return c.Render("error", fiber.Map{
			"Status_Code": 	fiber.StatusInternalServerError,
			"Message":     "Internal Server Error!",
			"Title":       "Error",
		})
	}
	return nil

}

func Protect(c *fiber.Ctx) error {

	var user models.User
	err := verifyToken(c, &user)
	
	if err != nil {
		return err
	}

	c.Locals("loggedInUser", user)

	return c.Next()
}
