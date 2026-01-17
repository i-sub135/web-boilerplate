package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/i-sub135/web-boilerplate/internal/web/exampletocopy"
)

type HTTPServer struct {
	app *fiber.App
}

func NewHTTPServer(app *fiber.App) *HTTPServer {
	return &HTTPServer{
		app: app,
	}
}

func (s *HTTPServer) MountRoutes() *fiber.App {
	// Mount your routes here

	exRepo := exampletocopy.NewMemoryRepo()
	exUsecase := exampletocopy.NewUsecase(exRepo)
	exHandler := exampletocopy.NewHandler(exUsecase)

	// mounting example to copy feature
	userRouter := s.app.Group("/exampletocopy")
	exampletocopy.Routes(userRouter, exHandler)

	return s.app
}
