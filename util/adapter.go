package util

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func Render(ctx *fiber.Ctx, comp templ.Component) error {
	handler := adaptor.HTTPHandler(templ.Handler(comp))
	return handler(ctx)
}
