package main

import (
	"fmt"

	"github.com/PiamNaJa/CourseZ_Backend/configs"
	"github.com/PiamNaJa/CourseZ_Backend/server"
)

func main() {
	configs.Init()
	configs.ConnectDB()
	app := server.Create()
	server.SetupRoute(app)
	server.WipeAndSeedDatabaseData()
	configs.RandomData()
	server.PrintRoutes(app)
	cron := server.CreateCron()
	server.StartCron(cron)
	defer server.ShutdownCron(cron)
	defer server.Shutdown(app)
	if err := server.Listen(app); err != nil {
		fmt.Println(err)
	}
}
