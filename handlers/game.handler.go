package handlers

import (
	"fmt"

	"github.com/MarcelArt/game_planner_htmx/middleware"
	"github.com/MarcelArt/game_planner_htmx/models"
	"github.com/MarcelArt/game_planner_htmx/repositories"
	"github.com/gofiber/fiber/v2"
)

type GameHandler struct {
	gameRepo repositories.IGameRepo
}

func NewGameHandler(gameRepo repositories.IGameRepo) *GameHandler {
	return &GameHandler{
		gameRepo: gameRepo,
	}
}

func (h *GameHandler) GamesView(c *fiber.Ctx) error {
	var gamesModel []models.Game
	games := h.gameRepo.Read(c, gamesModel)
	return c.Render("games", fiber.Map{
		"data": games,
		"page": games.Page + 1,
		"prev": games.Page,
		"next": games.Page + 1,
	})
}

func (h *GameHandler) CreateGameView(c *fiber.Ctx) error {
	return c.Render("create_game", nil)
}

func (h *GameHandler) CreateGame(c *fiber.Ctx) error {
	var gameInput models.Game
	if err := c.BodyParser(&gameInput); err != nil {
		return c.Status(fiber.StatusInternalServerError).Render("partials/toast", fiber.Map{"error": err.Error()})
	}

	pictureFile, err := c.FormFile("picture")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).Render("partials/toast", fiber.Map{"error": err.Error()})
	}
	c.SaveFile(pictureFile, fmt.Sprintf("./public/%s", pictureFile.Filename))
	picture := fmt.Sprintf("/public/%s", pictureFile.Filename)

	profile, err := middleware.GetCurrentProfile(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).Render("partials/toast", fiber.Map{"error": err.Error()})
	}

	gameInput.Picture = picture
	gameInput.ProfileID = profile.ID

	_, err = h.gameRepo.Create(&gameInput)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).Render("partials/toast", fiber.Map{"error": err.Error()})
	}

	var gamesModel []models.Game
	games := h.gameRepo.Read(c, gamesModel)
	return c.Render("games", fiber.Map{
		"data": games,
		"page": games.Page + 1,
		"prev": games.Page,
		"next": games.Page + 1,
	})
}

func (h *GameHandler) MyCreatedGamesView(c *fiber.Ctx) error {
	profile, err := middleware.GetCurrentProfile(c)
	if err != nil {
		c.Set("HX-Redirect", "/login")
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	games := h.gameRepo.GetByProfileID(c, profile.ID)

	return c.Render("created_games", fiber.Map{
		"data": games,
		"page": games.Page + 1,
		"prev": games.Page,
		"next": games.Page + 1,
	})
}

func (h *GameHandler) CreatedGameDetailView(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.Render("created_game_detail", fiber.Map{"id": id})
}

func (h *GameHandler) UpdateGameView(c *fiber.Ctx) error {
	id := c.Params("id")

	game, err := h.gameRepo.GetByID(id)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Render("update_game", game)
}

func (h *GameHandler) UpdateGame(c *fiber.Ctx) error {
	id := c.Params("id")
	var gameInput models.Game
	if err := c.BodyParser(&gameInput); err != nil {
		return c.Status(fiber.StatusInternalServerError).Render("partials/toast", fiber.Map{"error": err.Error()})
	}

	pictureFile, err := c.FormFile("picture")
	if err != nil {
		pictureFile = nil
	}

	if pictureFile != nil {
		c.SaveFile(pictureFile, fmt.Sprintf("./public/%s", pictureFile.Filename))
		picture := fmt.Sprintf("/public/%s", pictureFile.Filename)

		gameInput.Picture = picture
	}

	if err := h.gameRepo.Update(id, &gameInput); err != nil {
		return c.Status(fiber.StatusInternalServerError).Render("partials/toast", fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).Render("partials/toast", fiber.Map{"message": "Game updated successfully"})
}
