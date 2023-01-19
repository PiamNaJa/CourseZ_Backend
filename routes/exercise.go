package routes

import (
	"github.com/PiamNaJa/CourseZ_Backend/handlers"
	m "github.com/PiamNaJa/CourseZ_Backend/middleware"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ExerciseRoutes(app fiber.Router, db *gorm.DB) {
	app.Post("/", m.IsLogin, m.IsTeacher, m.IsCourseOwner, m.IsVideoOwner, handlers.CreateExercise(db))
	app.Get("/", m.IsLogin, m.IsCourseOwner, m.IsVideoOwner, handlers.GetAllExercise(db))
	app.Get("/:exercise_id", handlers.GetExerciseById(db))
	app.Delete("/:exercise_id", m.IsCourseOwner, m.IsVideoOwner, handlers.DeleteExerciseID(db))
	app.Put("/:exercise_id", m.IsLogin, m.IsTeacher, m.IsCourseOwner, m.IsVideoOwner, handlers.UpdateExercise(db))
}
