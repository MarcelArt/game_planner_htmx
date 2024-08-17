package handlers

import (
	"github.com/MarcelArt/game_planner_htmx/repositories"
	"github.com/gofiber/fiber/v2"
)

type RecipeHandler struct {
	recipeRepo repositories.IRecipeRepo
	itemRepo   repositories.IItemRepo
}

func NewRecipeHandler(recipeRepo repositories.IRecipeRepo, itemRepo repositories.IItemRepo) *RecipeHandler {
	return &RecipeHandler{
		recipeRepo: recipeRepo,
		itemRepo:   itemRepo,
	}
}

func (h *RecipeHandler) CreateView(c *fiber.Ctx) error {
	itemID := c.Params("item_id")

	item, err := h.itemRepo.GetByID(itemID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).Render("partials/toast", fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).Render("partials/create_recipe_modal", fiber.Map{
		"item": item,
	})
}
