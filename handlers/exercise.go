package handlers

import (
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

		if err := db.Where("video_id = ?", c.Params("video_id")).Preload("Choices").Find(&exercise).Error; err != nil {
			return utils.HandleFindError(err)
		}
		return c.Status(fiber.StatusOK).JSON(&exercise)
	}
}

func GetExerciseById(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var exercise models.Exercise

		if err := db.Where("video_id = ?", c.Params("video_id")).Preload("Choices").First(&exercise, c.Params("exercise_id")).Error; err != nil {
			return utils.HandleFindError(err)
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
		if err := db.Where("video_id = ?", c.Params("video_id")).First(&exercise, c.Params("exercise_id")).Error; err != nil {
			return utils.HandleFindError(err)
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

func AddPointAndDoneExercise(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get("authorization")
		claims, err := utils.GetClaims(token)

		if err != nil {
			return utils.Unauthorized(err.Error())
		}

		var point map[string]int32
		if err := c.BodyParser(&point); err != nil {
			return utils.BadRequest(err.Error())
		}

		var user models.User
		if err := db.Where("user_id = ?", claims["user_id"]).First(&user).Error; err != nil {
			return utils.HandleFindError(err)
		}

		var video models.Video
		if err := db.Where("video_id = ?", c.Params("video_id")).First(&video).Error; err != nil {
			return utils.HandleFindError(err)
		}

		user.Point += point["point"]
		user.DoneExercise = append(user.DoneExercise, &video)

		if err := db.Save(&user).Error; err != nil {
			return utils.Unexpected(err.Error())
		}

		return c.Status(fiber.StatusOK).JSON("Point added")
	}
}

func IsDoneExercise(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get("authorization")
		claims, err := utils.GetClaims(token)

		if err != nil {
			return utils.Unauthorized(err.Error())
		}

		var user models.User
		if err := db.Where("user_id = ?", claims["user_id"]).Preload("DoneExercise").First(&user).Error; err != nil {
			return utils.HandleFindError(err)
		}

		videoID, err := c.ParamsInt("video_id")
		if err != nil {
			return utils.BadRequest(err.Error())
		}

		for _, video := range user.DoneExercise {
			if int(video.Video_id) == videoID {
				return c.Status(fiber.StatusOK).JSON(true)
			}
		}
		return c.Status(fiber.StatusOK).JSON(false)
	}
}
