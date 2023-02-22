package handlers

import (
	"fmt"

	"github.com/PiamNaJa/CourseZ_Backend/constants"
	"github.com/PiamNaJa/CourseZ_Backend/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetChat(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var chat models.ChatRoom
		if err := db.Where("inbox_id = ?", c.Params("inbox_id")).Preload("Conversations").First(&chat).Error; err != nil {
			fmt.Println(err)
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"err": "Record not found",
			})
		}
		return c.Status(fiber.StatusOK).JSON(&chat)
	}
}

func NewConversaion(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get("authorization")
		claims, err := constants.GetClaims(token)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}

		var conversation models.Conversation
		if err := c.BodyParser(&conversation); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"err": "Something went wrong",
			})
		}
		conversation.Sender_id = int32(claims["user_id"].(float64))
		var chat models.ChatRoom
		if err := db.Where("inbox_id = ?", c.Params("inbox_id")).Preload("Conversations").First(&chat).Error; err != nil {
			fmt.Println(err)
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"err": "Record not found",
			})
		}

		var inbox models.Inbox
		if err := db.Where("inbox_id = ?", c.Params("inbox_id")).First(&inbox).Error; err != nil {
			fmt.Println(err)
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"err": "Record not found",
			})
		}

		tx := db.Begin()
		conversation.ChatRoom_id = chat.Inbox_id
		if err := tx.Create(&conversation).Error; err != nil {
			tx.Rollback()
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"err": "Something went wrong",
			})
		}
		inbox.Last_message = conversation.Message
		inbox.LastMessageUserID = conversation.Sender_id
		if err := tx.Save(&inbox).Error; err != nil {
			tx.Rollback()
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"err": "Something went wrong",
			})
		}
		tx.Model(&chat).Association("Conversations").Append(&conversation)
		if err := tx.Save(&chat).Error; err != nil {
			tx.Rollback()
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"err": "Something went wrong",
			})
		}
		tx.Commit()
		return c.Status(fiber.StatusCreated).JSON(&chat)
	}
}

func GetInbox(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get("authorization")
		claims, err := constants.GetClaims(token)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}
		var inbox []models.Inbox
		if err := db.Where("user1_id = ? OR user2_id = ?", claims["user_id"], claims["user_id"]).Preload("User1", func(db *gorm.DB) *gorm.DB {
			return db.Select("user_id, nickname, picture")
		}).Preload("User2", func(db *gorm.DB) *gorm.DB {
			return db.Select("user_id, nickname, picture")
		}).Find(&inbox).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"err": "Record not found",
			})
		}
		return c.Status(fiber.StatusOK).JSON(&inbox)
	}
}