package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func CORS() fiber.Handler {
	return cors.New(cors.Config{
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowMethods:     "GET, POST, PATCH, DELETE",
		AllowOrigins:    "*",

	})
}

func CORSEmbeddeerPolicy(c *fiber.Ctx) error {
	c.Set("Cross-Origin-Embedder-Policy", "credentialless")
	return c.Next()
}