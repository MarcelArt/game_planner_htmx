package routes

import (
	"github.com/MarcelArt/game_planner_htmx/database"
	"github.com/MarcelArt/game_planner_htmx/handlers"
	"github.com/MarcelArt/game_planner_htmx/repositories"
	"github.com/gofiber/fiber/v2"
)

func SetupAuthRoutes(app *fiber.App) {
	authHandler := handlers.NewAuthHandler(
		repositories.NewUserRepo(database.DB),
		repositories.NewConnectedDeviceRepo(database.DB),
		repositories.NewProfileRepo(database.DB),
	)

	app.Get("/register", authHandler.RegisterView)
	app.Post("/register", authHandler.Register)

	app.Get("/login", authHandler.LoginView)
	app.Post("/login", authHandler.Login)
}

func SetupAuthRoutesAfterMiddleware(app *fiber.App) {
	authHandler := handlers.NewAuthHandler(
		repositories.NewUserRepo(database.DB),
		repositories.NewConnectedDeviceRepo(database.DB),
		repositories.NewProfileRepo(database.DB),
	)

	app.Get("/logout", authHandler.Logout)
}
