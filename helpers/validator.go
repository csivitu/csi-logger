package helpers

import (
	"github.com/csivitu/csi-logger/schemas"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate *validator.Validate

func ValidateLogURLParams(c *fiber.Ctx) (*schemas.LogFetchSchema, error) {
	validate = validator.New()

	schema := new(schemas.LogFetchSchema)
	
	if schema.Limit == 0 {
		schema.Limit = 20
	}

	err := c.QueryParser(schema)
	if err != nil {
		return nil, err
	}

	err = validate.Struct(schema)
	if err != nil {
		return nil, err
	}

	return schema, nil
}
