package router

import (
	"database/sql"
	_ "digilounge/delivery/http/controllers"

	"github.com/gofiber/fiber/v2"
)

func NewRouter(app *fiber.App, db *sql.DB, c *fiber.Ctx) {

}
