package handlers

import (
	"github.com/PiamNaJa/CourseZ_Backend/models"
	"github.com/PiamNaJa/CourseZ_Backend/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func SearchALL(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var name = "%" + c.Query("name") + "%"
		var courses []models.Course
		var coursesVideos []models.Course = []models.Course{}

		if err := db.Preload("Videos", func(tx *gorm.DB) *gorm.DB { //Search video in course
			return tx.Where("video_name LIKE ?", &name)
		}).Preload("Subject").Preload("Videos.Reviews").
			Find(&courses).Error; err != nil {
			return utils.HandleRecordNotFoundErr(err)
		}

		for _, course := range courses {
			if len(*course.Videos) > 0 {
				coursesVideos = append(coursesVideos, course) //Add course that have video
			}
		}

		if err := db.Preload(clause.Associations).Preload("Videos.Reviews").Where("course_name LIKE ?", &name).Find(&courses).Error; err != nil {
			return utils.HandleRecordNotFoundErr(err)
		} //Search course

		var teachers []map[string]interface{}

		if err := db.Table("users").
			Select("MIN(users.user_id) as user_id, MIN(user_teachers.teacher_id) as teacher_id, MIN(nickname) as nickname, MIN(fullname) as fullname, MIN(class_level) as class_level, MIN(users.picture) as picture, COALESCE(AVG(rating), 0.0) as rating").
			Joins("JOIN user_teachers ON users.user_id = user_teachers.user_id").
			Joins("JOIN courses ON user_teachers.teacher_id = courses.teacher_id").
			Joins("JOIN subjects ON courses.subject_id = subjects.subject_id").
			Joins("LEFT JOIN review_tutors ON user_teachers.teacher_id = review_tutors.teacher_id").
			Where("role = 'Teacher' OR role = 'Tutor'").
			Group("users.user_id").
			Having("nickname LIKE ?", name).
			Order("rating DESC").
			Find(&teachers).Error; err != nil {
			return utils.HandleRecordNotFoundErr(err)
		}
		if teachers == nil {
			teachers = []map[string]interface{}{}
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"courses": &courses,
			"videos":  &coursesVideos,
			"tutors":  &teachers,
		})
	}
}

func SearchCourse(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var course []models.Course
		var name = "%" + c.Query("name") + "%"

		if err := db.Preload(clause.Associations).Preload("Videos.Reviews").Where("course_name LIKE ?", &name).Find(&course).Error; err != nil {
			return utils.HandleRecordNotFoundErr(err)
		}
		return c.Status(fiber.StatusOK).JSON(&course)
	}
}

func SearchTutor(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var result []map[string]interface{}
		var name = "%" + c.Query("name") + "%"
		if err := db.Table("users").
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
			return utils.HandleRecordNotFoundErr(err)
		}
		if result == nil {
			result = []map[string]interface{}{}
		}
		return c.Status(fiber.StatusOK).JSON(&result)
	}
}

func SearchVideo(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var courses, coursesSearch []models.Course
		var name = "%" + c.Query("name") + "%"

		if err := db.Preload("Videos", func(tx *gorm.DB) *gorm.DB {
			return tx.Where("video_name LIKE ?", &name)
		}).Preload("Subject").Preload("Videos.Reviews").
			Find(&courses).Error; err != nil {
			return utils.HandleRecordNotFoundErr(err)
		}

		for _, course := range courses {
			if len(*course.Videos) > 0 {
				coursesSearch = append(coursesSearch, course)
			}
		}
		return c.Status(fiber.StatusOK).JSON(&coursesSearch)
	}
}
