package models

import (
	"time"

	"gorm.io/gorm"
)

type Admin struct {
	Id        uint           `form:"id" json:"id" gorm:"primaryKey"`
	Name      string         `form:"name" json:"name" validate:"required"`
	Email     string         `form:"email" json:"email" validate:"required"`
	Username  string         `form:"username" json:"username" validate:"required"`
	Password  string         `form:"password" json:"password" validate:"required"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func RegisAdmin(db *gorm.DB, newAdmin *Admin) (err error) {
	err = db.Create(newAdmin).Error
	if err != nil {
		return err
	}
	return nil
}
func ReadAdmin(db *gorm.DB, admin *[]Admin) (err error) {
	err = db.Find(admin).Error
	if err != nil {
		return err
	}
	return nil
}
