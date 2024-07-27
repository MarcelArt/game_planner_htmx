package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type env struct {
	PORT       string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBHost     string
	JwtSecret  string
	IsProd     bool
}

var Env *env

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Panic(err.Error())
	}

	isProd, err := strconv.ParseBool(os.Getenv("IS_PROD"))
	if err != nil {
		log.Panic(err.Error())
	}

	Env = &env{
		PORT:       os.Getenv("PORT"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		DBHost:     os.Getenv("DB_HOST"),
		JwtSecret:  os.Getenv("JWT_SECRET"),
		IsProd:     isProd,
	}
}
