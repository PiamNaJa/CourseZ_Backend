package handlers

//CompileDaemon -command="./CourseZ_Backend"
import (
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
		id := c.Params("course_id")

		if err := c.BodyParser(&updateCourseData); err != nil {
			return utils.BadRequest(err.Error())
		}

		if err := db.Preload("Subject").First(&course, &id).Error; err != nil {
			return utils.HandleFindError(err)
		}
		course.SubjectID = updateCourseData.SubjectID
		course.Course_name = updateCourseData.Course_name
		course.Picture = updateCourseData.Picture
		course.Description = updateCourseData.Description

		if err := db.Save(&course).Error; err != nil {
			return utils.Unexpected(err.Error())
		}

		return c.Status(fiber.StatusOK).JSON("Update Success")
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

func checkIsLikeCourse(course models.Course, userLike []*models.Course) bool {
	for _, u := range userLike {
		if u.Course_id == course.Course_id {
			return true
		}
	}
	return false
}
