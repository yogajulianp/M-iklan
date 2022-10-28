package models

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Iklan struct {
	gorm.Model
	Name        string
	Image       string
	Video       string
	Description string
	View        int `gorm:"default:0"`
	Revenue     float64
	Vendor_fk   int
	IsPublished bool `gorm:"default:false"`
}

func GetAllIklanPublished(db *gorm.DB) (ListIklan []Iklan, err error) {
	_, err = db.Where("is_published IS TRUE AND deleted_at IS NULL").Find(&ListIklan).Rows()
	if len(ListIklan) == 0 {
		return ListIklan, fiber.ErrNotFound
	}
	return
}
func GetAllIklanPublishedById(db *gorm.DB, id int) (ListIklan *Iklan, err error) {
	_, err = db.Where("is_published IS TRUE AND deleted_at IS NULL AND id = ?", id).First(&ListIklan).Rows()
	return
}
func ReadIklanById(db *gorm.DB, iklan *Iklan, id int) (err error) {
	err = db.Where("id=?", id).First(iklan).Error
	if err != nil {
		return err
	}
	return nil
}
func UpdatePublikasiIklan(db *gorm.DB, iklan *Iklan) (err error) {

	fmt.Println("REVENUE", iklan.Revenue)
	db.Save(iklan)

	return nil
}
func CancelPublikasiById(db *gorm.DB, id int) (err error) {
	db.Where("id=?", id).Update("is_published", false)
	return nil
}
func UpdateViewsIklan(db *gorm.DB, iklan *Iklan, id int, view int) (err error) {

	err = db.Model(iklan).Where("id = ?", id).Update("view", view).Error
	if err != nil {
		return err
	}
	return nil
}

func GetTotalRevenueVendor(db *gorm.DB, id_vendor int) (float64, int, error) {
	rows, err := db.Raw("SELECT SUM(revenue),COUNT(id) FROM iklans WHERE vendor_fk = ? ", id_vendor).Rows()
	var totalRevenue float64
	var countIklan int
	defer rows.Close()
	for rows.Next() {
		fmt.Println("TESTS")
		rows.Scan(&totalRevenue, &countIklan)
	}
	return totalRevenue, countIklan, err
}

func RevenueCalculation(revenue float64) float64 {

	money := 10000  //dapat Rp. 10.000
	perView := 1000 //per 1.000 view
	revenue += float64(money) / float64(perView)

	return float64(revenue)
}
