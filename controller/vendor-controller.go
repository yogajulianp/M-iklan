package controller

import (
	//"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/M-iklan/database"
	"github.com/M-iklan/models"
)

type VendorController struct {
	// declare variables
	Db *gorm.DB
}

func InitVendorController() *VendorController {
	db, _ := database.NewDatabasePostgres()
	// gorm
	db.AutoMigrate(&models.Vendor{})

	return &VendorController{Db: db}
}

// routing
// GET /vendor
func (controller *VendorController) SeeVendor(c *fiber.Ctx) error {
	// load all vendor and iklan
	var vendors []models.Vendor
	err := models.ReadVendor(controller.Db, &vendors)
	if err != nil {
		return c.SendStatus(500) // http 500 internal server error
	}
	return c.Render("products", fiber.Map{
		"Title":  "Daftar Iklan",
		"Vendor": vendors,
	})
}

// GET /vendor/create
func (controller *VendorController) FormRegisVendor(c *fiber.Ctx) error {
	return c.Render("register", fiber.Map{
		"Title": "Registrasi Vendor",
	})
}

// POST /vendor/create
func (controller *VendorController) RegisVendor(c *fiber.Ctx) error {

	var myform models.Vendor

	if err := c.BodyParser(&myform); err != nil {
		return c.Redirect("/login")
	}
	// save vendor
	err := models.RegisVendor(controller.Db, &myform)
	if err != nil {
		return c.Redirect("/login")
	}
	// if succeed
	return c.Redirect("/login")
}
