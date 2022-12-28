package routes

import (
	"fmt"

	"github.com/PiamNaJa/CourseZ_Backend/handlers"
	_ "github.com/PiamNaJa/CourseZ_Backend/handlers"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CourseRoutes(app fiber.Router, db *gorm.DB) {
	app.Post("/CreateCourse", handlers.CreateCourse(db))
	app.Get("/GetAll", handlers.GetAll(db))
	app.Get("/GetById/:id", func(c *fiber.Ctx) error {
		fmt.Fprintf(c, "%s\n", c.Params("id"))
		return nil
	})
}
