package handlers

import (
	"strconv"

	"github.com/PiamNaJa/CourseZ_Backend/constants"
	"github.com/PiamNaJa/CourseZ_Backend/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateReviewVideo(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var review models.Review_Video

		if err := c.BodyParser(&review); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		video_id, err := strconv.ParseInt(c.Params("video_id"), 10, 64)
		review.VideoID = int32(video_id)

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

		if err := db.Model(&models.Review_Video{}).Create(&review).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusCreated).JSON(&review)
	}
}

func GetReviewVideoByFilter(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var review []models.Review_Video

		if err := db.Model(&models.Review_Video{}).Where("video_id = ?", c.Params("video_id")).Find(&review).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "No record",
			})
		}
		return c.Status(fiber.StatusOK).JSON(&review)
	}
}
