package handlers

//CompileDaemon -command="./CourseZ_Backend"
import (
	"github.com/PiamNaJa/CourseZ_Backend/ai"
	"github.com/PiamNaJa/CourseZ_Backend/models"
	"github.com/PiamNaJa/CourseZ_Backend/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateCourse(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var course models.Course

		if err := c.BodyParser(&course); err != nil {
			return utils.BadRequest(err.Error())
		}
		var subject models.Subject
		if err := c.BodyParser(&subject); err != nil {
			return utils.BadRequest(err.Error())
		}

		if err := db.Where("subject_title = ? AND class_level = ?", subject.Subject_title, subject.Class_level).First(&subject).Error; err != nil {
			return utils.HandleFindError(err)
		}
		course.SubjectID = subject.Subject_id
		course.Subject = &subject

		if err := utils.Validate.Struct(course); err != nil {
			return utils.BadRequest(err.Error())
		}

		if err := db.Create(&course).Error; err != nil {
			return utils.HandleFindError(err)
		}
		return c.Status(fiber.StatusCreated).JSON(&course)
	}
}

func GetAllCourse(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var course []models.Course

		if err := db.Preload("Subject").Preload("Videos").Preload("Videos.Reviews").Find(&course).Error; err != nil {
			return utils.HandleFindError(err)
		}
		return c.Status(fiber.StatusOK).JSON(&course)
	}
}

func GetCourseById(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var course models.Course
		id := c.Params("course_id")

		if err := db.Preload("Subject").Preload("Videos").Preload("Videos.Reviews").First(&course, &id).Error; err != nil {
			return utils.HandleFindError(err)
		}
		return c.Status(fiber.StatusOK).JSON(&course)
	}
}

func DeleteCourseByID(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var course models.Course
		id := c.Params("course_id")

		if err := db.Delete(&course, &id).Error; err != nil {
			return utils.HandleFindError(err)
		}
		return c.Status(fiber.StatusNoContent).JSON("Deleted")
	}
}

func UpdateCourse(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var course models.Course
		var updateCourseData models.Course
		var subject models.Subject
		id := c.Params("course_id")

		if err := c.BodyParser(&updateCourseData); err != nil {
			return utils.BadRequest(err.Error())
		}

		if err := db.First(&subject, updateCourseData.SubjectID).Error; err != nil {
			return utils.HandleFindError(err)
		}

		if err := db.Preload("Subject").First(&course, &id).Error; err != nil {
			return utils.HandleFindError(err)
		}

		course.Subject = &subject
		course.Course_name = updateCourseData.Course_name
		course.Picture = updateCourseData.Picture
		course.Description = updateCourseData.Description

		if err := db.Save(&course).Error; err != nil {
			return utils.Unexpected(err.Error())
		}

		return c.Status(fiber.StatusOK).JSON(course)
	}
}

func LikeCourse(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get("authorization")
		claims, err := utils.GetClaims(token)
		if err != nil {
			return utils.Unauthorized(err.Error())
		}
		var user models.User
		if err := db.Preload("LikeCourses").First(&user, claims["user_id"]).Error; err != nil {
			return utils.HandleFindError(err)
		}
		var course models.Course
		id := c.Params("course_id")

		if err := db.First(&course, &id).Error; err != nil {
			return utils.HandleFindError(err)
		}
		isLike := checkIsLikeCourse(course, user.LikeCourses)

		tx := db.Begin()
		var resMessage string
		if !isLike {
			course.Like += 1
			user.LikeCourses = append(user.LikeCourses, &course)
			resMessage = "Like Success"
			err = tx.Save(&user).Error
		} else {
			course.Like -= 1
			resMessage = "Unlike Success"
			err = tx.Model(&user).Where("course_course_id = ?", &course.Course_id).Association("LikeCourses").Clear()
		}
		if err != nil {
			tx.Rollback()
			return utils.Unexpected(err.Error())
		}
		if err = tx.Save(&course).Error; err != nil {
			tx.Rollback()
			return utils.Unexpected(err.Error())
		}
		tx.Commit()
		return c.Status(fiber.StatusOK).SendString(resMessage)
	}
}

func IsLikeCourse(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get("authorization")
		claims, err := utils.GetClaims(token)
		if err != nil {
			return utils.Unauthorized(err.Error())
		}
		var user models.User
		if err := db.Preload("LikeCourses").First(&user, claims["user_id"]).Error; err != nil {
			return utils.HandleFindError(err)
		}
		var course models.Course
		if err := db.First(&course, c.Params("course_id")).Error; err != nil {
			return utils.HandleFindError(err)
		}
		userLike := user.LikeCourses
		isLike := checkIsLikeCourse(course, userLike)

		return c.Status(fiber.StatusOK).JSON(isLike)
	}
}

func GetRecommendCourse(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		recCourseIds, err := ai.GetRecommendCourse(c.Params("user_id"))
		if err != nil {
			return utils.Unexpected(err.Error())
		}
		var recCourses []models.Course
		if err := db.Where("course_id IN ?", recCourseIds).Preload("Subject").Preload("Videos").Preload("Videos.Reviews").Find(&recCourses).Error; err != nil {
			return utils.HandleFindError(err)
		}
		for i := 0; i < len(recCourses); i++ {
			rating := 0.0
			count := 0.0
			for j := 0; j < len(*recCourses[i].Videos); j++ {
				videos := *recCourses[i].Videos
				for k := 0; k < len(*videos[j].Reviews); k++ {
					reviews := *videos[j].Reviews
					rating += float64(reviews[k].Rating)
					count++
				}
			}
			recCourses[i].Rating = 0
			if count != 0 {
				recCourses[i].Rating = rating / count
			}
			recCourses[i].Videos = nil
			recCourses[i].Subject = nil
		}
		return c.Status(fiber.StatusOK).JSON(recCourses)
	}
}

func checkIsLikeCourse(course models.Course, userLike []*models.Course) bool {
	for _, u := range userLike {
		if u.Course_id == course.Course_id {
			return true
		}
	}
	return false
}
