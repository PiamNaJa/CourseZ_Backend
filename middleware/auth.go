package middleware

import (
	"os"
	"strconv"
	"time"

	"github.com/PiamNaJa/CourseZ_Backend/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func IsLogin(c *fiber.Ctx) error {
	if c.Get("authorization") == "" {
		return utils.Unauthorized("Unauthorized")
	}
	return c.Next()
}

func IsUser(c *fiber.Ctx) error {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(c.Get("authorization"), &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_Secret")), nil
	})
	if err != nil {
		return utils.Unauthorized(err.Error())
	}
	exp := claims["exp"].(float64)
	if int64(exp) < time.Now().Unix() {
		return utils.Unauthorized("Token Expired")
	}
	user_id := claims["user_id"].(float64)

	userUpdateId, err := strconv.ParseFloat(c.Params("user_id"), 64)
	if err != nil {
		return utils.BadRequest(err.Error())
	}
	if userUpdateId != user_id {
		return utils.Unauthorized("Unauthorized")
	}
	return c.Next()
}

func IsTeacher(c *fiber.Ctx) error {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(c.Get("authorization"), &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_Secret")), nil
	})
	if err != nil {
		return utils.Unauthorized(err.Error())
	}
	exp := claims["exp"].(float64)
	if int64(exp) < time.Now().Unix() {
		return utils.Unauthorized("Token Expired")
	}
	role := claims["role"].(string)
	if role != "Teacher" && role != "Tutor" {
		return utils.Unauthorized("Unauthorized")
	}
	return c.Next()
}
