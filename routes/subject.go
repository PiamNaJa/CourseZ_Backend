package routes

import (
	"github.com/PiamNaJa/CourseZ_Backend/handlers"
	_ "github.com/PiamNaJa/CourseZ_Backend/handlers"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SubjectRoutes(app fiber.Router, db *gorm.DB) {
	app.Post("/CreateSubject", handlers.CreateSubject(db))
	app.Get("/GetAllSubject", handlers.GetAllSubject(db))
}
