package controllers

import (
	"github.com/csivitu/csi-logger/config"
	"github.com/csivitu/csi-logger/helpers"
	"github.com/csivitu/csi-logger/initializers"
	"github.com/csivitu/csi-logger/models"
	"github.com/csivitu/csi-logger/schemas"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateResource(c *fiber.Ctx) error {
	var reqBody schemas.ResourceCreateSchema

	if err := c.BodyParser(&reqBody); err != nil {
		return &fiber.Error{Code: fiber.StatusBadRequest, Message: "Validation Failed"}
	}
	userID, err := uuid.Parse(c.GetRespHeader("loggedInUserID"))
	if err != nil {
		return &fiber.Error{Code: fiber.StatusBadRequest, Message: "Invalid User ID"}
	}

	newResource := models.Resource{
		Name: reqBody.Name,
		HostedURL: reqBody.HostedURL,
		UserID: userID,
	}

	result := initializers.DB.Create(&newResource)
	if result.Error != nil {
		return helpers.AppError{Code: fiber.StatusInternalServerError, Message: config.DATABASE_ERROR , Err: result.Error}
	}

	apiKey, err := helpers.Encrypt([]byte(newResource.ID.String()))
	if err != nil {
		go helpers.LogServerError("Error while encrypting API Key.", err, c.Path())
		return helpers.AppError{Code: fiber.StatusInternalServerError, Message: config.SERVER_ERROR, Err: err}
	}

	newResource.APIKey = string(apiKey)

	result = initializers.DB.Save(&newResource)
	if result.Error != nil {
		return helpers.AppError{Code: fiber.StatusInternalServerError, Message: config.DATABASE_ERROR, Err: result.Error}
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"message": "Resource created",
		"api_key": newResource.APIKey,
	})
}

func GetAllResources(c *fiber.Ctx) error {
	var resources []models.Resource

	result := initializers.DB.Find(&resources)
	if result.Error != nil {
		return &fiber.Error{Code: fiber.StatusInternalServerError, Message: config.DATABASE_ERROR}
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"resources": resources,
	})
}

func DeleteResource (c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return &fiber.Error{Code: fiber.StatusBadRequest, Message: "Invalid ID"}
	}

	var resource models.Resource

	result := initializers.DB.Where("id = ?", id).First(&resource)
	if result.Error != nil {
		return helpers.AppError{Code: fiber.StatusNotFound, Message: "Resource not found", Err: result.Error}
	}

	result = initializers.DB.Delete(&resource)
	if result.Error != nil {
		return helpers.AppError{Code: fiber.StatusInternalServerError, Message: config.DATABASE_ERROR, Err: result.Error}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"message": "Resource deleted",
	})
}

func UpdateResource(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return &fiber.Error{Code: fiber.StatusBadRequest, Message: "Invalid ID"}
	}

	var reqBody schemas.ResourceUpdateSchema

	if err := c.BodyParser(&reqBody); err != nil {
		return &fiber.Error{Code: fiber.StatusInternalServerError, Message: "Validation Failed"}
	}

	var resource models.Resource

	result := initializers.DB.Where("id = ?", id).First(&resource)
	if result.Error != nil {
		return helpers.AppError{Code: fiber.StatusNotFound, Message: "Resource not found", Err: result.Error}
	}

	resource.HostedURL = reqBody.HostedURL

	result = initializers.DB.Save(&resource)
	if result.Error != nil {
		return helpers.AppError{Code: fiber.StatusInternalServerError, Message: config.DATABASE_ERROR, Err: result.Error}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"message": "Resource updated",
	})
}