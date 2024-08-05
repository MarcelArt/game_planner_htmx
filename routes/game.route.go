package routes

import (
	"github.com/MarcelArt/game_planner_htmx/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupGameRoutes(app *fiber.App) {
	gameHandler := handlers.NewGameHandler()

	app.Get("/game", gameHandler.GamesView)
}
