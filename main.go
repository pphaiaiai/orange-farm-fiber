package main

import (
	"os"

	"github.com/gofiber/fiber/v2"

	"github.com/pphaiaiai/orange-farm-fiber/app/pkg/configs"
	"github.com/pphaiaiai/orange-farm-fiber/app/pkg/utils"
)

func main() {
	// Define Fiber config.
	config := configs.FiberConfig()

	// Define a new Fiber app with config.
	app := fiber.New(config)

	// Routes.

	// Middlewares.

	// Start server (with or without graceful shutdown).
	if os.Getenv("STAGE_STATUS") == "dev" {
		utils.StartServer(app)
	} else {
		utils.StartServerWithGracefulShutdown(app)
	}
}
