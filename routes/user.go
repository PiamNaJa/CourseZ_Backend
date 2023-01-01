package routes

import (
	"github.com/PiamNaJa/CourseZ_Backend/handlers"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func UserRoutes(app fiber.Router, db *gorm.DB) {
	app.Post("/register/student", handlers.RegisterStudent(db))
	app.Post("/register/teacher", handlers.RegisterTeacher(db))
	app.Post("/login", handlers.LoginUser(db))
	app.Put("/:id", handlers.Update(db))
}
