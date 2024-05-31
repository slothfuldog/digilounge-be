package http

import (
	"digilounge/delivery/http/router"
	"digilounge/infrastructure"

	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/gofiber/fiber/v2"
)

func NewHttpDelivery(dir string) *fiber.App {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3001, http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	var c *fiber.Ctx

	db := infrastructure.NewDatabaseConnect(dir)

	router.NewRouter(app, db, c)

	return app
}
