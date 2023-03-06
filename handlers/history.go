package handlers

import (
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
		var history models.History
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

		var history models.History
		if err := c.BodyParser(&history); err != nil {
			return utils.BadRequest(err.Error())
		}

		history.UserID = int32(claims["user_id"].(float64))
		if err := db.Create(&history).Error; err != nil {
			return utils.Unexpected(err.Error())
		}
		return c.Status(fiber.StatusCreated).JSON("Created")
	}
}
