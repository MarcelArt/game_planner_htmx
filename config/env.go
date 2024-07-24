package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type env struct {
	PORT       string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBHost     string
}

var Env *env

func init() {
	err := godotenv.Load(".env")
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
	}
}
