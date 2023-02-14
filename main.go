package main

import (
	"log"

	"github.com/PiamNaJa/CourseZ_Backend/configs"
	_ "github.com/PiamNaJa/CourseZ_Backend/docs"
	"github.com/PiamNaJa/CourseZ_Backend/routes"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
)

func main() {
	app := fiber.New(fiber.Config{JSONEncoder: json.Marshal, JSONDecoder: json.Unmarshal})
	app.Use(recover.New())
	app.Use(cors.New(
		cors.Config{
			AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
			AllowOrigins:     "*",
			AllowCredentials: true,
			AllowMethods:     "GET,POST,PUT,DELETE",
		},
	))

	configs.Init()
	configs.ConnectDB()
	configs.WipeData()
	configs.MigrateData()
	configs.SeedDB()
	app.Use(logger.New(logger.Config{
		Format:     "${pid} ${status} - ${method} ${path}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "Asia/Bangkok",
		Output:     configs.File,
	}))
	sql, _ := configs.DB.DB()
	defer sql.Close()
	defer configs.File.Close()

	router := app.Group("/api")

	router.Get("/swagger/*", swagger.HandlerDefault)
	routes.UserRoutes(router.Group("/user"), configs.DB)
	routes.CourseRoutes(router.Group("/course"), configs.DB)
	routes.VideoRoutes(router.Group("/course/:course_id/video"), configs.DB)
	routes.ExerciseRoutes(router.Group("/course/:course_id/video/:video_id/exercise"), configs.DB)
	routes.SubjectRoutes(router.Group("/subject"), configs.DB)
	routes.SearchRoutes(router.Group("/search"), configs.DB)
	routes.ReviewVideoRoutes(router.Group("/video/:video_id/review"), configs.DB)
	routes.ReviewTutorRoutes(router.Group("/teacher/:teacher_id/review"), configs.DB)
	routes.PostRoutes(router.Group("/post"), configs.DB)
	routes.RewardRoutes(router.Group("/reward"), configs.DB)
	routes.PaymentRoutes(router.Group("/payment"), configs.DB)
	routes.WithdrawRoutes(router.Group("/withdraw"), configs.DB)

	log.Fatal(app.Listen(":5000"))
}
