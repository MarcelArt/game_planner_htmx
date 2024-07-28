package handlers

import (
	"fmt"
	"log"
	"time"

	"github.com/MarcelArt/game_planner_htmx/database"
	"github.com/MarcelArt/game_planner_htmx/enums"
	"github.com/MarcelArt/game_planner_htmx/models"
	"github.com/MarcelArt/game_planner_htmx/repositories"
	"github.com/MarcelArt/game_planner_htmx/utils"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	userRepo            repositories.IUserRepo
	connectedDeviceRepo repositories.IConnectedDeviceRepo
	profileRepo         repositories.IProfileRepo
}

func NewAuthHandler(userRepo repositories.IUserRepo, connectedDeviceRepo repositories.IConnectedDeviceRepo, profileRepo repositories.IProfileRepo) *AuthHandler {
	return &AuthHandler{
		userRepo:            userRepo,
		connectedDeviceRepo: connectedDeviceRepo,
		profileRepo:         profileRepo,
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

	tx := database.DB.Begin()
	createdUser, err := h.userRepo.CreateTrx(tx, &user)
	if err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).Render("register", fiber.Map{
			"error": err.Error(),
		}, "layouts/main")
	}

	_, err = h.profileRepo.CreateTrx(tx, &models.Profile{
		Name:   createdUser.Username,
		Avatar: fmt.Sprintf("https://ui-avatars.com/api/?name=%s", createdUser.Username),
		UserID: createdUser.ID,
	})
	if err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).Render("register", fiber.Map{
			"error": err.Error(),
		}, "layouts/main")
	}

	tx.Commit()
	return c.Status(fiber.StatusPermanentRedirect).Redirect("/")
}

func (h *AuthHandler) LoginView(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).Render("login", nil, "layouts/main")
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	isRememberStr := c.FormValue("isRemember", "false")
	isRemember := utils.ParseCheckboxToBool(isRememberStr)

	user, err := h.userRepo.GetByUsernameOrEmail(username)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).Render("login", fiber.Map{
			"error": err.Error(),
		}, "layouts/main")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).Render("login", fiber.Map{
			"error": err.Error(),
		}, "layouts/main")
	}

	accessToken, refreshToken, err := utils.GenerateTokenPair(user, isRemember)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).Render("login", fiber.Map{
			"error": err.Error(),
		}, "layouts/main")
	}

	aCookie := utils.CreateCookie("at", accessToken, time.Now().Add(5*time.Minute))

	expireAt := time.Now().Add(enums.Day)
	if isRemember {
		expireAt = time.Now().Add(enums.Month)
	}
	rCookie := utils.CreateCookie("rt", refreshToken, expireAt)

	device := &models.ConnectedDevice{
		RefreshToken: refreshToken,
		UserAgent:    c.Get("User-Agent", ""),
		Ip:           c.IP(),
		UserID:       user.ID,
	}
	device, err = h.connectedDeviceRepo.Create(device)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).Render("login", fiber.Map{
			"error": err.Error(),
		}, "layouts/main")
	}

	c.Cookie(aCookie)
	c.Cookie(rCookie)

	return c.Status(fiber.StatusPermanentRedirect).Redirect("/")
}
