package routes

import (
	"github.com/PiamNaJa/CourseZ_Backend/handlers"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SearchRoutes(app fiber.Router, db *gorm.DB) {
	app.Get("/", handlers.SearchALL(db))
	app.Get("/course", handlers.SearchCourse(db))
	app.Get("/video", handlers.SearchVideo(db))
	app.Get("/tutor", handlers.SearchTutor(db))
}
