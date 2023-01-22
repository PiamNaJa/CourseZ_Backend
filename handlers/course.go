package handlers

//CompileDaemon -command="./CourseZ_Backend"
import (
	"github.com/PiamNaJa/CourseZ_Backend/constants"
	"github.com/PiamNaJa/CourseZ_Backend/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateCourse(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var course *models.Course

		if err := c.BodyParser(&course); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		if err := constants.Validate.Struct(course); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(err.Error())
		}

		if err := db.Model(&models.Course{}).Create(&course).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusCreated).JSON(&course)
	}
}

func GetAllCourse(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var course *[]models.Course

		if err := db.Model(&models.Course{}).Preload("Subject").Preload("Videos").Find(&course).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "record not found",
			})
		}
		return c.Status(fiber.StatusOK).JSON(&course)
	}
}

func GetCourseById(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var course *models.Course
		id := c.Params("course_id")

		if err := db.Model(&models.Course{}).Preload("Subject").Preload("Videos").First(&course, id).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Course not found",
			})
		}
		return c.Status(fiber.StatusOK).JSON(&course)
	}
}

func DeleteCourseByID(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var course *models.Course
		id := c.Params("course_id")

		if err := db.Model(&models.Course{}).Delete(&course, id).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Course not found",
			})
		}
		return c.Status(fiber.StatusOK).JSON("Deleted")
	}
}

func UpdateCourse(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var course *models.Course
		var updateCourseData *models.Course
		id := c.Params("course_id")

		if err := db.Model(&models.Course{}).Preload("Subject").First(&course, id).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Course not found",
			})
		}

		if err := c.BodyParser(&updateCourseData); err != nil {
			return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		course.SubjectID = updateCourseData.SubjectID
		course.Course_name = updateCourseData.Course_name
		course.Picture = updateCourseData.Picture
		course.Description = updateCourseData.Description

		if err := db.Save(&course).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON("Update Success")
	}
}
