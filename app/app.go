package app

import (
	"log"

	"github.com/boydmeyer/fiber/app/database"
	"github.com/boydmeyer/fiber/app/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Run(port string) {
	app := fiber.New()
	app.Use(cors.New())

	database.Connect()

	router.Setup(app)

	log.Fatal(app.Listen(port))

}
