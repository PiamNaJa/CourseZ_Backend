package routes

import (
	"github.com/PiamNaJa/CourseZ_Backend/handlers"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RewardRoutes(app fiber.Router, db *gorm.DB) {
	app.Post("/item", handlers.CreateRewardItem(db))
	app.Get("/item", handlers.GetAllRewardItems(db))
	app.Get("/item/:item_id", handlers.GetRewardItemById(db))
	app.Delete("/item/:item_id", handlers.DeleteRewardItemById(db))
	app.Put("/item/:item_id", handlers.UpdateRewardItem(db))
}