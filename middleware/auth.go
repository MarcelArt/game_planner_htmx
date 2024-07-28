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

type AuthMiddleware struct {
	userRepo            repositories.IUserRepo
	connectedDeviceRepo repositories.IConnectedDeviceRepo
}

func NewAuthMiddleware(userRepo repositories.IUserRepo, connectedDeviceRepo repositories.IConnectedDeviceRepo) *AuthMiddleware {
	return &AuthMiddleware{
		userRepo:            userRepo,
		connectedDeviceRepo: connectedDeviceRepo,
	}
}

func (m *AuthMiddleware) Auth(c *fiber.Ctx) error {
	path := c.Path()
	if path == "/login" || path == "/register" {
		c.Next()
	}

	at := c.Cookies("at", "")
	rt := c.Cookies("rt", "")

	if at == "" || rt == "" {
		log.Println("Not logged in")
		return c.Redirect("/login")
	}

	aClaims, isAccessExpired, err := utils.ParseToken(at)
	if isAccessExpired {
		log.Println("Token refreshed")
		return m.refreshTokenPair(c, rt)
	}
	if err != nil {
		log.Println("1. Token invalid redirect to login")
		log.Println(err.Error())
		return c.Redirect("/login")
	}

	c.Locals("userId", aClaims["userId"])

	log.Println("Authorized")
	return c.Next()
}

func (m *AuthMiddleware) refreshTokenPair(c *fiber.Ctx, rt string) error {
	rClaims, isRefreshExpired, err := utils.ParseToken(rt)
	if isRefreshExpired || err != nil {
		return c.Redirect("/login")
	}

	isRemember := rClaims["isRemember"].(bool)
	userID := utils.ClaimsNumberToString(rClaims["userId"])

	user, err := m.userRepo.GetByID(userID)
	if err != nil {
		log.Println("2. Token invalid redirect to login")
		return c.Redirect("/login")
	}

	accessToken, refreshToken, err := utils.GenerateTokenPair(user, isRemember)
	if err != nil {
		log.Println("3. Token invalid redirect to login")
		return c.Redirect("/login")
	}

	aCookie := utils.CreateCookie("at", accessToken, time.Now().Add(5*time.Minute))

	expireAt := time.Now().Add(enums.Day)
	if isRemember {
		expireAt = time.Now().Add(enums.Month)
	}
	rCookie := utils.CreateCookie("rt", refreshToken, expireAt)

	device, err := m.connectedDeviceRepo.GetByToken(rt)
	if err != nil {
		log.Println("4. Token invalid redirect to login")
		return c.Redirect("/login")
	}
	device.RefreshToken = refreshToken
	device.Ip = c.IP()
	device.UserAgent = c.Get("User-Agent")
	deviceID := strconv.Itoa(int(device.ID))
	if err := m.connectedDeviceRepo.Update(deviceID, device); err != nil {
		log.Println("5. Token invalid redirect to login")
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

func GetCurrentProfile(c *fiber.Ctx) (*models.Profile, error) {
	user, err := GetCurrentUser(c)
	if err != nil {
		return nil, err
	}

	profileRepo := repositories.NewProfileRepo(database.DB)
	profile, err := profileRepo.GetByUserID(user.ID)

	return profile, err
}
