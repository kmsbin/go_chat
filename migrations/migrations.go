package migration

import (
	"go_chat/model"
	"log"

	"gorm.io/gorm"
)

type Migration struct {
	GormConection *gorm.DB
}

func (conn *Migration) Migrate() {
	log.Println("migrating...")
	conn.GormConection.AutoMigrate(&model.User{}, &model.Groups{})

	// conn.GormConection.Migrator().CreateConstraint(&model.User{}, "fk_users")

	// conn.GormConection.Model(&model.User{})

}
