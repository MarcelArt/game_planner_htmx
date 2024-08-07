package routes

import (
	"github.com/MarcelArt/game_planner_htmx/database"
	"github.com/MarcelArt/game_planner_htmx/handlers"
	"github.com/MarcelArt/game_planner_htmx/repositories"
	"github.com/gofiber/fiber/v2"
)

func SetupGameRoutes(app *fiber.App) {
	gameHandler := handlers.NewGameHandler(repositories.NewGameRepo(database.DB))

	game := app.Group("/game")
	game.Get("/", gameHandler.GamesView)
	game.Get("/create", gameHandler.CreateGameView)
	game.Get("/created", gameHandler.MyCreatedGamesView)
	game.Get("/created/detail", gameHandler.CreatedGameDetailView)
	game.Post("/create", gameHandler.CreateGame)
}
