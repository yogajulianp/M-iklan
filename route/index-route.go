package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/M-iklan/controller"
	"github.com/gofiber/fiber/v2/middleware/session"
)



func RouteInit(r *fiber.App) {
	store := session.New()
	authController := controller.InitAuthController(store)
	vendorController := controller.InitVendorController()

	r.Get("/login", authController.LoginForm)
	r.Post("/login", authController.AuthLogin)
	r.Get("/logout", authController.Logout)
	r.Get("/", vendorController.FormRegisVendor)
	r.Get("/register", vendorController.FormRegisVendor)
	r.Post("/register", vendorController.RegisVendor)

}