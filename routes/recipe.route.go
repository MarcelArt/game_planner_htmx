package routes

import (
	"github.com/MarcelArt/game_planner_htmx/database"
	"github.com/MarcelArt/game_planner_htmx/handlers"
	"github.com/MarcelArt/game_planner_htmx/repositories"
	"github.com/gofiber/fiber/v2"
)

func SetupRecipeRoutes(app *fiber.App) {
	recipeHandler := handlers.NewRecipeHandler(
		repositories.NewRecipeRepo(database.DB),
		repositories.NewItemRepo(database.DB),
	)

	recipe := app.Group("/recipe")
	recipe.Get("/:item_id/create", recipeHandler.CreateView)
	recipe.Get("/:item_id/add", recipeHandler.AddRecipeItem)
}
