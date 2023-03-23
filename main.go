package main

import (
	"fmt"

	"github.com/PiamNaJa/CourseZ_Backend/ai"
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
	defer server.Shutdown(app)
	if err := server.Listen(app); err != nil {
		fmt.Print(err)
	}
	ai.TrainData(configs.DB)
	// fmt.Println(ai.Predict())
}
