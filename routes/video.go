package routes

import (
	"github.com/PiamNaJa/CourseZ_Backend/handlers"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func VideoRoutes(app fiber.Router, db *gorm.DB) {
	app.Post("/", handlers.CreateVideo(db))
	app.Get("/", handlers.GetAllVideo(db))
	app.Get("/:id", handlers.GetVideoById(db))
	app.Delete("/:id", handlers.DeleteVideoByID(db))
	app.Put("/:id", handlers.UpdateVideo(db))
}
