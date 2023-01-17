package routes

import (
	"github.com/PiamNaJa/CourseZ_Backend/handlers"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SearchRoutes(app fiber.Router, db *gorm.DB) {
	app.Get("/:name", handlers.SearchALL(db))
	app.Get("/course/:name", handlers.SearchCourse(db))
	app.Get("/video/:name", handlers.SearchVideo(db))
	app.Get("/tutor/:name", handlers.SearchTutor(db))
}
