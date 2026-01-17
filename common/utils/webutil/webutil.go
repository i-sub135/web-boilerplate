package webutil

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/flosch/pongo2/v7"
	"github.com/gofiber/fiber/v2"
)

// Daftar folder template, urut prioritas
var templateFolders = []string{
	"templates",    // global
	"internal/web", // per-feature
}

// Loader custom
type MultiFolderLoader struct{}

func (l MultiFolderLoader) Abs(base, name string) string {
	// Kita pakai path relatif saja
	return name
}

func (l MultiFolderLoader) Get(filename string) (io.Reader, error) {
	// normalize path
	filename = strings.TrimLeft(filename, "/")
	for _, folder := range templateFolders {
		fullPath := filepath.Join(folder, filename)
		if _, err := os.Stat(fullPath); err == nil {
			f, err := os.Open(fullPath)
			if err != nil {
				return nil, err
			}
			return f, nil
		}
	}
	return nil, fmt.Errorf("template not found: %s", filename)
}

var tplSet *pongo2.TemplateSet

// Inisialisasi template
func InitTemplates() {
	tplSet = pongo2.NewSet("app_templates", MultiFolderLoader{})
}

// RenderPage
func RenderPage(c *fiber.Ctx, tpl string, ctx pongo2.Context) error {
	if tplSet == nil {
		return fmt.Errorf("template engine not initialized")
	}

	t, err := tplSet.FromCache(tpl)
	if err != nil {
		return fmt.Errorf("unable to load template %s: %w", tpl, err)
	}

	// set content type ke HTML
	c.Type("html", "utf-8")

	return t.ExecuteWriter(ctx, c.Response().BodyWriter())
}
