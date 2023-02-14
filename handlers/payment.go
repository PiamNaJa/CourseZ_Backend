package handlers

import (
	"github.com/PiamNaJa/CourseZ_Backend/constants"
	"github.com/PiamNaJa/CourseZ_Backend/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func VideosPayment(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get("authorization")
		claims, err := constants.GetClaims(token)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}
		var request map[string]interface{}
		if err := c.BodyParser(&request); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		var user models.User
		if err := db.Model(&models.User{}).Select("user_id", "role").Where("user_id = ?", claims["user_id"]).First(&user).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "User not found",
			})
		}
		var videos []*models.Video
		for _, video_id := range request["video_id"].([]interface{}) {
			var video *models.Video
			if err := db.Model(&models.Video{}).Select("video_id", "price").Where("video_id = ?", video_id).First(&video).Error; err != nil {
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
					"error": "Video not found",
				})
			}
			videos = append(videos, video)
		}
		user.PaidVideos = append(user.PaidVideos, videos...)
		if err := db.Save(&user).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		var courses []models.Course
		if err := db.Model(&models.Course{}).Preload("Videos", "video_id IN ?", request["video_id"]).Find(&courses).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
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
		if err := db.Model(&models.UserTeacher{}).First(&teacher, &course.TeacherID).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		teacher.Money += totalPrice
		if err := db.Save(&teacher).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		var userInPayment models.User
		if err := db.Model(&models.User{}).Preload("PaidVideos").Where("user_id = ?", user.User_id).First(&userInPayment).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "User not found",
			})
		}

		var payment models.Payment
		payment.Payee = &userInPayment
		payment.Recipient = &teacher
		payment.Money = totalPrice
		payment.Text = "pay for videos"

		if err := db.Model(&models.Payment{}).Create(&payment).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(payment)
	}
}

func GetPaidVideos(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get("authorization")
		claims, err := constants.GetClaims(token)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}
		var user models.User
		if err := db.Model(&models.User{}).Preload("PaidVideos").Select("user_id", "role").Where("user_id = ?", claims["user_id"]).First(&user).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "User not found",
			})
		}

		var videosId []int32
		for _, video := range user.PaidVideos {
			videosId = append(videosId, video.Video_id)
		}
		if len(videosId) == 0 {
			videosId = []int32{}
		}
		return c.Status(fiber.StatusOK).JSON(&videosId)
	}
}
