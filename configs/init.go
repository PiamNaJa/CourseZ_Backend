package configs

import (
	"errors"
	"log"
	"os"

	"github.com/PiamNaJa/CourseZ_Backend/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

var File *os.File

func Init() {
	file, err := os.OpenFile(".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	File = file

	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func CustomCors() func(*fiber.Ctx) error {
	return cors.New(
		cors.Config{
			AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
			AllowOrigins:     "*",
			AllowCredentials: true,
			AllowMethods:     "GET,POST,PUT,DELETE",
		},
	)
}

var CustomErrorHandler = func(c *fiber.Ctx, err error) error {
	// Status code defaults to 500
	code := fiber.StatusInternalServerError

	// Retrieve the custom status code if it's a *fiber.Error
	var defaulterr *fiber.Error
	var customerr *utils.Error
	if errors.As(err, &customerr) {
		code = customerr.Status
	}
	if errors.As(err, &defaulterr) {
		code = defaulterr.Code
	}

	// Set Content-Type: application/json; charset=utf-8
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)

	// Return status code with error message
	return c.Status(code).JSON(err.Error())
}
