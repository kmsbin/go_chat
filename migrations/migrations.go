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
	conn.GormConection.Migrator().DropTable(&model.User{})
	conn.GormConection.Migrator().DropTable(&model.Group{})

	conn.GormConection.AutoMigrate(&model.User{})
	conn.GormConection.AutoMigrate(&model.Group{})
	// conn.GormConection.Migrator().CreateTable(&model.User{})

	// conn.GormConection.Migrator().CreateTable(&model.Groups{})
	// conn.GormConection.Migrator().CreateConstraint(&model.Groups{}, "fk_clients_group")

	// conn.GormConection.Migrator().CreateConstraint(&model.User{}, "fk_users")

	// conn.GormConection.Model(&model.User{})

}
