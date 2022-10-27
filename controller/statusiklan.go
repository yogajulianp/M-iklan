package controller

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/M-iklan/models"
)

type IklanStatusController struct {
	Db *gorm.DB
}

func InitIklanStatusController(db *gorm.DB) *IklanStatusController {
	return &IklanStatusController{Db: db}
}

// route
func (controller *IklanStatusController) AdminRoute(app *fiber.App) {
	stat := app.Group("/statusiklan")
	stat.Get("/", controller.GetIklanStatues)
	stat.Post("/detail/:id", controller.GetDetailIklanStatus) // input form name, qunatity, price, picture
	stat.Put("/editstatus/:id", controller.EditStatusIklan)
}

// get iklanstatues
func (controller *IklanStatusController) GetIklanStatues(c *fiber.Ctx) error {
	var iklanStatuses []models.IklanStatus
	err := models.ReadIklanStatus(controller.Db, &iklanStatuses)
	if err != nil {
		return c.SendStatus(500) // http 500 internal server error
	}

	return c.JSON(fiber.Map{
		"products": iklanStatuses,
	})
}

// generate iklanstatus from iklan --> dimasukkan dalam controller iklan bagian bawah pembuatan iklan
// mengenerate iklanstatus setelah iklan ditambahkan
//
//		var iklanStatus models.IklanStatus
//		if err:=models.CreateIklanStatus(controller.Db, &iklanStatus, iklan.ID); err!=nil {
//			c.SendStatus(500)
//		}
//		return c.JSON(fiber.Map{
//			"iklanStatus": iklanStatus,
//		})
//

// get detail
func (controller *IklanStatusController) GetDetailIklanStatus(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")

	var iklanStatus models.IklanStatus
	if err := models.ReadStatusById(controller.Db, &iklanStatus, uint(id)); err != nil {
		return c.SendStatus(500) // http 500 internal server error
	}
	var iklan models.Iklan
	if err := models.ReadIklanById(controller.Db, &iklan, iklanStatus.IklanID); err != nil {
		return c.SendStatus(500)
	}
	return c.JSON(fiber.Map{
		"iklan status": iklanStatus,
		"detail iklan": iklan,
	})
}

// edit iklan status
func (controller *IklanStatusController) EditStatusIklan(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")

	var statusIklan models.IklanStatus
	if err := models.ReadStatusById(controller.Db, &statusIklan, uint(id)); err != nil {
		return c.SendStatus(500) // http 500 internal server error
	}
	var updateStatus models.IklanStatus
	if err := c.BodyParser(&updateStatus); err != nil {
		return c.SendStatus(400)
	}
	statusIklan.IklanID = updateStatus.IklanID

	// save product
	models.UpdateIklanStatus(controller.Db, &statusIklan)

	return c.JSON(statusIklan)
}
