package routes

import (
	"github.com/PiamNaJa/CourseZ_Backend/handlers"
	_ "github.com/PiamNaJa/CourseZ_Backend/handlers"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CourseRoutes(app fiber.Router, db *gorm.DB) {
	app.Post("/CreateCourse", handlers.CreateCourse(db))
	app.Get("/GetAllCourse", handlers.GetAllCourse(db))
	app.Get("/GetCourse/:id", handlers.GetCourseById(db))
	app.Delete("/DeleteCourse/:id", handlers.DeleteCourseByID(db))
	app.Put("/UpdateCourse/:id", handlers.UpdateCourse(db))
}
