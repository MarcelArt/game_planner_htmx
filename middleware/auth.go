package middleware

import (
	"log"
	"strconv"
	"time"

	"github.com/MarcelArt/game_planner_htmx/database"
	"github.com/MarcelArt/game_planner_htmx/enums"
	"github.com/MarcelArt/game_planner_htmx/models"
	"github.com/MarcelArt/game_planner_htmx/repositories"
	"github.com/MarcelArt/game_planner_htmx/utils"
	"github.com/gofiber/fiber/v2"
)

func Auth(c *fiber.Ctx) error {
	path := c.Path()
	if path == "/login" || path == "/register" {
		c.Next()
	}

	at := c.Cookies("at", "")
	rt := c.Cookies("rt", "")

	if at == "" || rt == "" {
		return c.Redirect("/login")
	}

	aClaims, isAccessExpired, err := utils.ParseToken(at)
	if isAccessExpired {
		return refreshTokenPair(c, rt)
	}
	if err != nil {
		log.Println(err.Error())
		return c.Redirect("/login")
	}

	c.Locals("userId", aClaims["userId"])

	return c.Next()
}

func refreshTokenPair(c *fiber.Ctx, rt string) error {
	rClaims, isRefreshExpired, err := utils.ParseToken(rt)
	if isRefreshExpired || err != nil {
		return c.Redirect("/login")
	}

	userRepo := repositories.NewUserRepo(database.DB)

	isRemember := rClaims["isRemember"].(bool)
	userID := utils.ClaimsNumberToString(rClaims["userId"])

	user, err := userRepo.GetByID(userID)
	if err != nil {
		return c.Redirect("/login")
	}

	accessToken, refreshToken, err := utils.GenerateTokenPair(user, isRemember)
	if err != nil {
		return c.Redirect("/login")
	}

	aCookie := utils.CreateCookie("at", accessToken, time.Now().Add(5*time.Minute))

	expireAt := time.Now().Add(enums.Day)
	if isRemember {
		expireAt = time.Now().Add(enums.Month)
	}
	rCookie := utils.CreateCookie("rt", refreshToken, expireAt)

	connectedDeviceRepo := repositories.NewConnectedDeviceRepo(database.DB)
	device, err := connectedDeviceRepo.GetByToken(rt)
	if err != nil {
		return c.Redirect("/login")
	}
	device.RefreshToken = refreshToken
	device.Ip = c.IP()
	device.UserAgent = c.Get("User-Agent")
	deviceID := strconv.Itoa(int(device.ID))
	if err := connectedDeviceRepo.Update(deviceID, device); err != nil {
		return c.Redirect("/login")
	}

	c.Cookie(aCookie)
	c.Cookie(rCookie)
	c.Locals("userId", userID)

	return c.Next()
}

func GetCurrentUser(c *fiber.Ctx) (*models.User, error) {
	userID, ok := c.Locals("userId").(string)
	if !ok {
		userID = utils.ClaimsNumberToString(c.Locals("userId"))
	}

	userRepo := repositories.NewUserRepo(database.DB)
	user, err := userRepo.GetByID(userID)

	return user, err
}
