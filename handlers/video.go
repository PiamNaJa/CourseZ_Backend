package handlers

//CompileDaemon -command="./CourseZ_Backend"
import (
	"github.com/PiamNaJa/CourseZ_Backend/models"
	"github.com/PiamNaJa/CourseZ_Backend/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateVideo(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var video models.Video
		var course models.Course
		if err := c.BodyParser(&video); err != nil {
			return utils.BadRequest(err.Error())
		}

		if err := db.Preload("Subject").Where("course_id = ?", c.Params("course_id")).First(&course).Error; err != nil {
			return utils.HandleFindError(err)
		}
		video.CourseID = course.Course_id
		video.Class_level = course.Subject.Class_level
		if err := utils.Validate.Struct(video); err != nil {
			return utils.BadRequest(err.Error())
		}

		if err := db.Create(&video).Error; err != nil {
			return utils.Unexpected(err.Error())
		}

		return c.Status(fiber.StatusCreated).JSON(&video)
	}
}

func GetAllVideo(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var videos []models.Video

		if err := db.Preload(clause.Associations).Where("course_id = ?", c.Params("course_id")).Find(&videos).Error; err != nil {
			return utils.HandleFindError(err)
		}
		return c.Status(fiber.StatusOK).JSON(&videos)
	}
}

func GetVideoByFilter(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var videos []models.Video

		if err := db.Where("class_level = ?", c.Params("class_level")).Find(&videos).Error; err != nil {
			return utils.HandleFindError(err)
		}
		return c.Status(fiber.StatusOK).JSON(&videos)
	}
}

func GetVideoById(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var video models.Video
		id := c.Params("video_id")

		if err := db.Preload(clause.Associations).Where("course_id = ?", c.Params("course_id")).First(&video, &id).Error; err != nil {
			return utils.HandleFindError(err)
		}
		return c.Status(fiber.StatusOK).JSON(&video)
	}
}

func DeleteVideoByID(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var video models.Video
		id := c.Params("video_id")

		r := db.Where("course_id = ?", c.Params("course_id")).Delete(&video, id)

		if r.Error != nil {
			return utils.HandleFindError(r.Error)
		}
		if r.RowsAffected == 0 {
			return utils.NotFound("record not found")
		}
		return c.Status(fiber.StatusNoContent).JSON("Deleted")
	}
}

func UpdateVideo(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var video models.Video
		var updateVideoData models.Video
		id := c.Params("video_id")

		if err := db.Preload(clause.Associations).Where("course_id = ?", c.Params("course_id")).First(&video, id).Error; err != nil {
			return utils.HandleFindError(err)
		}

		if err := c.BodyParser(&updateVideoData); err != nil {
			return utils.BadRequest(err.Error())
		}

		video.Video_name = updateVideoData.Video_name
		video.Price = updateVideoData.Price
		video.Picture = updateVideoData.Picture
		video.Description = updateVideoData.Description
		video.Url = updateVideoData.Url

		if err := db.Save(&video).Error; err != nil {
			return utils.Unexpected(err.Error())
		}

		return c.Status(fiber.StatusOK).JSON("Updated")
	}
}

func LikeVideo(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get("authorization")
		claims, err := utils.GetClaims(token)
		if err != nil {
			return utils.Unauthorized(err.Error())
		}
		var user models.User
		if err := db.Preload("LikeVideos").First(&user, claims["user_id"]).Error; err != nil {
			return utils.HandleFindError(err)
		}
		var video models.Video
		id := c.Params("video_id")

		if err := db.Where("course_id = ?", c.Params("course_id")).First(&video, &id).Error; err != nil {
			return utils.HandleFindError(err)
		}
		isLike := checkIsLikeVideo(video, user.LikeVideos)
		var resMessage string
		tx := db.Begin()
		if !isLike {
			video.Like += 1
			user.LikeVideos = append(user.LikeVideos, &video)
			if err := tx.Save(&user).Error; err != nil {
				tx.Rollback()
				return utils.Unexpected(err.Error())
			}
			resMessage = "Like Success"
		} else {
			video.Like -= 1
			if err := tx.Model(&user).Where("video_video_id = ?", video.Video_id).Association("LikeVideos").Clear(); err != nil {
				return utils.Unexpected(err.Error())
			}
			resMessage = "Unlike Success"
		}
		if err := tx.Save(&video).Error; err != nil {
			tx.Rollback()
			return utils.Unexpected(err.Error())
		}
		tx.Commit()
		return c.Status(fiber.StatusOK).SendString(resMessage)
	}
}
func IsLikeVideo(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get("authorization")
		claims, err := utils.GetClaims(token)
		if err != nil {
			return utils.Unauthorized(err.Error())
		}
		var user models.User
		if err := db.Preload("LikeVideos").First(&user, claims["user_id"]).Error; err != nil {
			return utils.HandleFindError(err)
		}
		var video models.Video
		id := c.Params("video_id")

		if err := db.Where("course_id = ?", c.Params("course_id")).First(&video, &id).Error; err != nil {
			return utils.HandleFindError(err)
		}
		isLike := checkIsLikeVideo(video, user.LikeVideos)
		return c.Status(fiber.StatusOK).JSON(isLike)
	}
}
func checkIsLikeVideo(video models.Video, userLike []*models.Video) bool {
	for _, u := range userLike {
		if u.Video_id == video.Video_id {
			return true
		}
	}
	return false
}
