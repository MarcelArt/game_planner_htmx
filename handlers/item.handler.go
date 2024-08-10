package handlers

import (
	"fmt"
	"strconv"

	"github.com/MarcelArt/game_planner_htmx/models"
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
	gameID := c.Params("game_id")

	items := h.itemRepo.GetByGameID(c, gameID)

	return c.Status(fiber.StatusOK).Render("items", fiber.Map{
		"data":   items,
		"gameId": gameID,
	})
}

func (h *ItemHandler) CreateView(c *fiber.Ctx) error {
	gameID := c.Params("game_id")
	return c.Status(fiber.StatusOK).Render("partials/create_item_modal", fiber.Map{
		"gameId": gameID,
	})
}

func (h *ItemHandler) Create(c *fiber.Ctx) error {
	gameID := c.Params("game_id")
	var itemInput models.Item
	if err := c.BodyParser(&itemInput); err != nil {
		return c.Status(fiber.StatusBadRequest).Render("partials/toast", fiber.Map{"error": err.Error()})
	}

	gameIDInt, err := strconv.Atoi(gameID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).Render("partials/toast", fiber.Map{"error": err.Error()})
	}
	itemInput.GameID = uint(gameIDInt)

	pictureFile, err := c.FormFile("picture")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).Render("partials/toast", fiber.Map{"error": err.Error()})
	}

	if pictureFile != nil {
		c.SaveFile(pictureFile, fmt.Sprintf("./public/uploads/%s", pictureFile.Filename))
		picture := fmt.Sprintf("/public/uploads/%s", pictureFile.Filename)

		itemInput.Picture = picture
	}

	_, err = h.itemRepo.Create(&itemInput)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).Render("partials/toast", fiber.Map{"error": err.Error()})
	}

	items := h.itemRepo.GetByGameID(c, gameID)

	return c.Status(fiber.StatusOK).Render("items", fiber.Map{
		"data":   items,
		"gameId": gameID,
	})
}
