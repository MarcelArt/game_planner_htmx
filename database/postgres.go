package database

import (
	"fmt"
	"strconv"

	"github.com/MarcelArt/game_planner_htmx/config"
	"github.com/MarcelArt/game_planner_htmx/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	p := config.Env.DBPort
	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		panic("failed to parse database port")
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Env.DBHost, port, config.Env.DBUser, config.Env.DBPassword, config.Env.DBName)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")
	Migrate()
}

func Migrate() {
	DB.AutoMigrate(
		&models.User{},
		&models.ConnectedDevice{},
		&models.Profile{},
		&models.Game{},
		&models.Item{},
		&models.Recipe{},
		&models.RecipeDetail{},
	)
	fmt.Println("Database Migrated")
}
