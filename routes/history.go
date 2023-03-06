package routes

import (
	"github.com/PiamNaJa/CourseZ_Backend/handlers"
	m "github.com/PiamNaJa/CourseZ_Backend/middleware"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func HistoryRoutes(app fiber.Router, db *gorm.DB) {
	app.Get("/:video_id", m.IsLogin, handlers.CheckUserVideoHistory(db))
	app.Post("/", m.IsLogin, handlers.AddVideoHistory(db))
}
