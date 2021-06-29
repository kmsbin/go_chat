package migration

import (
	"go_chat/model"
	"log"
)

func (migration *Migration) Seed() {

	user := &model.User{
		ImageUrl:  "https://asgsadg",
		Name:      "aksdfkas",
		SecretKey: "563456",
	}
	response := migration.GormConection.Statement.DB.Create(user)

	log.Println(response)

}
