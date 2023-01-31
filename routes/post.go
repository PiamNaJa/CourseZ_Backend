package routes

import (
	"github.com/PiamNaJa/CourseZ_Backend/handlers"
	m "github.com/PiamNaJa/CourseZ_Backend/middleware"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func PostRoutes(app fiber.Router, db *gorm.DB) {
	app.Post("/", m.IsLogin, handlers.CreatePost(db))
	app.Get("/", handlers.GetAllPost(db))
	app.Get("/:post_id", handlers.GetPostById(db))
	app.Get("/subject/:subject_title", handlers.GetPostBySubject(db))
	app.Get("/class/:class_level", handlers.GetPostByClassLevel(db))
	app.Put("/:post_id", m.IsLogin, m.IsPostOwner, handlers.UpdatePost(db))
	app.Delete("/:post_id", m.IsLogin, m.IsPostOwner, handlers.DeletePostByID(db))
}
