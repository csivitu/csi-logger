package middlewares

import (
	"strings"

	"github.com/csivitu/csi-logger/initializers"
	"github.com/gofiber/fiber/v2"
)

func RootUserProtect(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	tokenArr := strings.Split(authHeader, " ")

	if len(tokenArr) != 2 {
		return &fiber.Error{Code: 401, Message: "Unauthorized."}
	}

	tokenString := tokenArr[1]

	if tokenString != initializers.CONFIG.ROOT_PASSWORD{
		return &fiber.Error{Code: 401, Message: "Not a root user"}
	}

	return c.Next()
}
