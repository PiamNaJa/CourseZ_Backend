package handlers

import (
	"errors"

	"github.com/PiamNaJa/CourseZ_Backend/models"
	"github.com/PiamNaJa/CourseZ_Backend/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func WithdrawMoney(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get("authorization")
		claims, err := utils.GetClaims(token)
		if err != nil {
			return utils.Unauthorized(err.Error())
		}

		var request map[string]interface{}
		if err := c.BodyParser(&request); err != nil {
			return utils.BadRequest(err.Error())
		}

		var money int32 = int32(request["money"].(float64))
		tx := db.Begin()
		var teacher models.UserTeacher
		if err := tx.Where("teacher_id = ?", claims["teacher_id"]).First(&teacher).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return utils.NotFound(err.Error())
			}
			return utils.Unexpected(err.Error())
		}

		teacher.Money -= money

		if teacher.Money < 0 {
			return utils.BadRequest("Not enough money")
		}

		if err := tx.Save(&teacher).Error; err != nil {
			tx.Rollback()
			return utils.Unexpected(err.Error())
		}

		var withdraw models.Withdraw
		withdraw.Recipient = &teacher
		withdraw.Money = money
		withdraw.Text = "Withdraw money"

		if err := tx.Create(&withdraw).Error; err != nil {
			tx.Rollback()
			return utils.Unexpected(err.Error())
		}
		tx.Commit()
		return c.Status(fiber.StatusOK).JSON(&withdraw)
	}
}
