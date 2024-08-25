package controllers

import (
	"database/sql"
	com "digilounge/infrastructure/functions"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Login(db *sql.DB, c *fiber.Ctx) error {

	com.PrintLog("======= Login Function START =========")

	var username string

	query := "SELECT * FROM user_inf WHERE username %s"

	err := db.QueryRow(query).Scan(username)

	if err != nil {
		com.PrintLog(fmt.Sprintf("(4224) %s", err))
	}

	com.PrintLog("======= Login Function END   =========")

	return c.Status(200).JSON(fiber.Map{
		"status": 200,
	})
}
