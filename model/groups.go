package model

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	Name  string `gorm:"column:name"`
	Users []User `gorm:"many2many:GroupsID"`
}
