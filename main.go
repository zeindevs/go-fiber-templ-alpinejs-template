package main

import (
	"embed"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/zeindevs/go-fiber-templ-alpinejs-template/config"
	"github.com/zeindevs/go-fiber-templ-alpinejs-template/handlers"
	"github.com/zeindevs/go-fiber-templ-alpinejs-template/util"
)

//go:embed static/*
var embedStaticDir embed.FS

func main() {
	cfg := config.NewConfig()

	app := fiber.New()
	app.Use(logger.New())
	app.Use("/static", filesystem.New(filesystem.Config{
		Root:       http.FS(embedStaticDir),
		PathPrefix: "static",
	}))

	app.Get("/", handlers.HandleIndex)

	util.ErrorPanic(app.Listen(cfg.Get("PORT")))
}
