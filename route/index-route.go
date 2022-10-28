package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/M-iklan/controller"
	
)



func RouteInit(r *fiber.App) {
	authController := controller.InitAuthController()
	

	r.Post("/login", authController.AuthLogin)

	
}