package utils

import (
	"errors"
	"strconv"
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
		"exp":      time.Now().Add(time.Second * 5).Unix(),
	}

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
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
		"exp":        expireAt.Unix(),
	}

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	t, err := token.SignedString([]byte(config.Env.JwtSecret))

	return t, err
}

func ParseToken(t string) (jwt.MapClaims, bool, error) {
	token, err := jwt.Parse(t, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.Env.JwtSecret), nil
	})
	if err != nil && errors.Is(err, jwt.ErrTokenExpired) {
		return nil, true, nil
	}
	if err != nil {
		return nil, false, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, false, nil
	}

	return nil, false, jwt.ErrSignatureInvalid
}

func ClaimsNumberToString(i interface{}) string {
	number := i.(float64)

	return strconv.Itoa(int(number))
}
