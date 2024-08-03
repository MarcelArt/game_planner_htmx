package routes

import (
	"github.com/MarcelArt/game_planner_htmx/database"
	"github.com/MarcelArt/game_planner_htmx/handlers"
	"github.com/MarcelArt/game_planner_htmx/repositories"
	"github.com/gofiber/fiber/v2"
)

func SetupProfileRoutes(app *fiber.App) {
	profileHandler := handlers.NewProfileHandler(repositories.NewProfileRepo(database.DB))

	app.Get("/profile", profileHandler.MyProfileView)
	app.Put("/profile", profileHandler.Update)
}
