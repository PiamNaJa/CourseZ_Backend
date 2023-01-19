package handlers

//CompileDaemon -command="./CourseZ_Backend"
import (
	"github.com/PiamNaJa/CourseZ_Backend/constants"
	"github.com/PiamNaJa/CourseZ_Backend/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"strconv"
)

func CreateVideo(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var video *models.Video
		if err := c.BodyParser(&video); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		course_id, err := strconv.ParseInt(c.Params("course_id"), 10, 64)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		video.CourseID = int32(course_id)

		if err := constants.Validate.Struct(video); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(err.Error())
		}

		if err := db.Model(&models.Video{}).Create(&video).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusCreated).JSON(&video)
	}
}

func GetAllVideo(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var video *[]models.Video

		if err := db.Model(&models.Video{}).Preload("Reviews").Preload("Exercises").Where("course_id = ?", c.Params("course_id")).Find(&video).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "No record",
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
		id := c.Params("video_id")

		if err := db.Model(&models.Video{}).Preload("Reviews").Preload("Exercises").Where("course_id = ?", c.Params("course_id")).First(&video, id).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Video not found",
			})
		}
		return c.Status(fiber.StatusOK).JSON(&video)
	}
}

func DeleteVideoByID(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var video *models.Video
		id := c.Params("video_id")

		r := db.Model(&models.Video{}).Where("course_id = ?", c.Params("course_id")).Delete(&video, id)

		if r.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Video not found",
			})
		}
		if r.RowsAffected == 0 {
			return c.SendStatus(fiber.StatusNotFound)
		}
		return c.Status(fiber.StatusOK).JSON("Deleted")
	}
}

func UpdateVideo(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var video *models.Video
		var updateVideoData *models.Video
		id := c.Params("video_id")

		if err := db.Model(&models.Video{}).Preload("Reviews").Preload("Exercises").Where("course_id = ?", c.Params("course_id")).First(&video, id).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		if err := c.BodyParser(&updateVideoData); err != nil {
			return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		video.Video_name = updateVideoData.Video_name
		video.Price = updateVideoData.Price
		video.Picture = updateVideoData.Picture
		video.Description = updateVideoData.Description
		video.Url = updateVideoData.Url

		if err := db.Save(&video).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON("Updated")
	}
}
