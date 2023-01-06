package handlers

//CompileDaemon -command="./CourseZ_Backend"
import (
	"github.com/PiamNaJa/CourseZ_Backend/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateCourse(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var course *models.Course

		if err := c.BodyParser(&course); err != nil {
			return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		if err := db.Model(&models.Course{}).Create(&course).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(&course)
	}
}

func GetAllCourse(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var course *[]models.Course

		if err := db.Model(&models.Course{}).Preload("Subject").Find(&course).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(&course)
	}
}

func GetCourseById(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var course *models.Course
		id := c.Params("id")

		if err := db.Model(&models.Course{}).Preload("Subject").Preload("Videos").First(&course, id).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(&course)
	}
}

func DeleteCourseByID(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var course *models.Course
		id := c.Params("id")

		if err := db.Model(&models.Course{}).Delete(&course, id).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(&course)
	}
}

func UpdateCourse(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var course *models.Course
		var updateCourseData *models.Course
		id := c.Params("id")

		if err := db.Model(&models.Course{}).Preload("Subject").First(&course, id).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
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
