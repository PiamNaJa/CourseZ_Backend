package routes

import (
	"github.com/PiamNaJa/CourseZ_Backend/handlers"
	m "github.com/PiamNaJa/CourseZ_Backend/middleware"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func InboxRoutes(app fiber.Router, db *gorm.DB) {
	app.Get("/", m.IsLogin, handlers.GetInbox(db))
	app.Get("/:inbox_id", m.IsLogin, handlers.GetChat(db))
	app.Post("/:inbox_id", m.IsLogin, handlers.NewConversaion(db))
}
