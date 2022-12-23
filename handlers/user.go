package handlers

import (
	"github.com/PiamNaJa/CourseZ_Backend/repository"
	"github.com/gofiber/fiber/v2"
)

// @Sammary Test
// @description This is a sample swagger for Fiber
// @Router /api/user [get]
// @Success 200 {string} string "Hello, World 👋!"
// @Tags: Test
// @Produce json
func UserList(c *fiber.Ctx) error {
	userHandler := repository.NewUserRepository()
	result, err := userHandler.GetAllUser()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(result)
}
