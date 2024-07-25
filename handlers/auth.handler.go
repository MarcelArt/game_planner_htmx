package handlers

import (
	"log"

	"github.com/MarcelArt/game_planner_htmx/models"
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
	return c.Status(fiber.StatusOK).Render("register", nil, "layouts/main")
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var user *models.User
	if err := c.BodyParser(&user); err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).Render("register", fiber.Map{
			"error": err.Error(),
		}, "layouts/main")
	}

	return c.Status(fiber.StatusPermanentRedirect).Redirect("/")
}
