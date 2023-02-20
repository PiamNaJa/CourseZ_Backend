package middleware

import (
	"os"
	"time"

	"github.com/PiamNaJa/CourseZ_Backend/configs"
	"github.com/PiamNaJa/CourseZ_Backend/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

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

	var course models.Course
	if err := configs.DB.Where("course_id = ?", c.Params("course_id")).First(&course).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Course not found",
		})
	}
	courseOwner := course.TeacherID

	owner := claims["teacher_id"].(float64)
	if owner != float64(courseOwner) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}
	return c.Next()
}
