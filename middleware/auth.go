package middleware

import (
	"os"
	"strconv"
	"time"

	"github.com/PiamNaJa/CourseZ_Backend/configs"
	"github.com/PiamNaJa/CourseZ_Backend/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func IsLogin(c *fiber.Ctx) error {
	if c.Get("authorization") == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}
	return c.Next()
}

func IsUser(c *fiber.Ctx) error {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(c.Get("authorization"), &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_Secret")), nil
	})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}
	exp := claims["exp"].(float64)
	if int64(exp) < time.Now().Unix() {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Token Expired",
		})
	}
	user_id := claims["user_id"].(float64)

	userUpdateId, err := strconv.ParseFloat(c.Params("id"), 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Something went wrong",
		})
	}
	if userUpdateId != user_id {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}
	return c.Next()
}

func IsTeacher(c *fiber.Ctx) error {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(c.Get("authorization"), &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_Secret")), nil
	})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}
	exp := claims["exp"].(float64)
	if int64(exp) < time.Now().Unix() {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Token Expired",
		})
	}
	role := claims["role"].(string)
	if role != "Teacher" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}
	return c.Next()
}

func IsCourseOwner(c *fiber.Ctx) error {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(c.Get("authorization"), &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_Secret")), nil
	})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}
	exp := claims["exp"].(float64)
	if int64(exp) < time.Now().Unix() {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Token Expired",
		})
	}
	role := claims["role"].(string)
	if role != "Teacher" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}
	var course *models.Course
	configs.DB.Where("course_id = ?", c.Params("id")).First(&course)
	courseOwner := course.TeacherID

	owner := claims["teacher_id"].(float64)
	if owner != float64(courseOwner) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}
	return c.Next()
}
