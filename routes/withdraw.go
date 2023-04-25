package routes

import (
	"github.com/PiamNaJa/CourseZ_Backend/handlers"
	m "github.com/PiamNaJa/CourseZ_Backend/middleware"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func WithdrawRoutes(app fiber.Router, db *gorm.DB) {
	app.Post("/", m.IsLogin, handlers.WithdrawMoney(db))
	app.Get("/teacher/:teacher_id", m.IsLogin, handlers.GetWithdrawTeacher(db))
}
