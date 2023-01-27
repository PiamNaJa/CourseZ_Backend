package handlers

import (
	"github.com/PiamNaJa/CourseZ_Backend/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SearchALL(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var name = "%" + c.Query("name") + "%"
		var course *[]models.Course

		if err := db.Model(&models.Course{}).Preload("Subject").Where("course_name LIKE ?", &name).Find(&course).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Course not found",
			})
		}

		var video *[]models.Video

		if err := db.Model(&models.Video{}).Preload("Course").Where("video_name LIKE ?", &name).Find(&video).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Video not found",
			})
		}

		var user *[]models.User

		if err := db.Model(&models.User{}).Where("nickname LIKE ?", &name).Find(&user).Error; err != nil {
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
		var name = "%" + c.Query("name") + "%"

		if err := db.Model(&models.Course{}).Preload("Subject").Where("course_name LIKE ?", &name).Find(&course).Error; err != nil {
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
		var name = "%" + c.Query("name") + "%"

		if err := db.Model(&models.User{}).Preload("Teacher").Select("user_id", "nickname", "fullname", "picture").Preload("Teacher.Reviews").Where("nickname LIKE ?", &name).Find(&user).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Tutor not found",
			})
		}
		return c.Status(fiber.StatusOK).JSON(&user)
	}
}

func SearchVideo(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var video *[]models.Video
		var name = "%" + c.Query("name") + "%"

		if err := db.Model(&models.Video{}).Preload("Course").Where("video_name LIKE ?", &name).Find(&video).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Video not found",
			})
		}
		return c.Status(fiber.StatusOK).JSON(&video)
	}
}
