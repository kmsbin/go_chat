package migration

import (
	"go_chat/entities"
	"log"
)

func (migration *Migration) Up() {
	migration.GormConection.AutoMigrate(&entities.User{})
	migration.GormConection.AutoMigrate(&entities.Group{})

	user := &entities.User{
		ImageUrl:  "https://asgsadg",
		Name:      "aksdfkas",
		SecretKey: "563456",
	}
	response := migration.GormConection.Statement.DB.Create(user)

	log.Println(response)
}
