package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/M-iklan/models"
	"github.com/M-iklan/database"
	"github.com/gofiber/fiber/v2/middleware/session"
	//"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// type LoginRequest struct {
// 	Username	string    `form:"username" json:"username" validate:"required"`
// 	Password    string	`form:"password" json:"password" validate:"required"`
// }

type AuthController struct {
	// declare variables
	Db *gorm.DB
	store *session.Store
}

// GET /login
func (controller *AuthController) LoginForm (c *fiber.Ctx) error {
	

	return c.Render("vendorlogin", fiber.Map{
		"Title": "Login Vendor",
	})
}


func (controller *AuthController) AuthLogin(c *fiber.Ctx) error {
	sess, err := controller.store.Get(c)

	if err!=nil {
		panic(err)
	}

	var myform models.Vendor
	var data models.Vendor

	if err := c.BodyParser(&myform); err != nil {
		return c.JSON(fiber.Map{"error": err})
	}

	username := myform.Username
	plainPassword := myform.Password

	err2 := models.ReadVendorByUser(controller.Db, &data, username)

	if err2 != nil {
		return c.Redirect("/login")
	}
	
	hashPassword := data.Password

	check := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(plainPassword))

	status := check == nil

	if status {
		sess.Set("username", username)
		sess.Set("id", data.Id)
		sess.Save()
		return c.Redirect("/admin/iklan")
	} else {
		return c.Redirect("/login")
	}
}

// /logout
func (controller *AuthController) Logout(c *fiber.Ctx) error {
	sess, err := controller.store.Get(c)

	if err != nil {
		panic(err)
	}

	sess.Destroy()

	return c.Render("login", fiber.Map{
		"Title": "Login",
	})
}

func InitAuthController(s *session.Store) *AuthController {
	db,_ := database.NewDatabasePostgres()

	db.AutoMigrate(&models.Vendor{})
	
	return &AuthController{Db: db, store: s}
}