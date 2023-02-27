package handlers

import (
	"strconv"

	"github.com/PiamNaJa/CourseZ_Backend/models"
	"github.com/PiamNaJa/CourseZ_Backend/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateReviewVideo(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var review models.Review_Video

		if err := c.BodyParser(&review); err != nil {
			return utils.BadRequest(err.Error())
		}

		video_id, err := strconv.ParseInt(c.Params("video_id"), 10, 64)
		review.VideoID = int32(video_id)

		if err != nil {
			return utils.BadRequest(err.Error())
		}
		if err := utils.Validate.Struct(review); err != nil {
			return utils.BadRequest(err.Error())
		}

		if err := db.Create(&review).Error; err != nil {
			return utils.Unexpected(err.Error())
		}

		return c.Status(fiber.StatusCreated).JSON(&review)
	}
}

func GetReviewVideoByFilter(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var review []models.Review_Video

		if err := db.Where("video_id = ?", c.Params("video_id")).Find(&review).Error; err != nil {
			return utils.HandleFindError(err)
		}
		return c.Status(fiber.StatusOK).JSON(&review)
	}
}
