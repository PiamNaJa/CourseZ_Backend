package handlers

import (
	"github.com/PiamNaJa/CourseZ_Backend/constants"
	"github.com/PiamNaJa/CourseZ_Backend/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreatePost(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var post *models.Post

		if err := c.BodyParser(&post); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		if err := constants.Validate.Struct(post); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(err)
		}

		if err := db.Model(&models.Post{}).Create(&post).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusCreated).JSON(&post)
	}
}
