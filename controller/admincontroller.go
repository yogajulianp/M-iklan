package controller

import (
	"github.com/M-iklan/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AdminController struct {
	// declare variables
	Db *gorm.DB
}

func InitAdminController(db *gorm.DB) *AdminController {

	return &AdminController{Db: db}
}

// route
func (admincontroller *AdminController) AdminDashboardRoute(app *fiber.App) {
	admin := app.Group("/admindashboard")
	admin.Get("/", admincontroller.GetAllVendor)
}

// Get AdminDashboard

func (admincontroller *AdminController) GetAllVendor(c *fiber.Ctx) error {
	var admin []models.Admin
	erradmin := models.ReadAdmin(admincontroller.Db, &admin)
	if erradmin != nil {
		return c.SendStatus(500)
	}

	var vendor []models.Vendor
	errvendor := models.ReadVendor(admincontroller.Db, &vendor)
	if errvendor != nil {
		return c.SendStatus(500)
	}

	return c.Render("adminvideo/dashboardadmin", fiber.Map{
		"Title":  "Daftar Vendor",
		"Admin":  admin,
		"Vendor": vendor,
	})
}
