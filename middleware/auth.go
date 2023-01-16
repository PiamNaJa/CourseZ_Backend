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
	if role != "Teacher" && role != "Tutor" {
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

	var course *models.Course
	if err := configs.DB.Model(&models.Course{}).Where("course_id = ?", c.Params("id")).First(&course).Error; err != nil {
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

func IsVideoOwner(c *fiber.Ctx) error {
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

	var video *models.Video
	if err := configs.DB.Model(&models.Video{}).Where("video_id = ?", c.Params("id")).First(&video).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Video not found",
		})
	}

	var course *models.Course
	if err := configs.DB.Model(&models.Course{}).Where("course_id = ?", video.CourseID).First(&course).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Course not found",
		})
	}
	videoOwner := course.TeacherID

	owner := claims["teacher_id"].(float64)
	if owner != float64(videoOwner) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}
	return c.Next()
}
