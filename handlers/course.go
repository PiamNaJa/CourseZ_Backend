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

		if err := db.Model(&models.Course{}).Preload("Subject").First(&course, id).Error; err != nil {
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
		type updateCourse struct {
			SubjectID   int32  `json:"subject_id" gorm:"index;type:int"`
			Course_name string `json:"course_name" gorm:"not null;type:varchar(100)"`
			Picture     string `json:"picture" gorm:"not null;type:varchar(255)"`
			Description string `json:"description" gorm:"not null;type:text"`
		}
		var course *models.Course
		id := c.Params("id")

		if err := db.Model(&models.Course{}).Preload("Subject").First(&course, id).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		var updateCourseData updateCourse

		if err := c.BodyParser(&updateCourseData); err != nil {
			return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		course.SubjectID = updateCourseData.SubjectID
		course.Course_name = updateCourseData.Course_name
		course.Picture = updateCourseData.Picture
		course.Description = updateCourseData.Description

		db.Save(&course)

		return c.Status(fiber.StatusOK).JSON(&course)
	}
}
