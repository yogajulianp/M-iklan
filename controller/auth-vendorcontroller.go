package controller

import (
	"fmt"

	"github.com/M-iklan/database"
	"github.com/M-iklan/models"
	"github.com/M-iklan/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type LoginRequest struct {
	Username string `form:"username" json:"username" validate:"required"`
	Password string `form:"password" json:"password" validate:"required"`
}

type AuthController struct {
	// declare variables
	Db *gorm.DB
}

func InitAuthController() *AuthController {
	return &AuthController{}
}

// GET /login
func (controller *AuthController) LoginForm(c *fiber.Ctx) error {
	// load all vendor

	return c.Render("vendorlogin", fiber.Map{
		"Title": "Login Vendor",
	})
}

func (controller *AuthController) AuthLogin(c *fiber.Ctx) error {
	loginRequest := new(LoginRequest)
	if err := c.BodyParser(&loginRequest); err != nil {
		return c.Redirect("/login")
	}
	fmt.Println(loginRequest)

	//validasi request
	validate := validator.New()
	errValidate := validate.Struct(loginRequest)
	if errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	//check available user
	var vendor models.Vendor
	db, _ := database.NewDatabasePostgres()
	err := db.Where("username = ?", loginRequest.Username).First(&vendor).Error
	if err != nil {
		return c.Redirect("/login")
	}

	//check validasi password
	validPassword := utils.CheckPasswordHash(loginRequest.Password, vendor.Password)
	if !validPassword {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "wrong credentials",
		})
	}

	return c.Redirect("/admin/iklan")
}
