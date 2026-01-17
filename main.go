package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/i-sub135/web-boilerplate/common/utils/webutil"
	"github.com/i-sub135/web-boilerplate/internal/app"
)

func main() {
	fibApp := fiber.New()

	// Add logger middleware
	fibApp.Use(logger.New(logger.Config{
		Format: "${time} - ${ip} - ${status} - ${method} ${path}\n",
	}))

	// initialize template engine
	webutil.InitTemplates()

	fibApp.Static("/public", "./public")
	fibApp.Static("/features", "./internal/web")

	httpServer := app.NewHTTPServer(fibApp)
	httpServer.MountRoutes()

	fibApp.Listen(":3003")
}
