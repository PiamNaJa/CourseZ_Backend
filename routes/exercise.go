package routes

import (
	"github.com/PiamNaJa/CourseZ_Backend/handlers"
	m "github.com/PiamNaJa/CourseZ_Backend/middleware"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ExerciseRoutes(app fiber.Router, db *gorm.DB) {
	app.Post("/", m.HaveCourse, m.HaveVideo, handlers.CreateExercise(db))
	app.Get("/", m.HaveCourse, m.HaveVideo, handlers.GetAllExercise(db))
	app.Get("/:id", handlers.GetExerciseById(db))
	app.Delete("/:id", m.HaveCourse, m.HaveVideo, handlers.DeleteExerciseID(db))
	// app.Put("/:id", m.IsLogin, m.IsCourseOwner, handlers.UpdateCourse(db))
}
