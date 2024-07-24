package handlers

import (
	"github.com/MarcelArt/game_planner_htmx/repositories"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	userRepo repositories.IUserRepo
}

func NewAuthHandler(userRepo repositories.IUserRepo) *AuthHandler {
	return &AuthHandler{
		userRepo: userRepo,
	}
}

func (h *AuthHandler) RegisterView(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).Render("register", nil)
}
