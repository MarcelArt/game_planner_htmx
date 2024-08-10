package routes

import (
	"github.com/MarcelArt/game_planner_htmx/database"
	"github.com/MarcelArt/game_planner_htmx/handlers"
	"github.com/MarcelArt/game_planner_htmx/repositories"
	"github.com/gofiber/fiber/v2"
)

func SetupItemRoutes(app *fiber.App) {
	itemHandler := handlers.NewItemHandler(repositories.NewItemRepo(database.DB))

	item := app.Group("/item")
	item.Get("/:game_id", itemHandler.ItemsView)
	item.Get("/:game_id/create", itemHandler.CreateView)

	item.Post("/:game_id/create", itemHandler.Create)

	item.Delete("/:id", itemHandler.Delete)
}
