package handlers

import (
	_"github.com/PiamNaJa/CourseZ_Backend/constants"
	"github.com/PiamNaJa/CourseZ_Backend/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// @Sammary Test
// @description This is a sample swagger for Fiber
// @Router /api/user/registerstudent [post]
// @Success 200 {string} string "Hello, World 👋!"
// @Tags: Test
// @Produce json

func RegisterStudent(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var user *models.User
		if err := c.BodyParser(&user); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		encryptPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		user.Password = string(encryptPassword)
		db.Model(models.User{}).Create(&user)
		return c.Status(fiber.StatusOK).JSON(user)
	}
}

func RegisterTeacher(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var user *models.User
		var userTeacher *models.UserTeacher
		var experience *models.Experience
		if err := c.BodyParser(&user); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		if err := c.BodyParser(&userTeacher); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		if err := c.BodyParser(&experience); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		encryptPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		user.Password = string(encryptPassword)

		if err := db.Model(models.Experience{}).Create(&experience).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		if err := db.Model(models.UserTeacher{}).Create(&userTeacher).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		//user.TeacherID = constants.Int32ToSQLNullInt32(userTeacher.Teacher_id)
		user.Teacher = userTeacher
		if err := db.Model(models.User{}).Create(&user).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(user)
	}
}
