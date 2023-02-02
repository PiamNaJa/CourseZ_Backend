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
		var exercise models.Exercise

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

		return c.Status(fiber.StatusCreated).JSON(&exercise)
	}
}

func GetAllExercise(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var exercise []models.Exercise

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
		var exercise models.Exercise

		if err := db.Model(&models.Exercise{}).Where("video_id = ?", c.Params("video_id")).Preload("Choices").First(&exercise, c.Params("exercise_id")).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(&exercise)
	}
}

func DeleteExerciseID(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		r := db.Where("video_id = ? and exercise_id = ?", c.Params("video_id"), c.Params("exercise_id")).Delete(&models.Exercise{})
		if r.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": r.Error.Error(),
			})
		}
		if r.RowsAffected == 0 {
			return c.SendStatus(fiber.StatusNotFound)
		}
		return c.Status(fiber.StatusOK).JSON("Deleted")
	}
}

func UpdateExercise(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var exercise models.Exercise
		var updateExercise models.Exercise

		if err := db.Model(&models.Exercise{}).Where("video_id = ?", c.Params("video_id")).First(&exercise, c.Params("exercise_id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		if err := c.BodyParser(&updateExercise); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		updateExercise.VideoID = exercise.VideoID
		if err := constants.Validate.Struct(updateExercise); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(err.Error())
		}

		exercise.Question = updateExercise.Question
		exercise.Image = updateExercise.Image
		exercise.Choices = updateExercise.Choices

		if err := db.Save(&exercise).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON("Updated")
	}
}
