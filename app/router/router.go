package router

import (
	"github.com/boydmeyer/fiber/app/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Setup(app *fiber.App) {
	api := app.Group("/", logger.New())
	api.Get("/", handler.Hello)
}
