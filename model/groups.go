package model

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	Name  string `gorm:"column:name"`
	Users []User `gorm:"many2many:GroupsID"`
}

type User struct {
	ID        uint    `gorm:"primarykey"`
	ImageUrl  string  `gorm:"column:image_url"`
	Name      string  `gorm:"column:name"`
	Groups    []Group `gorm:"many2many:GroupsID"`
	SecretKey uint    `gorm:"unique"`
	gorm.Model
}
