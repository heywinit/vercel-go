package handler

import (
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	// Create a new Fiber app
	app := fiber.New()

	// Add a route
	app.Get("/*", func(c *fiber.Ctx) error {
		programName := os.Getenv("PROGRAM_NAME")
		if programName == "" {
			programName = "Vercel Go with Fiber"
		}
		return c.SendString("<h1>Hello from Go Fiber! " + programName + "</h1>")
	})

	// Convert Fiber app to http.Handler and serve the request
	handler := adaptor.FiberApp(app)
	handler.ServeHTTP(w, r)
}

