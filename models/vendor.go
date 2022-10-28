package models

import(
	"gorm.io/gorm"
	"time"
)

type Vendor struct {
	Id       	uint      `form:"id" json:"id" gorm:"primaryKey"`
	Name     	string  `form:"name" json:"name" validate:"required"`
	Email	 	string  `form:"email" json:"email" validate:"required"`
	Username	string    `form:"username" json:"username" validate:"required"`
	Password    string	`form:"password" json:"password" validate:"required"`
	Vendor_fk 	int
	CreatedAt time.Time		`json:"created_at"`
  	UpdatedAt time.Time		`json:"updated_at"`
 	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	
}

func RegisVendor(db *gorm.DB, newVendor *Vendor) (err error) {
	err = db.Create(newVendor).Error
	if err != nil {
		return err
	}
	return nil
}
func ReadVendor(db *gorm.DB, vendor *[]Vendor)(err error) {
	err = db.Find(vendor).Error
	if err != nil {
		return err
	}
	return nil
}


func ReadVendorById(db *gorm.DB, vendor *Vendor, id int)(err error) {
	err = db.Where("id=?", id).First(vendor).Error
	if err != nil {
		return err
	}
	return nil
}
func UpdateVendor(db *gorm.DB, vendor *Vendor)(err error) {
	db.Save(vendor)
	
	return nil
}
func DeleteVendorById(db *gorm.DB, vendor *Vendor, id int)(err error) {
	db.Where("id=?", id).Delete(vendor)
	
	return nil
}