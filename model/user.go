package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ImageUrl  string `json:"image-url"`
	Name      string `json:"name"`
	SecretKey string `gorm:"primaryKey"`
}
