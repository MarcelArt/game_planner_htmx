package handlers

import (
	"github.com/MarcelArt/game_planner_htmx/repositories"
	"github.com/gofiber/fiber/v2"
)

type ItemHandler struct {
	itemRepo repositories.IItemRepo
}

func NewItemHandler(itemRepo repositories.IItemRepo) *ItemHandler {
	return &ItemHandler{
		itemRepo: itemRepo,
	}
}

func (h *ItemHandler) ItemsView(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).Render("items", nil)
}

func (h *ItemHandler) CreateView(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).Render("partials/create_item_modal", nil)
}
