package main

import (
	"log"
	"os"

	"github.com/M-iklan/controller"
	"github.com/M-iklan/database"
	"github.com/M-iklan/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
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

	err = db.AutoMigrate(models.Iklan{})

	if err != nil {
		log.Fatal(err)
	}

	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/", "./public", fiber.Static{
		Index: "",
	})

	adsDisplay := controller.NewAdsDisplay(db)

	adsDisplay.MountRouter(app)

	app.Listen(os.Getenv("SERVER_PORT"))

}
