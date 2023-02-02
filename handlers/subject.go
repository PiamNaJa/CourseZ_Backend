package handlers

import (
	"github.com/PiamNaJa/CourseZ_Backend/constants"
	"github.com/PiamNaJa/CourseZ_Backend/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateSubject(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var subject models.Subject
		if err := c.BodyParser(&subject); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		if err := constants.Validate.Struct(subject); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(err.Error())
		}

		if err := db.Model(&models.Subject{}).Create(&subject).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusCreated).JSON(&subject)
	}
}

func GetAllSubject(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var subject []models.Subject

		if err := db.Model(&models.Subject{}).Find(&subject).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(&subject)
	}
}
