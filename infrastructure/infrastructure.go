package infrastructure

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	com "digilounge/infrastructure/functions"

	"github.com/joho/godotenv"
)

func NewDatabaseConnect(dir string) *sql.DB {
	currDir := fmt.Sprint(dir, "/.env")
	err := godotenv.Load(currDir)
	if err != nil {
		log.Fatal("(INFRASTRUCTURE:1001): ", err)
	}

	sqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("dbUser"), os.Getenv("dbPassword"),
		os.Getenv("dbEnv"), os.Getenv("dbPort"),
		os.Getenv("dbName"))

	db, err := sql.Open("mysql", sqlInfo)
	if err != nil {
		log.Fatal("(INFRASTRUCTURE:1002): ", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("(INFRASTRUCTURE:1003): ", err)
	}

	fmt.Println("Database is successfully connected")

	com.PrintLog("Database is successfully connected")

	return db
}
