package main

import (
	"fmt"
	"log"
	"os"

	"github.com/M-iklan/database"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	db, err := database.NewDatabasePostgres()

	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate()

	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error { fmt.Println("test"); return nil })

	app.Listen(os.Getenv("SERVER_PORT"))

}
