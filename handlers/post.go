package handlers

import (
	"errors"
	"net/url"

	"github.com/PiamNaJa/CourseZ_Backend/models"
	"github.com/PiamNaJa/CourseZ_Backend/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreatePost(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get("authorization")
		claims, err := utils.GetClaims(token)
    
		if err != nil {
			return utils.Unauthorized(err.Error())
		}
		var post models.Post

		if err := c.BodyParser(&post); err != nil {
			return utils.BadRequest(err.Error())
		}
		var user models.User
		err = db.Select("user_id", "role").Where("user_id = ?", claims["user_id"]).First(&user).Error

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.NotFound(err.Error())
		}

		if err != nil {
			return utils.Unexpected(err.Error())
		}
		post.UserID = user.User_id

		var subject models.Subject
		if err := c.BodyParser(&subject); err != nil {
			return utils.BadRequest(err.Error())
		}

		err = db.Where("subject_title = ? AND class_level = ?", subject.Subject_title, subject.Class_level).First(&subject).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.NotFound(err.Error())
		}
		if err != nil {
			return utils.Unexpected(err.Error())
		}

		post.SubjectID = subject.Subject_id
		if err := utils.Validate.Struct(post); err != nil {
			return utils.BadRequest(err.Error())
		}

		if err := db.Create(&post).Error; err != nil {
			return utils.Unexpected(err.Error())
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
			return utils.Unexpected(err.Error())
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
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return utils.NotFound(err.Error())
			}
			return utils.Unexpected(err.Error())
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
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return utils.NotFound(err.Error())
			}
			return utils.Unexpected(err.Error())
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
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return utils.NotFound(err.Error())
			}
			return utils.Unexpected(err.Error())
		}
		return c.Status(fiber.StatusOK).JSON(&posts)
	}
}

func DeletePostByID(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var post models.Post
		id := c.Params("post_id")

		if err := db.Delete(&post, &id).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return utils.NotFound(err.Error())
			}
			return utils.Unexpected(err.Error())
		}
		return c.Status(fiber.StatusNoContent).JSON("Deleted")
	}
}

func UpdatePost(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var post models.Post
		var updatePostData models.Post
		id := c.Params("post_id")

		if err := db.First(&post, &id).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return utils.NotFound(err.Error())
			}
			return utils.Unexpected(err.Error())
		}

		if err := c.BodyParser(&updatePostData); err != nil {
			return utils.BadRequest(err.Error())
		}

		post.Caption = updatePostData.Caption
		post.Post_picture = updatePostData.Post_picture

		if err := db.Save(&post).Error; err != nil {
			return utils.Unexpected(err.Error())
		}

		return c.Status(fiber.StatusOK).JSON("Update Success")
	}
}

func CreatePostComment(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get("authorization")
		claims, err := utils.GetClaims(token)
		if err != nil {
			return utils.Unauthorized(err.Error())
		}
		var comment models.Comment

		if err := c.BodyParser(&comment); err != nil {
			return utils.BadRequest(err.Error())
		}
		var user models.User
		if err := db.Select("user_id", "role").Where("user_id = ?", claims["user_id"]).First(&user).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return utils.NotFound(err.Error())
			}
			return utils.Unexpected(err.Error())
		}

		var post models.Post
		if err := db.Select("post_id").Where("post_id = ?", c.Params("post_id")).First(&post).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return utils.NotFound(err.Error())
			}
			return utils.Unexpected(err.Error())
		}
		comment.PostID = post.Post_id
		comment.UserID = user.User_id

		if err := utils.Validate.Struct(comment); err != nil {
			return utils.BadRequest(err.Error())
		}
		if err := db.Create(&comment).Error; err != nil {
			return utils.Unexpected(err.Error())
		}

		return c.Status(fiber.StatusCreated).JSON(&comment)
	}
}
