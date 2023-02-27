package middleware

import (
	"errors"
	"time"

	"github.com/PiamNaJa/CourseZ_Backend/configs"
	"github.com/PiamNaJa/CourseZ_Backend/models"
	"github.com/PiamNaJa/CourseZ_Backend/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func IsPostOwner(c *fiber.Ctx) error {
	claims, err := utils.GetClaims(c.Get("authorization"))
	if err != nil {
		return utils.Unauthorized(err.Error())
	}
	exp := claims["exp"].(float64)
	if int64(exp) < time.Now().Unix() {
		return utils.Unauthorized("Token Expired")
	}

	var post models.Post
	if err := configs.DB.Where("post_id = ?", c.Params("post_id")).First(&post).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.NotFound(err.Error())
		}
		return utils.Unexpected(err.Error())
	}
	postOwner := post.UserID

	owner := claims["user_id"].(float64)
	if owner != float64(postOwner) {
		return utils.Unauthorized("You are not the owner of this post")
	}
	return c.Next()
}
