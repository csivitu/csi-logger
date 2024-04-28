package middlewares

import (
	"strings"

	"github.com/csivitu/csi-logger/initializers"
	"github.com/csivitu/csi-logger/models"
	"github.com/gofiber/fiber/v2"
)

func RootUserProtect(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	tokenArr := strings.Split(authHeader, " ")

	if len(tokenArr) != 2 {
		return &fiber.Error{Code: fiber.StatusUnauthorized, Message: "Unauthorized."}
	}

	tokenString := tokenArr[1]

	if tokenString != initializers.CONFIG.ROOT_PASSWORD{
		return &fiber.Error{Code: fiber.StatusUnauthorized, Message: "Not a root user"}
	}

	return c.Next()
}


func AdminProtect(c *fiber.Ctx) error {
	var user models.User
	err := verifyToken(c, &user)
	if err != nil {
		return err
	}

	if !user.Admin {
		return &fiber.Error{Code: fiber.StatusUnauthorized, Message: "Not an admin"}
	}
	c.Set("loggedInUserID", user.ID.String())

	return c.Next()
}