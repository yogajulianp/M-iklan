package controllers

import (
	"fmt"
	"strconv"

	"github.com/M-iklan/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type IklanAPI struct {
	Db *gorm.DB
}

func NewIklanAPI(db *gorm.DB) *IklanAPI {
	return &IklanAPI{Db: db}
}

// Routing
func (iklanapicontroller *IklanAPI) RouteIklanAPI(app *fiber.App) {
	router := app.Group("/admin/iklan/api")
	router.Get("/", iklanapicontroller.GetAllIklanAPI)
	router.Post("/create", iklanapicontroller.CreateIklanAPI)
	router.Get("/detail/:id", iklanapicontroller.DetailIklanAPI)
	router.Put("/edit/:id", iklanapicontroller.EditIklanAPI)
	router.Delete("/delete/:id", iklanapicontroller.DeleteIklanAPI)
}

// GET /admin/iklan/api
func (iklanapicontroller *IklanAPI) GetAllIklanAPI(c *fiber.Ctx) error {
	// Load all iklans
	var iklans []models.Iklan
	err := models.ReadIklans(iklanapicontroller.Db, &iklans)
	if err != nil {
		return c.SendStatus(500)
	}

	return c.JSON(fiber.Map{
		"Iklans": iklans,
	})
}

// POST /admin/iklan/api/create
func (iklanapicontroller *IklanAPI) CreateIklanAPI(c *fiber.Ctx) error {
	var iklan models.Iklan

	if err := c.BodyParser(&iklan); err != nil {
		return c.SendStatus(400)
	}

	// Get image
	// Parse the multipart form:
	if form, err := c.MultipartForm(); err == nil {
		// => *multipart.Form

		// Get all files from "image" key:
		files := form.File["image"]
		// => []*multipart.FileHeader

		// Loop through files:
		for _, file := range files {
			fmt.Println(file.Filename, file.Size, file.Header["Content-Type"][0])

			// Save the files to disk:
			iklan.Image = fmt.Sprintf("/public/iklan/images/%s", file.Filename)
			if err := c.SaveFile(file, fmt.Sprintf("public/iklan/images/%s", file.Filename)); err != nil {
				return err
			}
		}
	}

	// Get video
	// Parse the multipart form:
	if form, err := c.MultipartForm(); err == nil {
		// => *multipart.Form

		// Get all files from "video" key:
		files := form.File["video"]
		// => []*multipart.FileHeader

		// Loop through files:
		for _, file := range files {
			fmt.Println(file.Filename, file.Size, file.Header["Content-Type"][0])

			// Save the files to disk:
			iklan.Video = fmt.Sprintf("/public/iklan/videos/%s", file.Filename)
			if err := c.SaveFile(file, fmt.Sprintf("public/iklan/videos/%s", file.Filename)); err != nil {
				return err
			}
		}
	}

	// Save data iklan
	err := models.CreateIklan(iklanapicontroller.Db, &iklan)
	if err != nil {
		return c.SendStatus(400)
	}
	// if succeed
	return c.JSON(fiber.Map{
		"status": 200,
		"Iklans": iklan,
	})
}

// GET /admin/iklan/api/detail:id
func (iklanapicontroller *IklanAPI) DetailIklanAPI(c *fiber.Ctx) error {
	params := c.AllParams()

	intId, errs := strconv.Atoi(params["id"])

	if errs != nil {
		fmt.Println(errs)
	}

	var iklan models.Iklan
	err := models.ReadIklanById(iklanapicontroller.Db, &iklan, intId)
	if err != nil {
		return c.SendStatus(500)
	}

	return c.JSON(fiber.Map{
		"status": 200,
		"Iklans": iklan,
	})
}

// PUT /admin/iklan/api/edit/:id
func (iklanapicontroller *IklanAPI) EditIklanAPI(c *fiber.Ctx) error {
	var iklan models.Iklan

	params := c.AllParams() // "{"id": "1"}"
	intId, _ := strconv.Atoi(params["id"])
	iklan.Id = intId

	if err := c.BodyParser(&iklan); err != nil {
		return c.SendStatus(400)
	}

	// Get image
	// Parse the multipart form:
	if form, err := c.MultipartForm(); err == nil {
		// => *multipart.Form

		// Get all files from "image" key:
		files := form.File["image"]
		// => []*multipart.FileHeader

		// Loop through files:
		for _, file := range files {
			fmt.Println(file.Filename, file.Size, file.Header["Content-Type"][0])

			// Save the files to disk:
			iklan.Image = fmt.Sprintf("/public/iklan/images/%s", file.Filename)
			if err := c.SaveFile(file, fmt.Sprintf("public/iklan/images/%s", file.Filename)); err != nil {
				return err
			}
		}
	}

	// Get video
	// Parse the multipart form:
	if form, err := c.MultipartForm(); err == nil {
		// => *multipart.Form

		// Get all files from "video" key:
		files := form.File["video"]
		// => []*multipart.FileHeader

		// Loop through files:
		for _, file := range files {
			fmt.Println(file.Filename, file.Size, file.Header["Content-Type"][0])

			// Save the files to disk:
			iklan.Video = fmt.Sprintf("/public/iklan/videos/%s", file.Filename)
			if err := c.SaveFile(file, fmt.Sprintf("public/iklan/videos/%s", file.Filename)); err != nil {
				return err
			}
		}
	}

	// Save iklan
	err := models.UpdateIklan(iklanapicontroller.Db, &iklan)
	if err != nil {
		return c.SendStatus(400)
	}

	// if succeed
	return c.JSON(fiber.Map{
		"status": 200,
		"Iklans": iklan,
	})
}

// GET /admin/iklan/api/hapus/:id
func (iklanapicontroller *IklanAPI) DeleteIklanAPI(c *fiber.Ctx) error {
	params := c.AllParams() // "{"id": "1"}"

	intId, errs := strconv.Atoi(params["id"])

	if errs != nil {
		fmt.Println(errs)
	}

	var iklan models.Iklan
	err := models.DeleteIklanById(iklanapicontroller.Db, &iklan, intId)
	if err != nil {
		return c.SendStatus(500)
	}

	return c.JSON(fiber.Map{
		"message": "Data berhasil dihapus!",
	})
}
