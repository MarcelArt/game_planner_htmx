package utils

import (
	"time"

	"github.com/MarcelArt/game_planner_htmx/config"
	"github.com/MarcelArt/game_planner_htmx/enums"
	"github.com/MarcelArt/game_planner_htmx/models"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateTokenPair(user *models.User, isRemember bool) (string, string, error) {
	accessToken, err := generateAccessToken(user)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := generateRefreshToken(user, isRemember)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func generateAccessToken(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"username": user.Username,
		"userId":   user.ID,
		"exp":      time.Now().Add(time.Minute * 5),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(config.Env.JwtSecret))

	return t, err
}

func generateRefreshToken(user *models.User, isRemember bool) (string, error) {
	expireAt := time.Now().Add(enums.Day)
	if isRemember {
		expireAt = time.Now().Add(enums.Month)
	}

	claims := jwt.MapClaims{
		"userId":     user.ID,
		"isRemember": isRemember,
		"exp":        expireAt,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(config.Env.JwtSecret))

	return t, err
}
