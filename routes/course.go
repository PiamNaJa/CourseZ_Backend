package routes

import (
	"github.com/PiamNaJa/CourseZ_Backend/handlers"
	m "github.com/PiamNaJa/CourseZ_Backend/middleware"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CourseRoutes(app fiber.Router, db *gorm.DB) {
	app.Post("/", m.IsLogin, m.IsTeacher, handlers.CreateCourse(db))
	app.Get("/", handlers.GetAllCourse(db))
	app.Get("/:id", handlers.GetCourseById(db))
	app.Delete("/:id", m.IsLogin, m.IsTeacher, m.IsCourseOwner, handlers.DeleteCourseByID(db))
	app.Put("/:id", m.IsLogin, m.IsCourseOwner, handlers.UpdateCourse(db))
}
