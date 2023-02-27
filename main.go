package main

import (
	"fmt"

	"github.com/PiamNaJa/CourseZ_Backend/server"
)

func main() {
	app := server.Create()
	server.SetupRoute(app)
	server.WipeAndSeedDatabaseData()
	server.PrintRoutes(app)
	defer server.Shutdown(app)
	if err := server.Listen(app); err != nil {
		fmt.Print(err)
	}
}
