package routes

import (
	"github.com/MarcelArt/game_planner_htmx/database"
	"github.com/MarcelArt/game_planner_htmx/handlers"
	"github.com/MarcelArt/game_planner_htmx/middleware"
	"github.com/MarcelArt/game_planner_htmx/repositories"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	app.Use(cors.New())
	app.Use(logger.New())

	app.Static("/scripts", "./views/scripts")
	app.Static("/public", "./public")

	SetupAuthRoutes(app)

	authMiddleware := middleware.NewAuthMiddleware(repositories.NewUserRepo(database.DB), repositories.NewConnectedDeviceRepo(database.DB))
	app.Use(authMiddleware.Auth)
	app.Get("/", handlers.Index)
	SetupProfileRoutes(app)
}
