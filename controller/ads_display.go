package controller

import (
	"fmt"
	"math/rand"
	"os"
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
	})

}
