package middlewares

import (
	"strings"

	"github.com/csivitu/csi-logger/helpers"
	"github.com/csivitu/csi-logger/initializers"
	"github.com/csivitu/csi-logger/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func ResourceProtect(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	tokenArr := strings.Split(authHeader, " ")

	if len(tokenArr) != 2 {
		return &fiber.Error{Code: fiber.StatusUnauthorized, Message: "You are Not Logged In."}
	}

	tokenString := tokenArr[1]
	resourceID, err := helpers.Decrypt([]byte(tokenString))
	if err != nil {
		return &fiber.Error{Code: fiber.StatusUnauthorized, Message: "Invalid API Key"}
	}

	parsedResourceID, err := uuid.Parse(string(resourceID))
	if err != nil {
		return &fiber.Error{Code: fiber.StatusUnauthorized, Message: "Invalid API Key"}
	}

	var resource models.Resource

	if err := initializers.DB.Where("id = ?", parsedResourceID).First(&resource).Error; err != nil {
		return &fiber.Error{Code: fiber.StatusUnauthorized, Message: "Resource Not Found"}
	}

	c.Set("resourceID", resource.ID.String())
	return c.Next()
}