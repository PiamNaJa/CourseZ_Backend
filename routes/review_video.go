package routes

import (
	"github.com/PiamNaJa/CourseZ_Backend/handlers"
	m "github.com/PiamNaJa/CourseZ_Backend/middleware"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ReviewVideoRoutes(app fiber.Router, db *gorm.DB) {
	app.Post("/", m.IsLogin, handlers.CreateReviewVideo(db))
	app.Get("/", handlers.GetReviewVideoByFilter(db))
}
