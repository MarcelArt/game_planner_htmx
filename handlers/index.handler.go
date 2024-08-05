package handlers

import (
	"github.com/MarcelArt/game_planner_htmx/middleware"
	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	profile, _ := middleware.GetCurrentProfile(c)
	return c.Render("index", fiber.Map{
		"profile": profile,
	}, "layouts/main")
}

// func Index(c *fiber.Ctx) error {
// 	profile, _ := middleware.GetCurrentProfile(c)
// 	return c.Render("index", fiber.Map{
// 		"profile": profile,
// 	}, "layouts/main")
// }
