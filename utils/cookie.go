package utils

import (
	"time"

	"github.com/MarcelArt/game_planner_htmx/config"
	"github.com/gofiber/fiber/v2"
)

func CreateCookie(key string, value string, expireAt time.Time) *fiber.Cookie {
	cookie := &fiber.Cookie{
		Name:     key,
		Value:    value,
		HTTPOnly: true,
		Secure:   config.Env.IsProd,
		SameSite: "lax",
		Path:     "/",
		Expires:  expireAt,
	}

	return cookie
}
