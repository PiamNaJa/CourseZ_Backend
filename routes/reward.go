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

	app.Post("/info", m.IsLogin, handlers.CreateRewardInfo(db))
	app.Get("/info", handlers.GetAllRewardInfo(db))
	app.Get("/info/user/:user_id", m.IsLogin, m.IsUser, handlers.GetRewardInfoByUser(db))
	app.Get("/info/:reward_id", m.IsLogin, m.IsRewardOwner, handlers.GetRewardInfoByID(db))
	app.Delete("/info/:reward_id", m.IsLogin, m.IsRewardOwner, handlers.DeleteRewardInfoById(db))
}
