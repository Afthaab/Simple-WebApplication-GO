package entities

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	First_Name string `gorm:"not null"`
	Last_name  string `gorm:"not null"`
	Username   string `gorm:"not null;unique;default:null"`
	Password   string `gorm:"not null"`
	IsAdmin    bool   `gorm:"default:false"`
}
