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
func (controller *AdminController) AdminDashboardRoute(app *fiber.App) {
	stat := app.Group("/admindashboard")
	stat.Get("/", controller.GetVendor)
	stat.Post("/detailvendor/:id", controller.GetDetailVendor) // input form name, qunatity, price, picture
}

// Get AdminDashboard

func (controller *AdminController) GetAllVendor(app *fiber.App) {
	var admin []models.Admin
	err := models.ReadAdmin(controller.Db, &admin)
	if err != nil {
		return app.SendStatus(500) // http 500 internal server error
	}
	return app.Render("adminvideo/dashboardadmin", fiber.Map{
		"Title": "Daftar Vendor",
		"Admin": admin,
	})
}
