package migration

import (
	"go_chat/model"
)

func (migration *Migration) Up() {
	migration.GormConection.AutoMigrate(&model.User{})
	migration.GormConection.AutoMigrate(&model.Group{})
}
