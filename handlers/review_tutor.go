package handlers

import (
	"errors"
	"strconv"

	"github.com/PiamNaJa/CourseZ_Backend/models"
	"github.com/PiamNaJa/CourseZ_Backend/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateReviewTutor(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var review models.Review_Tutor

		if err := c.BodyParser(&review); err != nil {
			return utils.BadRequest(err.Error())
		}

		teacher_id, err := strconv.ParseInt(c.Params("teacher_id"), 10, 64)
		review.TeacherID = int32(teacher_id)

		if err := utils.Validate.Struct(review); err != nil {
			return utils.BadRequest(err.Error())
		}

		if err != nil {
			return utils.Unexpected(err.Error())
		}

		if err := db.Create(&review).Error; err != nil {
			return utils.Unexpected(err.Error())
		}

		return c.Status(fiber.StatusCreated).JSON(&review)
	}
}

func GetReviewTutorByFilter(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var review []models.Review_Tutor

		if err := db.Where("teacher_id = ?", c.Params("teacher_id")).Find(&review).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return utils.NotFound(err.Error())
			}
			return utils.Unexpected(err.Error())
		}
		return c.Status(fiber.StatusOK).JSON(&review)
	}
}
