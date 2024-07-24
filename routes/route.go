package routes

import (
	"github.com/MarcelArt/game_planner_htmx/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func SetupRoutes(app *fiber.App) {
	app.Use(cors.New())

	app.Static("/scripts", "./views/scripts")
	// app.Static("/src", "./views/src")
	app.Get("/", handlers.Index)
}
