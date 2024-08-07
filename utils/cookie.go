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
		MaxAge:   1000 * 60 * 60 * 24 * 365 * 10,
		// Expires:  expireAt,
	}

	return cookie
}

func DeleteCookie(key string) *fiber.Cookie {
	return &fiber.Cookie{
		Name:    key,
		Path:    "/",
		MaxAge:  -1,
		Expires: time.Now().Add(-100 * time.Hour),
	}
}
