package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/M-iklan/controller"
	
)



func RouteInit(r *fiber.App) {
	authController := controller.InitAuthController()
	vendorController := controller.InitVendorController()
	r.Get("/login", authController.LoginForm)
	r.Post("/login", authController.AuthLogin)
	r.Get("/register", vendorController.FormRegisVendor)
	r.Post("/register", vendorController.RegisVendor)


	
}