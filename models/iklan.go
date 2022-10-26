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
