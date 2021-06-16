package migration

import "gorm.io/gorm"

type Migration struct {
	GormConection *gorm.DB
}
