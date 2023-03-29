package handlers

import (
	"time"

	"github.com/PiamNaJa/CourseZ_Backend/models"
	"github.com/PiamNaJa/CourseZ_Backend/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CheckUserVideoHistory(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get("authorization")
		claims, err := utils.GetClaims(token)
		if err != nil {
			return utils.Unauthorized(err.Error())
		}
		var history models.VideoHistory
		if err := db.Where("user_id = ? AND video_id = ?", claims["user_id"], c.Params("video_id")).First(&history).Error; err != nil {
			return utils.HandleFindError(err)
		}
		return c.Status(fiber.StatusOK).JSON(history.Duration)
	}
}

func AddVideoHistory(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get("authorization")
		claims, err := utils.GetClaims(token)
		if err != nil {
			return utils.Unauthorized(err.Error())
		}

		var history models.VideoHistory
		if err := c.BodyParser(&history); err != nil {
			return utils.BadRequest(err.Error())
		}
		videoDuration := history.Duration

		history.UserID = int32(claims["user_id"].(float64))
		result := db.FirstOrCreate(&history, models.VideoHistory{UserID: history.UserID, VideoID: history.VideoID})
		if result.Error != nil {
			return utils.BadRequest(result.Error.Error())
		}
		db.Model(&history).Update("update_at", time.Now().Unix())
		if result.RowsAffected == 0 {
			history.Duration = videoDuration
			if err := db.Save(&history).Error; err != nil {
				return utils.BadRequest(err.Error())
			}
		}
		var CourseHistory models.CourseHistory
		var video models.Video
		if err:= db.Where("video_id = ?", history.VideoID).First(&video).Error; err != nil {
			return utils.BadRequest(err.Error())
		}
		CourseHistory.UserID = history.UserID
		CourseHistory.CourseID = video.CourseID
		result = db.FirstOrCreate(&CourseHistory, models.CourseHistory{UserID: CourseHistory.UserID, CourseID: CourseHistory.CourseID})
		if result.Error != nil {
			return utils.BadRequest(result.Error.Error())
		}
		db.Model(&CourseHistory).Update("update_at", time.Now().Unix())
		return c.Status(fiber.StatusCreated).JSON(history)
	}
}
