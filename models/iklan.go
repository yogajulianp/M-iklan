package models

import (
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
