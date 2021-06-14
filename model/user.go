package model

import "gorm.io/gorm"

type User struct {
	ID        uint    `gorm:"primarykey"`
	ImageUrl  string  `gorm:"column:image_url"`
	Name      string  `gorm:"column:name"`
	Groups    []Group `gorm:"many2many:GroupsID"`
	SecretKey uint    `gorm:"unique"`
	gorm.Model
}
