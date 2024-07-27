package routes

import (
	"github.com/MarcelArt/game_planner_htmx/handlers"
	"github.com/MarcelArt/game_planner_htmx/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	app.Use(cors.New())
	app.Use(logger.New())

	app.Static("/scripts", "./views/scripts")

	SetupAuthRoutes(app)

	app.Use(middleware.Auth)
	app.Get("/", handlers.Index)
}
