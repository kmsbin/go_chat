package migration

import (
	"go_chat/entities"
)

func (migration *Migration) Down() {
	migration.GormConection.Migrator().DropTable(&entities.User{})
	migration.GormConection.Migrator().DropTable(&entities.Group{})
}
