package routes

import (
	_ "github.com/PiamNaJa/CourseZ_Backend/handlers"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func UserRoutes(app fiber.Router, db *gorm.DB) {
	//app.Get("/", handlers.UserList)
}
