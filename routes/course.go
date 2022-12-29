package routes

import (
	"github.com/PiamNaJa/CourseZ_Backend/handlers"
	_ "github.com/PiamNaJa/CourseZ_Backend/handlers"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CourseRoutes(app fiber.Router, db *gorm.DB) {
	app.Post("/", handlers.CreateCourse(db))
	app.Get("/", handlers.GetAllCourse(db))
	app.Get("/:id", handlers.GetCourseById(db))
	app.Delete("/:id", handlers.DeleteCourseByID(db))
	app.Put("/:id", handlers.UpdateCourse(db))
}
