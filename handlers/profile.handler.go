package handlers

import (
	"fmt"
	"log"

	"github.com/MarcelArt/game_planner_htmx/middleware"
	"github.com/MarcelArt/game_planner_htmx/models"
	"github.com/MarcelArt/game_planner_htmx/repositories"
	"github.com/gofiber/fiber/v2"
)

type ProfileHandler struct {
	profileRepo repositories.IProfileRepo
}

func NewProfileHandler(profileRepo repositories.IProfileRepo) *ProfileHandler {
	return &ProfileHandler{
		profileRepo: profileRepo,
	}
}

func (h *ProfileHandler) MyProfileView(c *fiber.Ctx) error {
	profile, _ := middleware.GetCurrentProfile(c)

	return c.Status(fiber.StatusOK).Render("profile", profile)
}

func (h *ProfileHandler) Update(c *fiber.Ctx) error {
	id := c.FormValue("profileId")
	name := c.FormValue("name")
	avatarFile, err := c.FormFile("file")
	if err != nil {
		log.Println(err.Error())
		avatarFile = nil
	}

	var avatar string
	if avatarFile != nil {
		c.SaveFile(avatarFile, fmt.Sprintf("./public/uploads/%s", name+avatarFile.Filename))
		avatar = fmt.Sprintf("/public/uploads/%s", name+avatarFile.Filename)
	}

	err = h.profileRepo.Update(id, &models.Profile{
		Name:   name,
		Avatar: avatar,
	})
	if err != nil {
		return c.Render("partials/toast", fiber.Map{"error": err.Error()})
	}

	return c.Render("partials/toast", fiber.Map{"message": "Success"})
}
