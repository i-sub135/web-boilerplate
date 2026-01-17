package exampletocopy

import (
	"github.com/flosch/pongo2/v7"
	"github.com/gofiber/fiber/v2"
	"github.com/i-sub135/web-boilerplate/common/utils/webutil"
)

type Handler struct {
	uc *Usecase
}

func NewHandler(uc *Usecase) *Handler {
	return &Handler{uc: uc}
}

func (h *Handler) Page(c *fiber.Ctx) error {
	users, _ := h.uc.GetAll(c.Context())
	return webutil.RenderPage(
		c,
		"exampletocopy/template/page.html",
		pongo2.Context{
			"users": users,
		},
	)
}
