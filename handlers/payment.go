package handlers

import (
	"github.com/PiamNaJa/CourseZ_Backend/models"
	"github.com/PiamNaJa/CourseZ_Backend/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func VideosPayment(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get("authorization")
		claims, err := utils.GetClaims(token)
		if err != nil {
			return utils.Unauthorized(err.Error())
		}
		var request map[string]interface{}
		if err := c.BodyParser(&request); err != nil {
			return utils.BadRequest(err.Error())
		}

		var user models.User
		if err = db.Where("user_id = ?", claims["user_id"]).First(&user).Error; err != nil {
			return utils.HandleFindError(err)
		}

		var totalPrice int32 = 0
		var videos []*models.Video
		for _, video_id := range request["video_id"].([]interface{}) {
			var video *models.Video
			if err := db.Select("video_id", "price").Where("video_id = ?", video_id).First(&video).Error; err != nil {
				return utils.HandleFindError(err)
			}
			videos = append(videos, video)
			totalPrice += video.Price
		}
		user.PaidVideos = append(user.PaidVideos, videos...)
		err = db.Save(&user).Error
		if err != nil {
			return utils.Unexpected(err.Error())
		}

		var course models.Course
		if err := db.Preload("Videos", "video_id IN ?", request["video_id"]).First(&course).Error; err != nil {
			return utils.Unexpected(err.Error())
		}

		var teacher models.UserTeacher
		if err := db.First(&teacher, &course.TeacherID).Error; err != nil {
			return utils.HandleFindError(err)
		}

		teacher.Money += totalPrice
		if err := db.Save(&teacher).Error; err != nil {
			return utils.Unexpected(err.Error())
		}
		var payments []models.Payment
		for _, video := range videos {
			payment := models.Payment{
				Payee:     &user,
				Recipient: &teacher,
				Money:     video.Price,
				Text:      "pay for videos",
				VideoID:   video.Video_id,
			}
			payments = append(payments, payment)
		}
		if err := db.Create(&payments).Error; err != nil {
			return utils.Unexpected(err.Error())
		}

		return c.Status(fiber.StatusOK).JSON("Success")
	}
}

func GetPaidVideos(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get("authorization")
		claims, err := utils.GetClaims(token)
		if err != nil {
			return utils.Unauthorized(err.Error())
		}
		var user models.User
		if err = db.Preload("PaidVideos").Select("user_id", "role").Where("user_id = ?", claims["user_id"]).First(&user).Error; err != nil {
			return utils.HandleFindError(err)
		}

		var videosId []int32 = []int32{}
		for _, video := range user.PaidVideos {
			videosId = append(videosId, video.Video_id)
		}

		return c.Status(fiber.StatusOK).JSON(&videosId)
	}
}

func GetPaymentTransaction(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get("authorization")
		claims, err := utils.GetClaims(token)
		if err != nil {
			return utils.Unauthorized(err.Error())
		}
		var payment []models.Payment
		if err := db.Where("teacher_id = ?", claims["teacher_id"]).Preload("Video", func (tx *gorm.DB) *gorm.DB {
			return tx.Select("video_id", "video_name", "picture")
		}).Find(&payment).Error; err != nil {
			return utils.Unexpected(err.Error())
		}
		return c.Status(fiber.StatusOK).JSON(payment)
	}
}
