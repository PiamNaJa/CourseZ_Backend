package handlers

import (
	"github.com/PiamNaJa/CourseZ_Backend/constants"
	"github.com/PiamNaJa/CourseZ_Backend/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreatePost(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var post *models.Post

		if err := c.BodyParser(&post); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		if err := constants.Validate.Struct(post); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		if err := db.Model(&models.Post{}).Create(&post).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusCreated).JSON(&post)
	}
}

func GetAllPost(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var post *[]models.Post

		if err := db.Model(&models.Post{}).Preload("Subject").Preload("User").Preload("Comments").Find(&post).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "record not found",
			})
		}
		return c.Status(fiber.StatusOK).JSON(&post)
	}
}

func GetPostById(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var post *models.Post
		id := c.Params("post_id")

		if err := db.Model(&models.Post{}).Preload("Subject").Preload("User").Preload("Comments").First(&post, id).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Post not found",
			})
		}
		return c.Status(fiber.StatusOK).JSON(&post)
	}
}

func GetPostBySubject(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var post *[]models.Post

		if err := db.Model(&models.Post{}).Where("subject_id = ?", c.Query("subject_id")).Find(&post).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "No record",
			})
		}
		return c.Status(fiber.StatusOK).JSON(&post)
	}
}

func GetPostByClassLevel(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var post *[]models.Post

		if err := db.Model(&models.Post{}).Where("class_level = ?", c.Params("class_level")).Find(&post).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "No record",
			})
		}
		return c.Status(fiber.StatusOK).JSON(&post)
	}
}

func DeletePostByID(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var post *models.Post
		id := c.Params("post_id")

		if err := db.Model(&models.Post{}).Delete(&post, id).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Post not found",
			})
		}
		return c.Status(fiber.StatusOK).JSON("Deleted")
	}
}

func UpdatePost(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var post *models.Post
		var updatePostData *models.Post
		id := c.Params("post_id")

		if err := db.Model(&models.Post{}).Preload("Subject").Preload("User").Preload("Comments").First(&post, id).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Post not found",
			})
		}

		if err := c.BodyParser(&updatePostData); err != nil {
			return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		post.Caption = updatePostData.Caption
		post.Post_picture = updatePostData.Post_picture

		if err := db.Save(&post).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON("Update Success")
	}
}
