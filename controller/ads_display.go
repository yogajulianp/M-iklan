package controller

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/M-iklan/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AdsDisplay struct {
	db *gorm.DB
}

func NewAdsDisplay(db *gorm.DB) *AdsDisplay {
	return &AdsDisplay{db: db}
}

func (ads *AdsDisplay) MountRouter(app *fiber.App) {
	router := app.Group("/ads")
	router.Get("/getads-image", ads.GetAdsImage)
	router.Get("/getads-video", ads.GetAdsVideo)
	router.Get("/iklandetail", ads.GetAdsAllType)
	router.Get("/iklandetail/:id", ads.GetAdsById)
	router.Get("/detailiklan/:id", ads.GetDetailsAds)
}

func (ads *AdsDisplay) GetAdsImage(c *fiber.Ctx) error {
	listIklan, err := models.GetAllIklanPublished(ads.db)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.ErrInternalServerError)
	}
	rand.Seed(time.Now().Unix())

	dataIklan := listIklan[rand.Intn(len(listIklan))]

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":       fiber.StatusOK,
		"image_path": fmt.Sprintf("localhost%s%s", os.Getenv("SERVER_PORT"), dataIklan.Image),
		"id_iklan":   dataIklan.ID,
		"id_user":    dataIklan.Vendor_fk,
	})
}

func (ads *AdsDisplay) GetAdsAllType(c *fiber.Ctx) error {
	listIklan, err := models.GetAllIklanPublished(ads.db)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.ErrInternalServerError)
	}
	rand.Seed(time.Now().Unix())

	dataIklan := listIklan[rand.Intn(len(listIklan))]

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":       fiber.StatusOK,
		"image_path": fmt.Sprintf("localhost%s%s", os.Getenv("SERVER_PORT"), dataIklan.Image),
		"video_path": fmt.Sprintf("localhost%s%s", os.Getenv("SERVER_PORT"), dataIklan.Video),
		"id_iklan":   dataIklan.ID,
		"id_user":    dataIklan.Vendor_fk,
	})

}

func (ads *AdsDisplay) GetAdsById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.ErrBadRequest)
	}
	iklan, err := models.GetAllIklanPublishedById(ads.db, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":       fiber.StatusOK,
		"image_path": fmt.Sprintf("localhost%s%s", os.Getenv("SERVER_PORT"), iklan.Image),
		"video_path": fmt.Sprintf("localhost%s%s", os.Getenv("SERVER_PORT"), iklan.Video),
		"id_iklan":   iklan.ID,
		"id_user":    iklan.Vendor_fk,
	})
}

func (ads *AdsDisplay) GetAdsVideo(c *fiber.Ctx) error {
	listIklan, err := models.GetAllIklanPublished(ads.db)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.ErrInternalServerError)
	}
	rand.Seed(time.Now().Unix())

	dataIklan := listIklan[rand.Intn(len(listIklan))]

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":       fiber.StatusOK,
		"video_path": fmt.Sprintf("localhost%s%s", os.Getenv("SERVER_PORT"), dataIklan.Video),
		"id_iklan":   dataIklan.ID,
		"id_user":    dataIklan.Vendor_fk,
	})

}

func (ads *AdsDisplay) GetDetailsAds(c *fiber.Ctx) error {
	id := c.Params("id")
	idn, _ := strconv.Atoi(id)
	var iklanDetail models.Iklan
	err := models.ReadIklanById(ads.db, &iklanDetail, idn)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.ErrInternalServerError)
	}
	fmt.Println(iklanDetail)

	return c.Render("DetailIklan", fiber.Map{
		"Title":     "Daftar Produk",
		"DataIklan": iklanDetail,
	})

}
