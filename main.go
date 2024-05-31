package main

import (
	"digilounge/delivery/http"
	"log"
	"os"
)

func main() {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal("(MAIN:1000): ", err)
	}
	currentDir := path
	app := http.NewHttpDelivery(currentDir)

	app.Listen(":3000")

}
