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
		if err = db.Select("user_id", "role").Where("user_id = ?", claims["user_id"]).First(&user).Error; err != nil {
			return utils.HandleFindError(err)
		}

		var videos []*models.Video
		for _, video_id := range request["video_id"].([]interface{}) {
			var video *models.Video
			if err := db.Select("video_id", "price").Where("video_id = ?", video_id).First(&video).Error; err != nil {
				return utils.HandleFindError(err)
			}
			videos = append(videos, video)
		}
		user.PaidVideos = append(user.PaidVideos, videos...)
		err = db.Save(&user).Error
		if err != nil {
			return utils.Unexpected(err.Error())
		}

		var courses []models.Course
		if err := db.Preload("Videos", "video_id IN ?", request["video_id"]).Find(&courses).Error; err != nil {
			return utils.Unexpected(err.Error())
		}

		var course models.Course
		var totalPrice int32

		for _, cour := range courses {
			if len(*cour.Videos) != 0 {
				course = cour
				for _, video := range *cour.Videos {
					totalPrice += video.Price
				}
				break
			}
		}
		var teacher models.UserTeacher
		if err := db.First(&teacher, &course.TeacherID).Error; err != nil {
			return utils.HandleFindError(err)
		}

		teacher.Money += totalPrice
		if err := db.Save(&teacher).Error; err != nil {
			return utils.Unexpected(err.Error())
		}

		var userInPayment models.User
		if err := db.Preload("PaidVideos").Where("user_id = ?", &user.User_id).First(&userInPayment).Error; err != nil {
			return utils.HandleFindError(err)
		}

		payment := models.Payment{
			Payee:     &userInPayment,
			Recipient: &teacher,
			Money:     totalPrice,
			Text:      "pay for videos",
		}
		if err := db.Create(&payment).Error; err != nil {
			return utils.Unexpected(err.Error())
		}

		return c.Status(fiber.StatusOK).JSON(payment)
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
