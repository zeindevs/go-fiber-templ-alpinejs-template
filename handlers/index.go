package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zeindevs/go-fiber-templ-alpinejs-template/util"
	"github.com/zeindevs/go-fiber-templ-alpinejs-template/views"
)

func HandleIndex(ctx *fiber.Ctx) error {
	index := views.Index("Go+Fiber+Templ+Alpinejs")
	return util.Render(ctx, index)
}
