package handlers

import (
	"github.com/MarcelArt/game_planner_htmx/middleware"
	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	user, _ := middleware.GetCurrentUser(c)
	return c.Render("index", fiber.Map{
		"Title": user.Username,
	}, "layouts/main")
}
