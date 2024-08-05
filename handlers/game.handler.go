package handlers

import "github.com/gofiber/fiber/v2"

type GameHandler struct {
}

func NewGameHandler() *GameHandler {
	return &GameHandler{}
}

func (h *GameHandler) GamesView(c *fiber.Ctx) error {
	return c.Render("games", nil)
}
