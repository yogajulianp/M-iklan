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
	app.Static("/public", "./public")

	adsDisplay := controllers.NewAdsDisplay(db)
	iklancontroller := controllers.NewIklan(db)

	adsDisplay.MountRouter(app)
	iklancontroller.RouteIklan(app)

	app.Get("/dashboard", func(c *fiber.Ctx) error {

		return c.Render("admin/dashboard", fiber.Map{
			"Title": "Dashboard",
		})
	})

	app.Get("/login", func(c *fiber.Ctx) error {

		return c.Render("admin/login", fiber.Map{
			"Title": "Login",
		})
	})

	app.Listen(os.Getenv("SERVER_PORT"))

}
