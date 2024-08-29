package controllers

import (
	"database/sql"
	en "digilounge/infrastructure/encryptor"
	com "digilounge/infrastructure/functions"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type LoginStruct struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(db *sql.DB, c *fiber.Ctx) error {

	com.PrintLog("======= Login Function START =========")

	var login LoginStruct
	var password string

	if err := c.BodyParser(login); err != nil {
		com.PrintLog(fmt.Sprintf("(0001)%s", err))
		return c.Status(501).JSON(fiber.Map{
			"status":  501,
			"message": "Error on Server Side",
		})
	}

	query := fmt.Sprintf("SELECT password FROM user_inf WHERE username %s", login.Username)

	err := db.QueryRow(query).Scan(password)

	if err != nil {
		com.PrintLog(fmt.Sprintf("(0002) %s", err))
		return c.Status(404).JSON(fiber.Map{
			"status":  404,
			"message": "User not found",
		})
	}

	res, ers := en.VerifyPassword(password, login.Password)

	if ers != nil {
		com.PrintLog(fmt.Sprintf("(0003) %s", err))
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": "Password or Username is wrong",
		})
	}

	if res {
		fmt.Println("good")
	}

	com.PrintLog("======= Login Function END   =========")

	return c.Status(200).JSON(fiber.Map{
		"status": 200,
	})
}
