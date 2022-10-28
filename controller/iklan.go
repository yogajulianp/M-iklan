package controller

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/M-iklan/models"
)

type Iklan struct {
	Db *gorm.DB
}

func NewIklan(db *gorm.DB) *Iklan {
	return &Iklan{Db: db}
}

// Routing
func (iklancontroller *Iklan) RouteIklan(app *fiber.App) {
	router := app.Group("/admin/iklan")
	router.Get("/", iklancontroller.GetAllIklan)
	router.Get("/create", iklancontroller.AddIklan)
	router.Post("/create", iklancontroller.AddPostedIklan)
	router.Get("/detail/:id", iklancontroller.DetailIklan)
	router.Get("/edit/:id", iklancontroller.EditIklan)
	router.Post("/edit/:id", iklancontroller.AddEditedIklan)
	router.Get("/delete/:id", iklancontroller.DeleteIklan)
}

// GET /admin/iklan
func (iklancontroller *Iklan) GetAllIklan(c *fiber.Ctx) error {
	// Load all Iklans
	var iklans []models.Iklan
	err := models.ReadIklans(iklancontroller.Db, &iklans)
	if err != nil {
		return c.SendStatus(500)
	}

	return c.Render("admin/iklan/index", fiber.Map{
		"Title":  "M-Iklan",
		"Iklans": iklans,
	})
}

// GET /admin/iklan/create
func (iklancontroller *Iklan) AddIklan(c *fiber.Ctx) error {
	return c.Render("admin/iklan/create", fiber.Map{
		"Title": "Tambah Iklan",
	})
}

// POST /admin/iklan/create
func (iklancontroller *Iklan) AddPostedIklan(c *fiber.Ctx) error {
	//myform := new(models.Iklan)
	var iklan models.Iklan

	if err := c.BodyParser(&iklan); err != nil {
		return c.Redirect("/admin/iklan")
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

	// save iklan
	err := models.CreateIklan(iklancontroller.Db, &iklan)
	if err != nil {
		return c.Redirect("/admin/iklan")
	}
	// if succeed
	return c.Redirect("/admin/iklan")
}

// GET /admin/iklan/detail:id
func (iklancontroller *Iklan) DetailIklan(c *fiber.Ctx) error {
	params := c.AllParams() // "{"id": "1"}"

	intId, errs := strconv.Atoi(params["id"])

	if errs != nil {
		fmt.Println(errs)
	}

	var iklan models.Iklan
	err := models.ReadIklanById(iklancontroller.Db, &iklan, intId)
	if err != nil {
		return c.SendStatus(500)
	}

	return c.Render("admin/iklan/detail", fiber.Map{
		"Title": "Detail Iklan",
		"Iklan": iklan,
	})
}

// GET /admin/iklan/ubah/:id
func (iklancontroller *Iklan) EditIklan(c *fiber.Ctx) error {
	params := c.AllParams() // "{"id": "1"}"

	intId, _ := strconv.Atoi(params["id"])

	var iklan models.Iklan
	err := models.ReadIklanById(iklancontroller.Db, &iklan, intId)
	if err != nil {
		return c.SendStatus(500)
	}

	return c.Render("admin/iklan/edit", fiber.Map{
		"Title": "Ubah Iklan",
		"Iklan": iklan,
	})
}

// POST /admin/iklan/ubah/:id
func (iklancontroller *Iklan) AddEditedIklan(c *fiber.Ctx) error {
	var iklan models.Iklan

	params := c.AllParams() // "{"id": "1"}"
	intId, _ := strconv.Atoi(params["id"])
	iklan.Id = intId

	if err := c.BodyParser(&iklan); err != nil {
		return c.Redirect("/admin/iklan/edit")
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

	// save iklan
	err := models.UpdateIklan(iklancontroller.Db, &iklan)
	if err != nil {
		return c.Redirect("/admin/iklan/edit")
	}
	// if succeed
	return c.Redirect("/admin/iklan/edit")
}

// GET /admin/iklan/hapus/:id
func (iklancontroller *Iklan) DeleteIklan(c *fiber.Ctx) error {
	params := c.AllParams() // "{"id": "1"}"

	intId, errs := strconv.Atoi(params["id"])

	if errs != nil {
		fmt.Println(errs)
	}

	var iklan models.Iklan
	err := models.DeleteIklanById(iklancontroller.Db, &iklan, intId)
	if err != nil {
		return c.SendStatus(500)
	}

	return c.Redirect("/admin/iklan")
}
