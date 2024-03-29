package server

import (
	"fmt"
	"os"

	"github.com/PiamNaJa/CourseZ_Backend/ai"
	"github.com/PiamNaJa/CourseZ_Backend/configs"
	"github.com/PiamNaJa/CourseZ_Backend/routes"
	"github.com/PiamNaJa/CourseZ_Backend/socket"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/websocket/v2"
	"github.com/robfig/cron/v3"
)

func Create() *fiber.App {
	app := fiber.New(fiber.Config{JSONEncoder: json.Marshal, JSONDecoder: json.Unmarshal, ErrorHandler: configs.CustomErrorHandler})

	app.Use(recover.New())
	app.Use(configs.CustomCors())
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})
	return app
}

func CreateCron() *cron.Cron {
	c := cron.New()
	c.AddFunc("@midnight", func() { ai.TrainData(configs.DB) })
	return c
}

func StartCron(c *cron.Cron) {
	c.Start()
}

func ShutdownCron(c *cron.Cron) {
	c.Stop()
}
func Listen(app *fiber.App) error {

	port := os.Getenv("PORT")
	if port == "" {
		port = "6000"
	}

	return app.Listen("0.0.0.0:" + port)
}

func SetupRoute(app *fiber.App) {
	app.Get("/metrics", monitor.New())
	app.Get("/ws/:id", websocket.New(socket.NewServer().HandleConnection))
	router := app.Group("/api")
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
	routes.InboxRoutes(router.Group("/inbox"), configs.DB)
	routes.WithdrawRoutes(router.Group("/withdraw"), configs.DB)
	routes.HistoryRoutes(router.Group("/history"), configs.DB)
}

func Shutdown(app *fiber.App) {
	sql, _ := configs.DB.DB()
	sql.Close()
	app.Shutdown()
}

func WipeAndSeedDatabaseData() {
	configs.WipeData()
	configs.MigrateData()
	configs.TestChat()
	configs.SeedDB()
}

func PrintRoutes(app *fiber.App) {
	data, _ := json.MarshalIndent(app.GetRoutes(true), "", "  ")
	var result []map[string]interface{}
	json.Unmarshal(data, &result)
	for _, route := range result {
		fmt.Println(route["method"], route["path"])
	}
}
