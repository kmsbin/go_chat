package migration

import (
	"go_chat/model"
)

func (migration *Migration) Down() {
	migration.GormConection.Migrator().DropTable(&model.User{})
	migration.GormConection.Migrator().DropTable(&model.Group{})
}
