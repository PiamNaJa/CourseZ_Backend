package handlers

import (
	"errors"
	"strconv"

	"github.com/PiamNaJa/CourseZ_Backend/models"
	"github.com/PiamNaJa/CourseZ_Backend/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateExercise(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var exercise models.Exercise

		if err := c.BodyParser(&exercise); err != nil {
			return utils.BadRequest(err.Error())
		}

		video_id, err := strconv.ParseInt(c.Params("video_id"), 10, 64)

		if err != nil {
			return utils.Unexpected(err.Error())
		}

		exercise.VideoID = int32(video_id)

		if err := utils.Validate.Struct(exercise); err != nil {
			return utils.BadRequest(err.Error())
		}

		if err := db.Create(&exercise).Error; err != nil {
			return utils.Unexpected(err.Error())
		}

		return c.Status(fiber.StatusCreated).JSON(&exercise)
	}
}

func GetAllExercise(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var exercise []models.Exercise

		err := db.Where("video_id = ?", c.Params("video_id")).Preload("Choices").Find(&exercise).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.NotFound(err.Error())
		}
		if err != nil {
			return utils.Unexpected(err.Error())
		}
		return c.Status(fiber.StatusOK).JSON(&exercise)
	}
}

func GetExerciseById(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var exercise models.Exercise

		err := db.Where("video_id = ?", c.Params("video_id")).Preload("Choices").First(&exercise, c.Params("exercise_id")).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.NotFound(err.Error())
		}
		if err != nil {
			return utils.Unexpected(err.Error())
		}
		return c.Status(fiber.StatusOK).JSON(&exercise)
	}
}

func DeleteExerciseID(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		r := db.Where("video_id = ? and exercise_id = ?", c.Params("video_id"), c.Params("exercise_id")).Delete(&models.Exercise{})
		if r.Error != nil {
			utils.Unexpected(r.Error.Error())
		}
		if r.RowsAffected == 0 {
			return utils.NotFound("record not found")
		}
		return c.Status(fiber.StatusNoContent).JSON("Deleted")
	}
}

func UpdateExercise(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var exercise models.Exercise
		var updateExercise models.Exercise
		
		if err := c.BodyParser(&updateExercise); err != nil {
			return utils.BadRequest(err.Error())
		}
		err := db.Where("video_id = ?", c.Params("video_id")).First(&exercise, c.Params("exercise_id")).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.NotFound(err.Error())
		}
		if err != nil {
			return utils.Unexpected(err.Error())
		}

		updateExercise.VideoID = exercise.VideoID
		if err := utils.Validate.Struct(updateExercise); err != nil {
			return utils.BadRequest(err.Error())
		}

		exercise.Question = updateExercise.Question
		exercise.Image = updateExercise.Image
		exercise.Choices = updateExercise.Choices

		if err := db.Save(&exercise).Error; err != nil {
			return utils.Unexpected(err.Error())
		}
		return c.Status(fiber.StatusOK).JSON("Updated")
	}
}
