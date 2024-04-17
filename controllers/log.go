package controllers

import (
	"fmt"
	"time"

	"github.com/csivitu/csi-logger/cache"
	"github.com/csivitu/csi-logger/config"
	"github.com/csivitu/csi-logger/helpers"
	"github.com/csivitu/csi-logger/initializers"
	"github.com/csivitu/csi-logger/models"
	"github.com/csivitu/csi-logger/schemas"
	"github.com/csivitu/csi-logger/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func AddLog(c *fiber.Ctx) error {
	resourceID, err := uuid.Parse(c.GetRespHeader("resourceID"))
	if err != nil {
		return &fiber.Error{Code: fiber.StatusBadRequest, Message: "Invalid Resource ID"}
	}

	var reqBody schemas.LogEntrySchema

	if err := c.BodyParser(&reqBody); err != nil {
		return &fiber.Error{Code: fiber.StatusBadRequest, Message: "Validation Failed"}
	}

	timestamp, err := time.Parse(time.RFC3339, reqBody.Timestamp)
	if err != nil {
		fmt.Println(err)
		return &fiber.Error{Code: fiber.StatusBadRequest, Message: "Invalid Timestamp"}
	}

	log := models.Log{
		Message:    reqBody.Message,
		Level:      reqBody.Level,
		Path:       reqBody.Path,
		ResourceID: resourceID,
		Timestamp:  timestamp,
	}

	if err := initializers.DB.Create(&log).Error; err != nil {
		go helpers.LogDatabaseError("Error creating log", err, "controllers/add_log.go")
		return helpers.AppError{Code: fiber.StatusInternalServerError, Message: config.DATABASE_ERROR, Err: err}
	}

	go cache.FlushCache(c.Context())

	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
		"message": "Log Created",
	})
}

func GetLogs(c *fiber.Ctx) error {

	resourceID, err := uuid.Parse(c.GetRespHeader("resourceID"))
	if err != nil {
		return &fiber.Error{Code: fiber.StatusBadRequest, Message: "Invalid Resource ID"}
	}

	urlParams, err := helpers.ValidateLogURLParams(c)
	if err != nil {
		return &fiber.Error{Code: fiber.StatusBadRequest, Message: "Invalid Query Params"}
	}

	cachedLogs, err := utils.FindCache(resourceID.String(), *urlParams, c.Context())
	if err == nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status": "success",
			"videos": cachedLogs,
		})
	}

	db := initializers.DB

	db = utils.FilterLogs(db, *urlParams)

	var logs []models.Log

	if err := db.Order("timestamp desc").Find(&logs, "resource_id = ?", resourceID).Error; err != nil {
		go helpers.LogDatabaseError("Error getting videos", err, "controllers/get_videos.go")
		return helpers.AppError{Code: fiber.StatusBadRequest, Message: config.DATABASE_ERROR, Err: err}
	}

	cacheKey := fmt.Sprintf("%s-%d-%d", resourceID.String(), urlParams.Limit, urlParams.Page)

	go cache.SetToCache(cacheKey, logs, c.Context())

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Logs successfully fetched",
		"logs":    logs,
	})
}
