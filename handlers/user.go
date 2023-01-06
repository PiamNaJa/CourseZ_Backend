package handlers

import (
	"github.com/PiamNaJa/CourseZ_Backend/constants"
	"github.com/PiamNaJa/CourseZ_Backend/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

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
		if err := db.Model(&models.User{}).Create(&user).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		tokenString, err := constants.GenerateToken(user.User_id, user.Role, -1)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		refeshTokenString, err := constants.GenerateRefreshToken(user.User_id, user.Role, -1)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		c.Set("authorization", tokenString)
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"token":        tokenString,
			"refreshToken": refeshTokenString,
		})
	}
}

func RegisterTeacher(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var user *models.User
		var userTeacher *models.UserTeacher
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
		encryptPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
		user.Password = string(encryptPassword)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		tx := db.Begin()
		if err := tx.Model(&models.User{}).Create(&user).Error; err != nil {
			tx.Rollback()
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		userTeacher.UserID = user.User_id

		if err := tx.Model(&models.UserTeacher{}).Create(&userTeacher).Error; err != nil {
			tx.Rollback()
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		tokenString, err := constants.GenerateToken(user.User_id, user.Role, userTeacher.Teacher_id)
		if err != nil {
			tx.Rollback()
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		refeshTokenString, err := constants.GenerateRefreshToken(user.User_id, user.Role, userTeacher.Teacher_id)
		if err != nil {
			tx.Rollback()
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		tx.Commit()
		c.Set("authorization", tokenString)
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"token":        tokenString,
			"refreshToken": refeshTokenString,
		})
	}
}

func LoginUser(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var user *models.User
		var err error
		if err := c.BodyParser(&user); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		password := user.Password

		if err := db.Model(&models.User{}).Preload("Teacher").Where("email = ?", user.Email).First(&user).Error; err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		var tokenString string

		if user.Role == "teacher" || user.Role == "Teacher" {
			tokenString, err = constants.GenerateToken(user.User_id, user.Role, user.Teacher.Teacher_id)
		} else {
			tokenString, err = constants.GenerateToken(user.User_id, user.Role, -1)
		}

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		var refeshTokenString string
		if user.Role == "teacher" || user.Role == "Teacher" {
			refeshTokenString, err = constants.GenerateRefreshToken(user.User_id, user.Role, user.Teacher.Teacher_id)
		} else {
			refeshTokenString, err = constants.GenerateRefreshToken(user.User_id, user.Role, -1)
		}
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		c.Set("authorization", tokenString)
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"token":        tokenString,
			"refreshToken": refeshTokenString,
		})
	}
}

func Update(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var user *models.User
		if err := c.BodyParser(&user); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		if user.Password != "" {
			encryptPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
			user.Password = string(encryptPassword)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": err.Error(),
				})
			}
		}
		if err := db.Model(&models.User{}).Where("user_id = ?", c.Params("id")).Updates(&user).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Update Success",
		})
	}
}
