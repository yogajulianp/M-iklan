package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Response struct {
	ID     string `json:"id"`
	Vendor   string `json:"vendor"`
	Views int    `json:"views"`
}

type AdsViewersController struct {
	Db *gorm.DB
}

func (controller *AdsViewersController) GetAdsViewer(c *fiber.Ctx) error {
	var url string
	var id = c.Query("id") // get specific vendor ad with id

	if id != "" {
		url = "http://localhost:3001/getviews/" + id //contoh url
	} else {
		return c.JSON(fiber.Map{
			"message": "Please input the Id",
		})
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Print(err.Error())
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}
	var responseObject Response
	json.Unmarshal(bodyBytes, &responseObject)

	return c.JSON(responseObject)
}


func InitAdsViewersController(db *gorm.DB) *AdsViewersController {
	return &AdsViewersController{Db: db}
}
