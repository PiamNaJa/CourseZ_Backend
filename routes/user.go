package routes

import (
	"github.com/PiamNaJa/CourseZ_Backend/handlers"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func UserRoutes(app fiber.Router, db *gorm.DB) {
	app.Post("/registerstudent", handlers.RegisterStudent(db))
	app.Post("/registerteacher", handlers.RegisterTeacher(db))
	app.Put("/:id", handlers.UpdateStudent(db))
}