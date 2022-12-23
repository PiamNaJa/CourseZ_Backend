package main

import (
	"github.com/PiamNaJa/CourseZ_Backend/configs"
	_ "github.com/PiamNaJa/CourseZ_Backend/docs"
	"github.com/PiamNaJa/CourseZ_Backend/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/goccy/go-json"
)

func main() {
	app := fiber.New(fiber.Config{JSONEncoder: json.Marshal, JSONDecoder: json.Unmarshal})
	configs.Init()
	configs.ConnectDB()
	sql, _ := configs.DB.DB()
	defer sql.Close()

	router := app.Group("/api")

	router.Get("/swagger/*", swagger.HandlerDefault)
	routes.UserRoutes(router.Group("/user"))

	app.Listen(":5000")
}
