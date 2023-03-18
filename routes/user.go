package routes

import (
	"github.com/PiamNaJa/CourseZ_Backend/handlers"
	m "github.com/PiamNaJa/CourseZ_Backend/middleware"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func UserRoutes(app fiber.Router, db *gorm.DB) {
	app.Post("/register/student", handlers.RegisterStudent(db))
	app.Post("/register/teacher", handlers.RegisterTeacher(db))
	app.Post("/login", handlers.LoginUser(db))
	app.Post("/newtoken", handlers.GetNewToken(db))
	app.Put("/:user_id", m.IsLogin, m.IsUser, handlers.Update(db))
	app.Get("/teacher", handlers.GetTeacher(db))
	app.Get("/teacher/:teacher_id", handlers.GetTeacherById(db))
	app.Get("/teacher/class/:class_level", handlers.GetTeacherByClassLevel(db))
	app.Get("/:user_id", handlers.GetProfile(db))
	app.Get(":user_id/address", m.IsLogin, m.IsUser, handlers.GetUserAddress(db))
	app.Post(":user_id/address", m.IsLogin, m.IsUser, handlers.CreateUserAddress(db))
	app.Put(":user_id/address", m.IsLogin, m.IsUser, handlers.UpdateUserAddress(db))
}
