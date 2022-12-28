package handlers

import (
	"github.com/PiamNaJa/CourseZ_Backend/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateCourse(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var course *models.Course

		if err := c.BodyParser(&course); err != nil {
			return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		if err := db.Model(models.User{}).Create(&course).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		// db.Model(models.Course{}).Create(&course)
		return c.Status(fiber.StatusOK).JSON(&course)
	}
}

func GetAll(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var course *[]models.Course
		db.Model(models.Course{}).Find(&course)
		return c.Status(fiber.StatusOK).JSON(&course)
	}
}

func GetById(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var course *models.Course
		id := c.Params("id")
		db.Model(models.Course{}).First(&course, id)
		return c.Status(fiber.StatusOK).JSON(&course)
	}
}
