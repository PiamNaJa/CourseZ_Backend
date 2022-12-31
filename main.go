package main

import (
	"github.com/PiamNaJa/CourseZ_Backend/configs"
	_ "github.com/PiamNaJa/CourseZ_Backend/docs"
	"github.com/PiamNaJa/CourseZ_Backend/routes"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
)

func main() {
	app := fiber.New(fiber.Config{JSONEncoder: json.Marshal, JSONDecoder: json.Unmarshal})
	app.Use(recover.New())
	app.Use(cors.New())

	configs.Init()
	configs.ConnectDB()

	sql, _ := configs.DB.DB()
	defer sql.Close()

	router := app.Group("/api")

	router.Get("/swagger/*", swagger.HandlerDefault)
	routes.UserRoutes(router.Group("/user"), configs.DB)
	routes.CourseRoutes(router.Group("/course"), configs.DB)
	routes.SubjectRoutes(router.Group("/subject"), configs.DB)

	app.Listen(":5000")
}
