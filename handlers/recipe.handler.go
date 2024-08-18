package handlers

import (
	"fmt"
	"strconv"

	"github.com/MarcelArt/game_planner_htmx/models"
	"github.com/MarcelArt/game_planner_htmx/repositories"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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

	dropdownItems, err := h.itemRepo.GetDropdownByGameID(item.GameID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).Render("partials/toast", fiber.Map{"error": err.Error()})
	}

	recipeItemID := uuid.New().String()

	return c.Status(fiber.StatusOK).Render("partials/create_recipe_modal", fiber.Map{
		"item":          item,
		"dropdownItems": dropdownItems,
		"recipeItemId":  recipeItemID,
		"lastIndex":     0,
	})
}

func (h *RecipeHandler) AddRecipeItem(c *fiber.Ctx) error {
	itemID := c.Params("item_id")
	lastIndex := c.Params("last_index")
	i, err := strconv.Atoi(lastIndex)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).Render("partials/toast", fiber.Map{"error": err.Error()})
	}

	item, err := h.itemRepo.GetByID(itemID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).Render("partials/toast", fiber.Map{"error": err.Error()})
	}

	dropdownItems, err := h.itemRepo.GetDropdownByGameID(item.GameID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).Render("partials/toast", fiber.Map{"error": err.Error()})
	}

	recipeItemID := uuid.New().String()
	return c.Status(fiber.StatusOK).Render("partials/recipe_item", fiber.Map{
		"item":          item,
		"dropdownItems": dropdownItems,
		"recipeItemId":  recipeItemID,
		"lastIndex":     i + 1,
	})
}

func (h *RecipeHandler) GetRecipeItemImage(c *fiber.Ctx) error {
	index := c.Params("index", "0")
	itemID := c.Query(fmt.Sprintf("recipeDetails[%s].itemId", index))

	item, err := h.itemRepo.GetByID(itemID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).Render("partials/toast", fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).Render("partials/recipe_item_image", fiber.Map{"image": item.Picture})
}

func (h *RecipeHandler) Create(c *fiber.Ctx) error {
	itemIDStr := c.Params("item_id")
	var recipeInput models.RecipeDto
	if err := c.BodyParser(&recipeInput); err != nil {
		c.Set("HX-Reswap", "innerHTML")
		return c.Status(fiber.StatusBadRequest).Render("partials/toast", fiber.Map{"error": err.Error()})
	}

	itemID, err := strconv.Atoi(itemIDStr)
	if err != nil {
		c.Set("HX-Reswap", "innerHTML")
		return c.Status(fiber.StatusBadRequest).Render("partials/toast", fiber.Map{"error": err.Error()})
	}

	recipeInput.ItemID = uint(itemID)

	if err := h.recipeRepo.CreateWithDetail(&recipeInput); err != nil {
		c.Set("HX-Reswap", "innerHTML")
		return c.Status(fiber.StatusInternalServerError).Render("partials/toast", fiber.Map{"error": err.Error()})
	}

	c.Set("HX-Reswap", "delete")
	return c.SendStatus(fiber.StatusOK)
}
