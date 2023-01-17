package handlers

import (
	"github.com/PiamNaJa/CourseZ_Backend/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SearchALL(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var course *[]models.Course

		if err := db.Model(&models.Course{}).Preload("Subject").Where("course_name = ?", c.Params("name")).Find(&course).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Course not found",
			})
		}

		var video *[]models.Video

		if err := db.Model(&models.Video{}).Preload("Reviews").Preload("Exercises").Where("video_name = ?", c.Params("name")).Find(&video).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Video not found",
			})
		}

		var user *[]models.User

		if err := db.Model(&models.User{}).Preload("Teacher").Where("fullname = ?", c.Params("name")).Find(&user).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Tutor not found",
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"course": &course,
			"video":  &video,
			"tutor":  &user,
		})
	}
}

func SearchCourse(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var course *[]models.Course

		if err := db.Model(&models.Course{}).Preload("Subject").Where("course_name = ?", c.Params("name")).Find(&course).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Course not found",
			})
		}
		return c.Status(fiber.StatusOK).JSON(&course)
	}
}

func SearchTutor(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var user *[]models.User

		if err := db.Model(&models.User{}).Preload("Teacher").Where("fullname = ?", c.Params("name")).Find(&user).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(&user)
	}
}

func SearchVideo(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var video *[]models.Video

		if err := db.Model(&models.Video{}).Preload("Reviews").Preload("Exercises").Where("video_name = ?", c.Params("name")).Find(&video).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Video not found",
			})
		}
		return c.Status(fiber.StatusOK).JSON(&video)
	}
}
