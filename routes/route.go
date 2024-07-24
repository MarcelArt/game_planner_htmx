package routes

import (
	"github.com/MarcelArt/game_planner_htmx/handlers"
	"github.com/MarcelArt/game_planner_htmx/lib"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	gofiberfirebaseauth "github.com/sacsand/gofiber-firebaseauth"
)

func SetupRoutes(app *fiber.App) {
	app.Use(cors.New())
	app.Use(gofiberfirebaseauth.New(gofiberfirebaseauth.Config{
		FirebaseApp: lib.FireApp,
		// IgnoreUrls: []string{},
	}))

	app.Static("/scripts", "./views/scripts")
	// app.Static("/src", "./views/src")
	app.Get("/", handlers.Index)
}
