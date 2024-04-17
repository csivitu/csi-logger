package controllers

import (
	"github.com/csivitu/csi-logger/config"
	"github.com/csivitu/csi-logger/helpers"
	"github.com/csivitu/csi-logger/initializers"
	"github.com/csivitu/csi-logger/models"
	"github.com/csivitu/csi-logger/schemas"
	"github.com/gofiber/fiber/v2"
)


func MakeAdmin(c *fiber.Ctx) error {
	var reqbody schemas.MakeAdminSchema

	if err := c.BodyParser(&reqbody); err != nil {
		return &fiber.Error{Code: fiber.StatusBadRequest, Message: "Validation Failed"}
	}

	var existingUser models.User

	if err := initializers.DB.Where("email = ?", reqbody.Email).First(&existingUser).Error; err != nil {
		return helpers.AppError{Code: fiber.StatusInternalServerError, Message: config.DATABASE_ERROR , Err: err}
	}

	existingUser.Admin = true

	if err  := initializers.DB.Save(&existingUser).Error; err != nil{
		return helpers.AppError{Code: fiber.StatusInternalServerError, Message: config.DATABASE_ERROR , Err: err}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"message": "User is now an admin",
	})
}