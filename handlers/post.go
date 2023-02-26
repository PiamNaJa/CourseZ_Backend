package handlers

import (
	"net/url"

	"github.com/PiamNaJa/CourseZ_Backend/constants"
	"github.com/PiamNaJa/CourseZ_Backend/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreatePost(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get("authorization")
		claims, err := constants.GetClaims(token)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}
		var post models.Post

		if err := c.BodyParser(&post); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		var user models.User
		if err := db.Select("user_id", "role").Where("user_id = ?", claims["user_id"]).First(&user).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "User not found",
			})
		}
		post.UserID = user.User_id

		var subject models.Subject
		if err := c.BodyParser(&subject); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		if err := db.Where("subject_title = ? AND class_level = ?", subject.Subject_title, subject.Class_level).First(&subject).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "User not found",
			})
		}
		post.SubjectID = subject.Subject_id
		if err := constants.Validate.Struct(post); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		if err := db.Create(&post).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusCreated).JSON(&post)
	}
}

func GetAllPost(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var posts []models.Post

		if err := db.Preload("Comments.User", func(tx *gorm.DB) *gorm.DB {
			return tx.Omit("Password", "Email")
		}).Preload(clause.Associations, func(tx *gorm.DB) *gorm.DB {
			return tx.Omit("Password", "Email")
		}).Find(&posts).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "record not found",
			})
		}

		return c.Status(fiber.StatusOK).JSON(&posts)
	}
}

func GetPostById(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var post models.Post
		id := c.Params("post_id")

		if err := db.Preload("Comments.User", func(tx *gorm.DB) *gorm.DB {
			return tx.Omit("Password", "Email")
		}).Preload(clause.Associations, func(tx *gorm.DB) *gorm.DB {
			return tx.Omit("Password", "Email")
		}).First(&post, &id).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Post not found",
			})
		}
		return c.Status(fiber.StatusOK).JSON(&post)
	}
}

func GetPostBySubject(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var posts []models.Post

		subject_title, err := url.QueryUnescape(c.Params("subject_title")) // decode url
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
		}
		if err := db.Joins("join subjects on posts.subject_id = subjects.subject_id").Where("subjects.subject_title = ?", &subject_title).Preload("Comments.User", func(tx *gorm.DB) *gorm.DB {
			return tx.Omit("Password", "Email")
		}).Preload(clause.Associations, func(tx *gorm.DB) *gorm.DB {
			return tx.Omit("Password", "Email")
		}).Find(&posts).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(&posts)
	}
}

func GetPostByClassLevel(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var posts []models.Post

		class_level := c.Params("class_level")

		if err := db.Joins("join subjects on posts.subject_id = subjects.subject_id").Where("subjects.class_level = ?", &class_level).Preload("Comments.User", func(tx *gorm.DB) *gorm.DB {
			return tx.Omit("Password", "Email")
		}).Preload(clause.Associations, func(tx *gorm.DB) *gorm.DB {
			return tx.Omit("Password", "Email")
		}).Find(&posts).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "No record",
			})
		}
		return c.Status(fiber.StatusOK).JSON(&posts)
	}
}

func DeletePostByID(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var post models.Post
		id := c.Params("post_id")

		if err := db.Delete(&post, &id).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Post not found",
			})
		}
		return c.Status(fiber.StatusOK).JSON("Deleted")
	}
}

func UpdatePost(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var post models.Post
		var updatePostData models.Post
		id := c.Params("post_id")

		if err := db.First(&post, &id).Error; err != nil {
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

func CreatePostComment(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get("authorization")
		claims, err := constants.GetClaims(token)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}
		var comment models.Comment

		if err := c.BodyParser(&comment); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		var user models.User
		if err := db.Select("user_id", "role").Where("user_id = ?", claims["user_id"]).First(&user).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "User not found",
			})
		}

		var post models.Post
		if err := db.Select("post_id").Where("post_id = ?", c.Params("post_id")).First(&post).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Post not found",
			})
		}
		comment.PostID = post.Post_id
		comment.UserID = user.User_id

		if err := constants.Validate.Struct(comment); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		if err := db.Create(&comment).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusCreated).JSON(&comment)
	}
}
