package middleware

import (
	"github.com/PiamNaJa/CourseZ_Backend/configs"
	"github.com/PiamNaJa/CourseZ_Backend/models"
	"github.com/gofiber/fiber/v2"
)

func HaveVideo(c *fiber.Ctx) error {
	var count int64
	if err := configs.DB.Model(&models.Video{}).Where("course_id = ? and video_id = ?", c.Params("course_id"), c.Params("video_id")).Count(&count).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if count == 0 {
		return c.SendStatus(fiber.StatusNotFound)
	}
	return c.Next()
}
