package handlers

//CompileDaemon -command="./CourseZ_Backend"
import (
	"github.com/PiamNaJa/CourseZ_Backend/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateVideo(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var video *models.Video

		if err := c.BodyParser(&video); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		if err := db.Model(&models.Video{}).Create(&video).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(&video)
	}
}

func GetAllVideo(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var video *[]models.Video

		if err := db.Model(&models.Video{}).Preload("Reviews").Preload("Exercises").Find(&video).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(&video)
	}
}

// func GetVideoByFilter(db *gorm.DB) fiber.Handler {
// 	return func(c *fiber.Ctx) error {
// 		var video *[]models.Video
// 		flt := c.Params("flt")

// 		if err := db.Model(&models.Video{}).Preload("Course").Find(&video, flt).Error; err != nil {
// 			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 				"error": err.Error(),
// 			})
// 		}
// 		return c.Status(fiber.StatusOK).JSON(&video)
// 	}
// }

func GetVideoById(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var video *models.Video
		id := c.Params("id")

		if err := db.Model(models.Video{}).Preload("Reviews").Preload("Exercises").First(&video, id).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(&video)
	}
}

func DeleteVideoByID(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var video *models.Video
		id := c.Params("id")

		if err := db.Model(models.Video{}).Delete(&video, id).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(&video)
	}
}

func UpdateVideo(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		type updateVideo struct {
			CourseID    int32  `json:"course_id" gorm:"index;type:int"`
			Video_name  string `json:"video_name" gorm:"not null;type:varchar(100)"`
			Price       int32  `json:"price" gorm:"not null;type:int"`
			Picture     string `json:"picture" gorm:"not null;type:varchar(255)"`
			Description string `json:"description" gorm:"not null;type:text"`
			Url         string `json:"url" gorm:"not null;type:varchar(255)"`
		}
		var video *models.Video
		id := c.Params("id")

		if err := db.Model(models.Video{}).Preload("Reviews").Preload("Exercises").First(&video, id).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		var updateVideoData updateVideo

		if err := c.BodyParser(&updateVideoData); err != nil {
			return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		video.CourseID = updateVideoData.CourseID
		video.Video_name = updateVideoData.Video_name
		video.Price = updateVideoData.Price
		video.Picture = updateVideoData.Picture
		video.Description = updateVideoData.Description
		video.Url = updateVideoData.Url

		db.Save(&video)

		return c.Status(fiber.StatusOK).JSON(&video)
	}
}
