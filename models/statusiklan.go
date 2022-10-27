package models

import (
	"gorm.io/gorm"
)

type IklanStatus struct {
	ID uint `form:"id" json:"id" gorm:"primaryKey"`
	// AdminID uint	`form:"adminid" json:"adminid"`
	IklanID uint   `form:"iklanid" json:"iklanid"`
	Status  string `form:"status" json:"status" gorm:"default:uncheck"`
}

func CreateIklanStatus(db *gorm.DB, newIklanStatus *IklanStatus, iklanId uint) (err error) {
	if err := db.Create(newIklanStatus).Error; err != nil {
		return nil
	}
	return nil
}

func ReadIklanStatus(db *gorm.DB, iklanStatuses *[]IklanStatus) (err error) {
	if err := db.Find(iklanStatuses).Error; err != nil {
		return err
	}
	return nil
}

func ReadStatusById(db *gorm.DB, iklanStatus *IklanStatus, id uint) (err error) {
	if err = db.Where("id=?", id).First(iklanStatus).Error; err != nil {
		return err
	}
	return nil
}
func UpdateIklanStatus(db *gorm.DB, iklanStatus *IklanStatus) (err error) {
	db.Save(iklanStatus)
	return nil
}
