package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type env struct {
	PORT string
}

var Env *env

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Panic(err.Error())
	}

	Env = &env{
		PORT: os.Getenv("PORT"),
	}
}
