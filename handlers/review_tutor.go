package handlers

import (
	"strconv"

	"github.com/PiamNaJa/CourseZ_Backend/constants"
	"github.com/PiamNaJa/CourseZ_Backend/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateReviewTutor(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var review models.Review_Tutor

		if err := c.BodyParser(&review); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		teacher_id, err := strconv.ParseInt(c.Params("teacher_id"), 10, 64)
		review.TeacherID = int32(teacher_id)

		if err := constants.Validate.Struct(review); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		if err := db.Model(&models.Review_Tutor{}).Create(&review).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusCreated).JSON(&review)
	}
}

func GetReviewTutorByFilter(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var review []models.Review_Tutor

		if err := db.Model(&models.Review_Tutor{}).Where("teacher_id = ?", c.Params("teacher_id")).Find(&review).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "No record",
			})
		}
		return c.Status(fiber.StatusOK).JSON(&review)
	}
}
