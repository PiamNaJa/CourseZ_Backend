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
		var result []map[string]interface{}
		var name = "%" + c.Query("name") + "%"
        if err:= db.Table("users").
		Select("MIN(users.user_id) as user_id, MIN(user_teachers.teacher_id) as teacher_id, MIN(nickname) as nickname, MIN(fullname) as fullname, MIN(class_level) as class_level, MIN(users.picture) as picture, COALESCE(AVG(rating), 0.0) as rating").
		Joins("JOIN user_teachers ON users.user_id = user_teachers.user_id").
		Joins("JOIN courses ON user_teachers.teacher_id = courses.teacher_id").
		Joins("JOIN subjects ON courses.subject_id = subjects.subject_id").
		Joins("LEFT JOIN review_tutors ON user_teachers.teacher_id = review_tutors.teacher_id").
		Where("role = 'Teacher' OR role = 'Tutor'").
		Group("users.user_id").
		Having("nickname LIKE ?", name).
		Order("rating DESC").
        Find(&result).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		if result == nil {
			result = []map[string]interface{}{}
		}
        return c.Status(fiber.StatusOK).JSON(&result)
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
