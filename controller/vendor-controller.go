package controller

import (
	//"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/M-iklan/database"
	"github.com/M-iklan/models"
	"golang.org/x/crypto/bcrypt"
)



type VendorController struct {
	// declare variables
	Db *gorm.DB
}
func InitVendorController() *VendorController {
	db,_ := database.NewDatabasePostgres()
	// gorm
	db.AutoMigrate(&models.Vendor{})

	return &VendorController{Db: db}
}


// routing
// GET 
func (controller *VendorController) SeeVendor(c *fiber.Ctx) error {
	// load all vendor
	var vendors []models.Vendor
	err := models.ReadVendor(controller.Db, &vendors)
	if err!=nil {
		return c.SendStatus(500) // http 500 internal server error
	}
	return c.Render("vendor", fiber.Map{
		"Title": "Daftar Vendor",
		"Vendor": vendors,
	})
}
// GET /products/create
func (controller *VendorController) FormRegisVendor(c *fiber.Ctx) error {
	return c.Render("register", fiber.Map{
		"Title": "Registrasi Vendor",
	})
}


// POST /products/create

func (controller *VendorController) RegisVendor(c *fiber.Ctx) error {
	var regis models.Vendor

		if err := c.BodyParser(&regis); err != nil {
			return c.Redirect("/register")
		}

		bytes, _ := bcrypt.GenerateFromPassword([]byte(regis.Password), 8)
		sHash := string(bytes)
		
		regis.Password = sHash

		err := models.RegisVendor(controller.Db, &regis)

		if err != nil {
			return c.Redirect("/register")
		}
		
		return c.Redirect("/login")
}
