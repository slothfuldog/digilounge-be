package router

import (
	"database/sql"
	"digilounge/delivery/http/controllers"

	"github.com/gofiber/fiber/v2"
)

func NewRouter(app *fiber.App, db *sql.DB, c *fiber.Ctx) {
	app.Post("/login", func(c *fiber.Ctx) error {
		return controllers.Login(db, c)
	})
}
