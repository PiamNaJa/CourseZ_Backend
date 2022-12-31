package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// @Sammary Test
// @description This is a sample swagger for Fiber
// @Router /api/user [get]
// @Success 200 {string} string "Hello, World 👋!"
// @Tags: Test
// @Produce json

func Register(c *fiber.Ctx, db *gorm.DB) error {
	//test := repository.NewUserRepository(db)
	//test.CreateUser()
	return c.SendString("Hello, World 👋!")
}