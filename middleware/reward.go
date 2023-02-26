package middleware

import (
	"errors"
	"os"
	"time"

	"github.com/PiamNaJa/CourseZ_Backend/configs"
	"github.com/PiamNaJa/CourseZ_Backend/models"
	"github.com/PiamNaJa/CourseZ_Backend/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

func IsRewardOwner(c *fiber.Ctx) error {
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

	var rewardInfo models.Reward_Info
	if err := configs.DB.Where("reward_id = ?", c.Params("reward_id")).First(&rewardInfo).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.NotFound(err.Error())
		}
		return utils.Unexpected(err.Error())
	}

	rewardOwner := rewardInfo.UserID

	owner := claims["user_id"].(float64)
	if owner != float64(rewardOwner) {
		return utils.Unauthorized("You are not the owner of this reward")
	}
	return c.Next()
}
