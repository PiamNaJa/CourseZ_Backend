package handlers

import (
	"github.com/PiamNaJa/CourseZ_Backend/constants"
	"github.com/PiamNaJa/CourseZ_Backend/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func WithdrawMoney(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get("authorization")
		if token == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}
		claims, err := constants.GetClaims(token)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}

		var request map[string]interface{}
		if err := c.BodyParser(&request); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		var money int32
		money = int32(request["money"].(float64))

		var teacher models.UserTeacher
		if err := db.Model(&models.UserTeacher{}).Where("teacher_id = ?", claims["teacher_id"]).First(&teacher).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "User not found",
			})
		}

		teacher.Money -= money
		if err := db.Save(&teacher).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		if teacher.Money < 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Not enough money",
			})
		}

		var withdraw models.Withdraw
		withdraw.Recipient = &teacher
		withdraw.Money = money
		withdraw.Text = "Withdraw money"

		if err := db.Model(&models.Withdraw{}).Create(&withdraw).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(&withdraw)
	}
}