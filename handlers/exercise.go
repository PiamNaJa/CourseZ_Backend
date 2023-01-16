package handlers

import (
	"strconv"

	"github.com/PiamNaJa/CourseZ_Backend/constants"
	"github.com/PiamNaJa/CourseZ_Backend/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateExercise(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var exercise *models.Exercise

		if err := c.BodyParser(&exercise); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		video_id, err := strconv.ParseInt(c.Params("video_id"), 10, 64)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		exercise.VideoID = int32(video_id)

		if err := constants.Validate.Struct(exercise); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(err.Error())
		}

		if err := db.Model(&models.Exercise{}).Create(&exercise).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusCreated).JSON("Exercise Created")
	}
}

func GetAllExercise(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var exercise *[]models.Exercise

		if err := db.Model(&models.Exercise{}).Where("video_id = ?", c.Params("video_id")).Preload("Choices").Find(&exercise).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(&exercise)
	}
}

func GetExerciseById(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var exercise *models.Exercise

		if err := db.Model(&models.Exercise{}).Where("exercise_id = ?", c.Params("id")).Preload("Choices").First(&exercise).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(&exercise)
	}
}

func DeleteExerciseID(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var exercise *models.Exercise

		if err := db.Model(&models.Exercise{}).Where("video_id = ?", c.Params("video_id")).Delete(&exercise, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON("Deleted")
	}
}

func UpdateExercise(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var exercise *models.Exercise
		if err := c.BodyParser(&exercise); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		if err := constants.Validate.Struct(exercise); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(err)
		}

		if err := db.Model(&models.Exercise{}).Create(&exercise).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(&exercise)
	}
}
