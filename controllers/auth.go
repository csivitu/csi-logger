package controllers

import (
	"time"

	"github.com/csivitu/csi-logger/config"
	"github.com/csivitu/csi-logger/helpers"
	"github.com/csivitu/csi-logger/initializers"
	"github.com/csivitu/csi-logger/models"
	"github.com/csivitu/csi-logger/schemas"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Register(c *fiber.Ctx) error {
	var reqBody schemas.UserCreateSchema

	if err := c.BodyParser(&reqBody); err != nil {
		return &fiber.Error{Code: fiber.StatusBadRequest, Message: "Validation Failed"}
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(reqBody.Password), 12)
	if err != nil {
		go helpers.LogServerError("Error while hashing Password.", err, c.Path())
		return helpers.AppError{Code: fiber.StatusInternalServerError, Message: config.SERVER_ERROR , Err: err}
	}

	newUser := models.User{
		Name: reqBody.Name,
		Email: reqBody.Email,
		Password: string(hash),
	}

	result := initializers.DB.Create(&newUser)
	if result.Error != nil {
		return helpers.AppError{Code: fiber.StatusInternalServerError, Message: config.DATABASE_ERROR, Err: result.Error}
	}

	access_token_claim := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": newUser.ID,
		"crt": time.Now().Unix(),
		"exp": time.Now().Add(config.ACCESS_TOKEN_TTL).Unix(),
	})

	access_token, err := access_token_claim.SignedString([]byte(initializers.CONFIG.JWT_SECRET))
	if err != nil {
		go helpers.LogServerError("Error while decrypting JWT Token.", err, c.Path())
		return helpers.AppError{Code: fiber.StatusInternalServerError, Message: config.SERVER_ERROR, Err: err}
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"message": "Account created",
		"token": access_token,
	})
}

func Login(c *fiber.Ctx) error {
	var reqBody schemas.UserLoginSchema

	if err := c.BodyParser(&reqBody); err != nil {
		return &fiber.Error{Code: fiber.StatusBadRequest, Message: "Validation Failed"}
	}

	var user models.User

	if err := initializers.DB.Session(&gorm.Session{SkipHooks: true}).First(&user, "email = ?", reqBody.Email).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return &fiber.Error{Code: fiber.StatusBadRequest, Message: "No account with these credentials found."}
		} else {
			return helpers.AppError{Code: fiber.StatusInternalServerError, Message: config.DATABASE_ERROR, Err: err}
		}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(reqBody.Password)); err != nil {
		return &fiber.Error{Code: fiber.StatusBadRequest, Message: "No account with these credentials found."}
	}

	access_token_claim := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"crt": time.Now().Unix(),
		"exp": time.Now().Add(config.ACCESS_TOKEN_TTL).Unix(),
	})

	access_token, err := access_token_claim.SignedString([]byte(initializers.CONFIG.JWT_SECRET))
	if err != nil {
		go helpers.LogServerError("Error while decrypting JWT Token.", err, c.Path())
		return helpers.AppError{Code: fiber.StatusInternalServerError, Message: config.SERVER_ERROR, Err: err}
	}

	sess, err := config.Store.Get(c)
    if err != nil {
        return err
    }

	sess.Set("jwt", access_token)

	if err := sess.Save(); err != nil {
        return err
    }

	return c.Redirect("/dashboard")

	// return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
	// 	"status": "success",
	// 	"message": "Successfully logged in!",
	// 	"token": access_token,
	// })
}

func LoginView(c *fiber.Ctx) error {
    c.Set("Cross-Origin-Embedder-Policy", "credentialless")
	return c.Render("index", fiber.Map{
		"Title": "Login | CSI Logger",
	})
}