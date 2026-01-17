package exampletocopy

import "github.com/gofiber/fiber/v2"

func Routes(r fiber.Router, h *Handler) {
	r.Get("/", h.Page)
}
