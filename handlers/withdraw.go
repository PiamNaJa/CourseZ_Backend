package handlers

import (
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
		var withdraw models.Withdraw
		if err := c.BodyParser(&withdraw); err != nil {
			return utils.BadRequest(err.Error())
		}
		money := withdraw.Money
		tx := db.Begin()
		var teacher models.UserTeacher
		if err := tx.Where("teacher_id = ?", claims["teacher_id"]).First(&teacher).Error; err != nil {
			return utils.HandleFindError(err)
		}

		teacher.Money -= money

		if teacher.Money < 0 {
			return utils.BadRequest("Not enough money")
		}

		if err := tx.Save(&teacher).Error; err != nil {
			tx.Rollback()
			return utils.Unexpected(err.Error())
		}
		withdraw.TeacherID = teacher.Teacher_id
		if err := tx.Create(&withdraw).Error; err != nil {
			tx.Rollback()
			return utils.Unexpected(err.Error())
		}
		tx.Commit()
		return c.Status(fiber.StatusCreated).JSON(&withdraw)
	}
}
