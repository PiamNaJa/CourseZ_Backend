package handlers

import (
	"os"

	"github.com/PiamNaJa/CourseZ_Backend/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
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
		if err := db.Model(models.User{}).Create(&user).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user_id": user.User_id,
			"role":    user.Role,
		})
		tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"token": tokenString,
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
		if err := tx.Model(models.User{}).Create(&user).Error; err != nil {
			tx.Rollback()
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		userTeacher.UserID = user.User_id

		if err := tx.Model(models.UserTeacher{}).Create(&userTeacher).Error; err != nil {
			tx.Rollback()
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		tx.Commit()
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user_id": user.User_id,
			"role":    user.Role,
		})
		tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"token": tokenString,
		})
	}
}

func LoginUser(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var user *models.User

		if err := c.BodyParser(&user); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		password := user.Password

		if err := db.Model(models.User{}).First(&user, user.Email).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user_id": user.User_id,
			"role":    user.Role,
		})
		tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"token": tokenString,
		})
	}
}

func UpdateStudent(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var user *models.User
		if err := c.BodyParser(&user); err != nil {
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
		if err := db.Model(models.User{}).Where("user_id = ?", c.Params("id")).Updates(&user).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Update Success",
		})
	}
}
