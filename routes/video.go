package routes

import (
	"github.com/PiamNaJa/CourseZ_Backend/handlers"
	m "github.com/PiamNaJa/CourseZ_Backend/middleware"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func VideoRoutes(app fiber.Router, db *gorm.DB) {
	app.Post("/", m.IsLogin, m.IsTeacher, handlers.CreateVideo(db))
	app.Get("/", handlers.GetAllVideo(db))
	app.Get("/getby/:class_level", handlers.GetVideoByFilter(db))
  app.Get("/:video_id", handlers.GetVideoById(db))
	app.Delete("/:video_id", m.IsLogin, m.IsTeacher, m.IsVideoOwner, handlers.DeleteVideoByID(db))
	app.Put("/:video_id", m.IsLogin, m.IsTeacher, handlers.UpdateVideo(db))
}
