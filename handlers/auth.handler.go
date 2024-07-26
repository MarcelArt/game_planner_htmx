package handlers

import (
	"log"

	"github.com/MarcelArt/game_planner_htmx/models"
	"github.com/MarcelArt/game_planner_htmx/repositories"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
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
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).Render("register", fiber.Map{
			"error": err.Error(),
		}, "layouts/main")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).Render("register", fiber.Map{
			"error": err.Error(),
		}, "layouts/main")
	}
	user.Password = string(hashedPassword)

	_, err = h.userRepo.Create(&user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).Render("register", fiber.Map{
			"error": err.Error(),
		}, "layouts/main")
	}

	return c.Status(fiber.StatusPermanentRedirect).Redirect("/")
}
