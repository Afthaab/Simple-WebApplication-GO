package entities

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	First_Name string
	Last_name  string
	Username   string
	Password   string
}
