package handlers

import (
	"github.com/PiamNaJa/CourseZ_Backend/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateSubject(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var subject *models.Subject
		if err := c.BodyParser(&subject); err != nil {
			return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		if err := db.Model(models.User{}).Create(&subject).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		db.Model(models.Subject{}).Create(&subject)
		return c.Status(fiber.StatusOK).JSON(&subject)
	}
}

func GetAllSubject(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var subject *[]models.Subject
		db.Model(models.Course{}).Find(&subject)
		return c.Status(fiber.StatusOK).JSON(&subject)
	}
}
