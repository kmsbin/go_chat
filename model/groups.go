package model

import "gorm.io/gorm"

type Groups struct {
	gorm.Model
	Name  string `json:"name"`
	Users []User `gorm:"foreignKey:SecretKey"`
}
