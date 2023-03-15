package main

import (
	"github.com/PiamNaJa/CourseZ_Backend/ai"
	"github.com/PiamNaJa/CourseZ_Backend/configs"
)

func main() {
	configs.Init()
	configs.ConnectDB()
	// app := server.Create()
	// server.SetupRoute(app)
	// server.WipeAndSeedDatabaseData()
	// server.PrintRoutes(app)
	// defer server.Shutdown(app)
	// if err := server.Listen(app); err != nil {
	// 	fmt.Print(err)
	// }
	ai.K_cluster(configs.DB, 3)
}
