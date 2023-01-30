package routes

import (
	"github.com/PiamNaJa/CourseZ_Backend/handlers"
	m "github.com/PiamNaJa/CourseZ_Backend/middleware"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RewardRoutes(app fiber.Router, db *gorm.DB) {
	app.Post("/item", handlers.CreateRewardItem(db))
	app.Get("/item", handlers.GetAllRewardItems(db))
	app.Get("/item/:item_id", handlers.GetRewardItemById(db))
	app.Delete("/item/:item_id", handlers.DeleteRewardItemById(db))
	app.Put("/item/:item_id", handlers.UpdateRewardItem(db))

	app.Post("/info", handlers.CreateRewardInfo(db))
	app.Get("/info", m.IsLogin, handlers.GetAllRewardInfo(db))
	app.Get("/info/user/:user_id", m.IsLogin, handlers.GetRewardInfoByUser(db)) //wait user refactor m.IsUser middleware
	app.Get("/info/:reward_id", m.IsLogin, handlers.GetRewardInfoByID(db))
	app.Delete("/info/:reward_id", m.IsLogin, handlers.DeleteRewardInfoById(db))
}
