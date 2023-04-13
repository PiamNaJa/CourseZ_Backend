package routes

import (
	"github.com/PiamNaJa/CourseZ_Backend/handlers"
	m "github.com/PiamNaJa/CourseZ_Backend/middleware"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func PaymentRoutes(app fiber.Router, db *gorm.DB) {
	app.Post("/videos", m.IsLogin, handlers.VideosPayment(db))
	app.Get("/paid/videos", m.IsLogin, handlers.GetPaidVideos(db))
	app.Get("/paid/videos/object", m.IsLogin, handlers.GetPaidVideosObject(db))
	app.Get("/teacher/:teacher_id", m.IsLogin, m.IsTeacher , handlers.GetPaymentTransaction(db))
}
