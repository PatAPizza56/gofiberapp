package main

import (
	"github.com/gofiber/fiber"

	routes "./routes"
)

var (
	app *fiber.App
)

func main() {
	start()
	handle()
	listen()
}

func start() {
	// Create a new fiber app
	app = fiber.New()
}

func handle() {
	// Set up a home route
	app.Get("/", routes.Home)

	// Set up different routes
	app.Get("/routes", routes.Routes)

	// Use * for a wildcard, and :paramName for a parameter
	app.Get("/api/*", routes.API)

	// Set a static files directory
	app.Static("/static", "./static")

}

func listen() {
	// Set a port, 80 is for http communication
	app.Listen(":80")
}
