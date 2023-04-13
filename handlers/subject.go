package handlers

import (
	"github.com/PiamNaJa/CourseZ_Backend/utils"
	"github.com/PiamNaJa/CourseZ_Backend/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateSubject(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var subject models.Subject
		if err := c.BodyParser(&subject); err != nil {
			return utils.BadRequest(err.Error())
		}

		if err := utils.Validate.Struct(subject); err != nil {
			return utils.BadRequest(err.Error())
		}

		if err := db.Create(&subject).Error; err != nil {
			return utils.BadRequest(err.Error())
		}

		return c.Status(fiber.StatusCreated).JSON(&subject)
	}
}

func GetAllSubject(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var subject []models.Subject

		if err := db.Find(&subject).Error; err != nil {
			return utils.Unexpected(err.Error())
		}
		return c.Status(fiber.StatusOK).JSON(&subject)
	}
}
